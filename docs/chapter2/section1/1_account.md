# アカウント機能の実装

## 本日の目的

`main.go` の handler の設定部分を見てみましょう。

```go
func main () {
    (省略)
    h := handler.NewHandler(db)  // [!code hl]
    e := echo.New()  // [!code hl]
    // [!code hl]
    e.GET("/cities/:cityName", h.GetCityInfoHandler)  // [!code hl]
    e.POST("/cities", h.PostCityHandler)  // [!code hl]
    // [!code hl]
    err = e.Start(":8080")  // [!code hl]
    (省略)
}
```

今回の目標は、 `/cities/` で始まる api 2 つ (`getCityInfoHandler`, `postCityHandler`) に対して、
ログインしているかどうかを判定して、ログインしていなければリクエストを拒否するように実装することです。

用語を使わずに言えば、`City` を新たに追加したり、`City` の情報を得るのにログインを必須にする、ということです。

実装は以下のように進めます。

1. アカウントを作成できるようにする
2. ログインを実装する
3. ログインしないと利用できないようにする

## ライブラリのインストール

新たにライブラリを導入するので以下のコマンドを実行します。

```sh
go get -u github.com/labstack/echo-contrib/session
go get -u github.com/srinathgs/mysqlstore
```

では、アカウントの作成を実装していきましょう。

アカウントの作成は、以下の手順で進んでいきます。

1. クライアントから`Username`と`Password`をリクエストとして受け取る
2. `Username`と`Password`のバリデーション(値が正当かのチェック)を行う
3. 既に同じ`Username`のユーザーが登録されていないかチェックする
4. `Password`をハッシュ化する
5. ユーザーをデーターベースに登録する

## 下準備

### テーブルの作成

アカウントを管理するテーブル `users` を作成しましょう。`main.go`に以下を追加します。

```go
func main() {
    (省略)
	// データベースに接続
    db, err := sqlx.Open("mysql", conf.FormatDSN())
    if err != nil {
    log.Fatal(err)
    }
	
    // usersテーブルが存在しなかったら、usersテーブルを作成する // [!code ++]
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))") // [!code ++]
    if err != nil { // [!code ++]
    log.Fatal(err) // [!code ++]
    } // [!code ++]
    
    h := handler.NewHandler(db)
    e := echo.New()
	(省略)
}
```

## signUpHandler の実装

続いて、アカウントを作成するハンドラーである `signUpHandler` を `handler.go` に実装していきましょう。

```go
func (h *Handler) SignUpHandler(c echo.Context) error { // [!code ++]
} // [!code ++]
```

この `signUpHandler` に以下のものを順番に実装していきます。

### 1. リクエストの受け取り

`signUpHandler` の外に以下の構造体を追加します。

```go
type LoginRequestBody struct { // [!code ++]
	Username string `json:"username,omitempty" form:"username"` // [!code ++]
	Password string `json:"password,omitempty" form:"password"` // [!code ++]
} // [!code ++]
```

次に、`signUpHandler`の中に以下を追加します。

```go
func (h *Handler) SignUpHandler(c echo.Context) error {
	// リクエストを受け取り、reqに格納する // [!code ++]
	req := LoginRequestBody{} // [!code ++]
    err := c.Bind(&req) // [!code ++]
    if err != nil { // [!code ++]
        return echo.NewHTTPError(http.StatusBadRequest, "bad request body") // [!code ++]
    } // [!code ++]
}
```

ここでは、req 変数に requestBody の json 情報を格納しています。`LoginRequestBody` 型を見れば分かる通り、ここには UserName と
Password が格納されています。

### 2. リクエストの検証

```go
func (h *Handler) SignUpHandler(c echo.Context) error {
    // リクエストを受け取り、reqに格納する
    req := LoginRequestBody{}
    err := c.Bind(&req)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
    }
    
    // バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す) // [!code ++]
    if req.Password == "" || req.Username == "" { // [!code ++]
        return c.String(http.StatusBadRequest, "Username or Password is empty") // [!code ++]
    } // [!code ++]
}
```

ここでは、UserName と Password が正しく入っているのかをチェック（バリデーションといいます）します。
入っていない場合は、与えられた入力が正しくない間違った形式なので、 400 (Bad Request) をレスポンスします。

### 3. アカウントの存在チェック

```go
func (h *Handler) SignUpHandler(c echo.Context) error {
    (省略)
	
	// 登録しようとしているユーザーが既にデータベース内に存在するかチェック// [!code ++]
	var count int// [!code ++]
	err = h.db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)// [!code ++]
    if err != nil {// [!code ++]
        log.Println(err)// [!code ++]
        return c.NoContent(http.StatusInternalServerError)// [!code ++]
    }// [!code ++]
    // 存在したら409 Conflictを返す// [!code ++]
    if count > 0 {// [!code ++]
        return c.String(http.StatusConflict, "Username is already used")// [!code ++]
    }// [!code ++]
}
```

`"SELECT COUNT(*) FROM users WHERE Username=?", req.Username` で、指定された UserName を持つユーザーの数を見ます。

