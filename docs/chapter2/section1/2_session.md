# セッション管理機構の実装

## セッションストアを設定する
`main.go`に以下を追加しましょう。
```go
func main() {
	(省略)
	// usersテーブルが存在しなかったら、usersテーブルを作成する
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
	}

	// セッションの情報を記憶するための場所をデータベース上に設定 // [!code ++]
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token")) // [!code ++]
	if err != nil { // [!code ++]
		log.Fatal(err) // [!code ++]
	} // [!code ++]

	h := handler.NewHandler(db)
	e := echo.New()
	e.Use(middleware.Logger())       // ログを取るミドルウェアを追加 // [!code ++]
	e.Use(session.Middleware(store)) // セッション管理のためのミドルウェアを追加 // [!code ++]

	e.POST("/signup", h.SignUpHandler)
	(省略)
}
```
これらはセッションストアの設定です。

最初に、セッションの情報を記憶するための場所をデータベース上に設定します。

この仕組みを使用するために、 `e.Use(session.Middleware(store))` を含めてセッションストアを使ってね〜、って echo に命令しています。

`e.Use(middleware.Logger())` は文字通りログを取るものです。ついでに入れましょう。

:::tip
`"secret-token"`は、暗号化/復号化の際に使われる秘密鍵です。  
実際に運用するときはこの"secret-token"を独自の値にしてください。環境変数などで管理するのが良いでしょう。
:::

## LoginHandler の実装
続いて、`LoginHandler` を `handler.go` に実装していきましょう。

```go
func (h *Handler) LoginHandler(c echo.Context) error { // [!code ++]
} // [!code ++]
```
`LoginHandler` の外に以下の構造体を追加します。
```go
type User struct { // [!code ++]
	Username   string `json:"username,omitempty"  db:"Username"` // [!code ++]
	HashedPass string `json:"-"  db:"HashedPass"` // [!code ++]
} // [!code ++]
```
`LoginHandler` を実装していきます。
```go
func (h *Handler) LoginHandler(c echo.Context) error {
	// リクエストを受け取り、reqに格納する // [!code ++]
	var req LoginRequestBody // [!code ++]
	err := c.Bind(&req) // [!code ++]
	if err != nil { // [!code ++]
		return c.String(http.StatusBadRequest, "bad request body") // [!code ++]
	} // [!code ++]

	// バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す) // [!code ++]
	if req.Password == "" || req.Username == "" { // [!code ++]
		return c.String(http.StatusBadRequest, "Username or Password is empty") // [!code ++]
	} // [!code ++]

	// データベースからユーザーを取得する // [!code ++]
	user := User{} // [!code ++]
	err = h.db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username) // [!code ++]
	if err != nil { // [!code ++]
		if errors.Is(err, sql.ErrNoRows) { // [!code ++]
			return c.NoContent(http.StatusUnauthorized) // [!code ++]
		} else { // [!code ++]
			log.Println(err) // [!code ++]
			return c.NoContent(http.StatusInternalServerError) // [!code ++]
		} // [!code ++]
	} // [!code ++]
}
```
req への代入は signUpHandler と同じです。UserName と Password が入っているかも確認しましょう。

パスワードの一致チェックをするために、データベースからユーザーを取得してきましょう。

ユーザーが存在しなかった場合は `sql.ErrNoRows` というエラーが返ってきます。
もしそのエラーなら 401 (Unauthorized)、そうでなければ 500 (Internal Server Error) です。
もし 404 (Not Found) とすると、「このユーザーはパスワードが違うのではなく存在しないんだ」という事がわかってしまい（このユーザーは存在していてパスワードは違う事も分かります）、セキュリティ上のリスクに繋がります。

:::tip
ここで、エラーチェックは基本的に errors.Is を使いましょう。     
参考: <https://pkg.go.dev/errors#Is>
:::
```go
func (h *Handler) LoginHandler(c echo.Context) error {
	(省略)
	// データベースからユーザーを取得する
	user := User{}
	err = h.db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.NoContent(http.StatusUnauthorized)
		} else {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	// パスワードが一致しているかを確かめる // [!code ++]
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password)) // [!code ++]
	if err != nil { // [!code ++]
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) { // [!code ++]
			return c.NoContent(http.StatusUnauthorized) // [!code ++]
		} else { // [!code ++]
			return c.NoContent(http.StatusInternalServerError) // [!code ++]
		} // [!code ++]
	} // [!code ++]
}
```

データベースに保存されているパスワードはハッシュ化されています。

ハッシュ化は不可逆な処理なので、ハッシュ化されたものから原文を調べることはできません。確認する際はもらったパスワードをハッシュ化することで行います。

これは `bcrypt.CompareHashAndPassword` が行ってくれるのでそれに乗っかりましょう。

- この関数はハッシュが一致すれば返り値が `nil` となります
- 一致しない場合、 `bcrypt.ErrMismatchedHashAndPassword` が返ってきます
- 処理中にこれ以外の問題が発生した場合は、返り値はエラー型の何かです