結果は `count` 変数に格納されます。

もしすでに居た場合は、そのユーザーが存在しているので処理は受け付けず、 409 (Conflict) をレスポンスします。

### 4. パスワードのハッシュ化

ここまでは「リクエストを実行しても本当に問題がないか」を検証していました。
ユーザーはまだ存在していなくて、パスワードとユーザー名がある事まで確認できれば、リクエストを処理できます。なのでここから処理を行っていきます。

```go
func (h *Handler) SignUpHandler(c echo.Context) error {
    (省略)
	
    // パスワードをハッシュ化する// [!code ++]
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)// [!code ++]
    // ハッシュ化に失敗したら500 InternalServerErrorを返す// [!code ++]
    if err != nil {// [!code ++]
        log.Println(err)// [!code ++]
        return c.NoContent(http.StatusInternalServerError)// [!code ++]
    }// [!code ++]
}
```

まずはパスワードのハッシュ化です。 **パスワードは平文で保存してはいけません！** パスワードを DB に保管するときは、必ずハッシュ化をしましょう。

:::info 参考: ソルトについて  
ソルトという手法を用いることで、事前計算されたテーブルを使用する攻撃から守ることができます。

ここで用いている`bcrypt`というライブラリでは、これを自動でやってくれています。

参考: <https://en.wikipedia.org/wiki/Salt_(cryptography)>
:::

`bcrypt`というのはいい感じにハッシュ化してくれるライブラリです。セキュリティに関わるものは自分で実装すると穴だらけになりやすいので、積極的にライブラリに頼りましょう。

### 5. ユーザーの作成

```go
func (h *Handler) SignUpHandler(c echo.Context) error {
    (省略)
	
	// ユーザーを登録する// [!code ++]
	_, err = h.db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)// [!code ++]
	// 登録に失敗したら500 InternalServerErrorを返す// [!code ++]
	if err != nil {// [!code ++]
		log.Println(err)// [!code ++]
		return c.NoContent(http.StatusInternalServerError)// [!code ++]
	}// [!code ++]
	// 登録に成功したら201 Createdを返す// [!code ++]
	return c.NoContent(http.StatusCreated)// [!code ++]
}
```

`Username`, `HashedPassword` を持つ User レコードをデータベースに追加しましょう。

何かしらのエラーによって生成できなかった場合は err にその内容が詰め込まれます。
ユーザーのリクエストは問題なく、ここでエラーが発生した場合はサーバー側で何かが発生したということなので、
500 (InternalServer Error) をレスポンスします。

ここで、どんなエラーが発生したかをユーザーに直接伝えるのはセキュリティの観点から △ です。
ログで出力するだけにして、ユーザー側には 500 という情報だけ渡しましょう。

もし err がなければ、それはうまく成功したということです。 201 (Created) をレスポンスしましょう！

## 完成！
これで実装は終わりです。すべてを実装したら、以下のようになります。

```go
type LoginRequestBody struct {
    Username string `json:"username,omitempty" form:"username"`
    Password string `json:"password,omitempty" form:"password"`
}

func (h *Handler) SignUpHandler(c echo.Context) error {
	// リクエストを受け取り、reqに格納する
	req := LoginRequestBody{}
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body")
	}

	// バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す)
	if req.Password == "" || req.Username == "" {
		return c.String(http.StatusBadRequest, "Username or Password is empty")
	}

	// 登録しようとしているユーザーが既にデータベース内に存在するかチェック
	var count int
	err = h.db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	// 存在したら409 Conflictを返す
	if count > 0 {
		return c.String(http.StatusConflict, "Username is already used")
	}

	// パスワードをハッシュ化する
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// ハッシュ化に失敗したら500 InternalServerErrorを返す
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// ユーザーを登録する
	_, err = h.db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)
	// 登録に失敗したら500 InternalServerErrorを返す
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	// 登録に成功したら201 Createdを返す
	return c.NoContent(http.StatusCreated)
}
```

最後に、`main.go` に、先ほど書いたハンドラーを追加しましょう。

```go
func main(){
    (省略)
    h := handler.NewHandler(db)
    e := echo.New()
    
    e.POST("/signup", h.SignUpHandler) // [!code ++]
    
    e.GET("/cities/:cityName", h.GetCityInfoHandler)
    e.POST("/cities", h.PostCityHandler)
    
    err = e.Start(":8080")
    (省略)
}
```

:::warning
このコードは後々の回で使用するので、エンドポイントのパス (`/signup` など) は変更しないでください！

エンドポイントの追加は問題ないので、試したい場合は新しくハンドラーを実装しましょう。
:::

ここまでできたら、実行して、Postman 等を用いて正しく実装できているかデバッグしてみましょう。以下は例です。
![](images/3/postman-debug-signup.png)
上手く作成できれば Status 201 が返ってくるはずです。

正しく API を叩いたあとに、テーブルに意図したユーザー名と、ハッシュ化されたパスワードが入っていますか？

:::details 確認に使う SQL クエリ

```sql
USE
world;
SELECT *
FROM users;
```

:::