従って、これらのエラーの内容に応じて、 500 (Internal Server Error), 401 (Unauthorized) を返却するか、処理を続行するかを選択していきます。
```go
func (h *Handler) LoginHandler(c echo.Context) error {
	(省略)
	// パスワードが一致しているかを確かめる
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.NoContent(http.StatusUnauthorized)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	// セッションストアに登録する // [!code ++]
	sess, err := session.Get("sessions", c) // [!code ++]
	if err != nil { // [!code ++]
		log.Println(err) // [!code ++]
		return c.String(http.StatusInternalServerError, "something wrong in getting session") // [!code ++]
	} // [!code ++]
	sess.Values["userName"] = req.Username // [!code ++]
	sess.Save(c.Request(), c.Response()) // [!code ++]

	return c.NoContent(http.StatusOK) // [!code ++]
}
```
セッションストアに登録します。
セッションの `userName` という値にそのユーザーの名前を格納していることは覚えておきましょう。

ここまで書いたら、 `LoginHandler` を使えるようにしましょう。

```go
func main() {
	(省略)
	e.Use(session.Middleware(store)) // セッション管理のためのミドルウェアを追加

	e.POST("/signup", h.SignUpHandler)
	e.POST("/login", h.LoginHandler) // [!code ++]

	e.GET("/cities/:cityName", h.GetCityInfoHandler)
	(省略)
}
```

:::details ここまでの全体像
::: code-group
<<<@/chapter2/section1/src/2_session/main.go{go:line-numbers}[main.go]
<<<@/chapter2/section1/src/2_session/handler.go{go:line-numbers}[handler.go]
:::

## userAuthMiddleware の実装

続いて、`userAuthMiddleware` を実装します。
まず、これは Handler ではなく Middleware と呼ばれます。

送られてくるリクエストは、Middleware を経由して、 Handler に流れていきます。

Middleware から次の Middleware/Handler を呼び出す際は `next(c)` と記述します。 Middleware の実装は難しいので、なんとなく理解できれば十分です。

以下を`handler.go`に追加しましょう。
```go
func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc { // [!code ++]
	return func(c echo.Context) error { // [!code ++]
		sess, err := session.Get("sessions", c) // [!code ++]
		if err != nil { // [!code ++]
			log.Println(err) // [!code ++]
			return c.String(http.StatusInternalServerError, "something wrong in getting session") // [!code ++]
		} // [!code ++]
		if sess.Values["userName"] == nil { // [!code ++]
			return c.String(http.StatusUnauthorized, "please login") // [!code ++]
		} // [!code ++]
		c.Set("userName", sess.Values["userName"].(string)) // [!code ++]
		return next(c) // [!code ++]
	} // [!code ++]
} // [!code ++]
```

関数が関数を呼び出していて混乱しそうですが、 2 行目から 13 行目が本質で、外側はおまじないと考えて良いです。

この Middleware はリクエストを送ったユーザーがログインしているのかをチェックし、
ログインしているなら Context (`c`) にそのユーザーの UserName を設定します。

セッションを取得し、ログイン時に設定した `userName` の値を確認しに行きます。

ここで名前が入っていればリクエストの送信者はログイン済みで、そうでなければログインをしていないことが分かります。

これを利用して、ログインしていない場合には処理をここで止めて 401 (Unauthorized) を返却し、していれば次の処理 (`next(c)`)
に進みます。

最後に、Middleware を設定しましょう。
グループ機能を利用して、 `withAuth` に設定されてるエンドポイントは `userAuthMiddleware` を処理してから処理する、という設定をします。

```go
func main() {
	(省略)
	e.POST("/login", h.LoginHandler)

	e.GET("/cities/:cityName", h.GetCityInfoHandler) // [!code --]
	e.POST("/cities", h.PostCityHandler) // [!code --]
	withAuth := e.Group("") // [!code ++]
	withAuth.Use(handler.UserAuthMiddleware) // [!code ++]
	withAuth.GET("/cities/:cityName", h.GetCityInfoHandler) // [!code ++]
	withAuth.POST("/cities", h.PostCityHandler) // [!code ++]

	err = e.Start(":8080")
	(省略)
}
```

これで、この章の目標である「ログインしないと利用できないようにする」が達成されました。

## GetMeHandler の実装

最後に、 `GetMeHandler` を実装します。叩いたときに自分の情報が返ってくるエンドポイントです。

以下を `handler.go` に追加しましょう。
```go
type Me struct { // [!code ++]
	Username string `json:"username,omitempty"  db:"username"` // [!code ++]
} // [!code ++]
```
```go
func GetMeHandler(c echo.Context) error { // [!code ++]
	return c.JSON(http.StatusOK, Me{ // [!code ++]
		Username: c.Get("userName").(string), // [!code ++]
	}) // [!code ++]
} // [!code ++]
```

アクセスしているユーザーの`userName`をセッションから取得して返しています。
`userAuthMiddleware` を実行したあとなので、`c.Get("userName").(string)` によって userName を取得できます。

`main.go`に`withAuth.GET("/me", handler.GetMeHandler)`を追加しましょう。
```go
func main() {
	(省略)
	withAuth := e.Group("")
	withAuth.Use(handler.UserAuthMiddleware)
	withAuth.GET("/me", handler.GetMeHandler) // [!code ++]
	withAuth.GET("/cities/:cityName", h.GetCityInfoHandler)
	withAuth.POST("/cities", h.PostCityHandler)

	err = e.Start(":8080")
	(省略)
}
```
