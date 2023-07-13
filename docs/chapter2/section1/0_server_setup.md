# main.go
書きます
:::warning
下に解説が書いてあります
:::


```
$ go get -u github.com/labstack/echo-contrib/session
$ go get -u github.com/srinathgs/mysqlstore
```


```go=
package main

import (
	"fmt"
	"errors"
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        sql.NullString `json:"name,omitempty"  db:"Name"`
	CountryCode sql.NullString `json:"countryCode,omitempty"  db:"CountryCode"`
	District    sql.NullString `json:"district,omitempty"  db:"District"`
	Population  sql.NullInt64    `json:"population,omitempty"  db:"Population"`
}

var (
	db *sqlx.DB
)

func main() {
	_db, err := sqlx.Connect(
		"mysql", 
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOSTNAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	db = _db

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/login", postLoginHandler)
	e.POST("/signup", postSignUpHandler)

	withLogin := e.Group("")
	withLogin.Use(checkLogin)
	withLogin.GET("/cities/:cityName", getCityInfoHandler)

	e.Start(":<ポート番号>")
}

type LoginRequestBody struct {
	Username string `json:"username,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}

type User struct {
	Username   string `json:"username,omitempty"  db:"Username"`
	HashedPass string `json:"-"  db:"HashedPass"`
}

func postSignUpHandler(c echo.Context) error {
	req := LoginRequestBody{}
	c.Bind(&req)

	// もう少し真面目にバリデーションするべき
	if req.Password == "" || req.Username == "" {
		// エラーは真面目に返すべき
		return c.String(http.StatusBadRequest, "項目が空です")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	// ユーザーの存在チェック
	var count int

	err = db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "ユーザーが既に存在しています")
	}

	_, err = db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}
	return c.NoContent(http.StatusCreated)
}

func postLoginHandler(c echo.Context) error {
	req := LoginRequestBody{}
	c.Bind(&req)

	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.NoContent(http.StatusForbidden)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userName"] = req.Username
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if sess.Values["userName"] == nil {
			return c.String(http.StatusForbidden, "please login")
		}
		c.Set("userName", sess.Values["userName"].(string))

		return next(c)
	}
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")

	city := City{}
	db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName)
	if !city.Name.Valid {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}

```

## 解説
### main関数
```go=34
_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	db = _db
```
データベースへ接続しています。

`os.GetEnv("ENV")`は ENV という名前の環境変数を取得する関数です。環境変数には DB の情報が入っています。

接続出来なかった場合は err にそのエラー内容が格納されます。

```go=47
store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
            panic(err)
	}
```

セッションストアを設定しています。セッションは HTTP リクエストを送ってきた User が誰かを確認するために使うのですが、それらを覚えておくための場所を用意しているイメージです(ちょい雑な気がする)

54 行目の`e.Use(session.Middleware(store))`が echo にこのセッションストアを使ってと言っている部分です。

それ以降では echo についての設定をしています。具体的にはハンドラの登録やサーバーの起動などを行っています。ここら辺は知っていると思っております。

### postSignUpHandler関数
これは User 登録を行なうための関数です。
```go=80
req := LoginRequestBody{}
c.Bind(&req)
```
req にリクエスト情報を入れています。ここには UserName と Password が格納されているはずです。

```go=83
// もう少し真面目にバリデーションするべき
if req.Password == "" || req.Username == "" {
	// エラーは真面目に返すべき
	return c.String(http.StatusBadRequest, "項目が空です")
}

```

ここでは、本当に UserName と Password が入っているのかをチェックしています。入っていなければ不正なリクエストなので、400(Bad Request)を返しています。

```go=89
hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
}
```
基本的に Password を平文で保存しておくのは危険です！　従って、パスワードを DB に格納するときはハッシュ化を行ってから格納します。(ハッシュ化については調べてください)

`bcrypt`というのはハッシュ化をいい感じにやってくれるライブラリです。それを使ってパスワードをハッシュ化しています。

```go=94
// ユーザーの存在チェック
var count int

err = db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)
if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
}

if count > 0 {
    return c.String(http.StatusConflict, "ユーザーが既に存在しています")
}

```
`db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)`という部分は、db に`req.Username`という名前の User は何人いますか？　という問い合わせをしています。

その結果は`count`に格納されています。同名のユーザーがいたら困るので、そういう場合はその名前の User はもういるからダメだよというレスポンスを返します。



```go=106
_, err = db.Exec("INSERT INTO users (Username, HashedPass) VALUES (?, ?)", req.Username, hashedPass)
if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
    }
return c.NoContent(http.StatusCreated)
```

`db.Exec`はクエリを実行する関数です。ここでは Username,HashedPassword を持つ User を生成しようとしてます。

何かしらのエラーによって生成できなかった場合は err にその内容が詰め込まれます。上までの処理から理論上は User を生成できるはずなので、ここで何かエラーが出たとするとそれはサーバー側の問題になります。従ってここで返しているエラーコードは 500 になっています。

### postLoginHandler関数
```go=114
req := LoginRequestBody{}
c.Bind(&req)

user := User{}
err := db.Get(&user, "SELECT * FROM users WHERE username=?", req.Username)
if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
}
```
req の代入のところは SignUp のところと同じです。

下の部分ではリクエストで送られてきた UserName と Password を持つユーザーは存在するのか？　という問い合わせをしています。存在した場合は`user`にそのユーザーの情報が入ります。

```go=123
err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
if err != nil {
    if err == bcrypt.ErrMismatchedHashAndPassword {
        return c.NoContent(http.StatusForbidden)
    } else {
        return c.NoContent(http.StatusInternalServerError)
    }
}
```

SignUp の方にも書きましたが、パスワードを平文で保存するのは良くないということでハッシュ化されています。

送られてきたパスワードが正しいパスワードなら、ユーザー生成時に行った計算と同様のものを行うことで同じハッシュ文字列を得ることができ、それによってこのパスワードが正しいかどうかを確認できます。その処理を行っています。

この関数はハッシュが一致するときに限り`err`が`nil`で返ってきます。一致しない場合のエラーは`bcrypt.ErrMismatchedHashAndPassword`というものです。

従って、これの場合はパスワードが違うよというレスポンスを返し、それ以外のエラーの場合は 500 を返しています。


```go=132
sess, err := session.Get("sessions", c)
if err != nil {
    fmt.Println(err)
    return c.String(http.StatusInternalServerError, "something wrong in getting session")
}
sess.Values["userName"] = req.Username
sess.Save(c.Request(), c.Response())

return c.NoContent(http.StatusOK)
```

セッションに登録する処理です。セッションとは、今来た人が次来たとき、同じ人であることを確認するための仕組みです。

ここではセッションにログインした人を登録しているんだなあということが分かれば大丈夫です。

### checklogin関数
まず、この関数は handler ではありません。handler 関数を返す関数です。(middleware と呼ばれています) リクエスト→middleware→handler という順序で処理されます。

middleware から次の handler を呼び出すには`next(c)`と書きます。

このミドルウェアはリクエストを送ったユーザーがログインしているのかをチェックし、ログインしているなら echo の Context にそのユーザーの UserName を登録します。

```go=144
return func(c echo.Context) error {
    sess, err := session.Get("sessions", c)
    if err != nil {
        fmt.Println(err)
        return c.String(http.StatusInternalServerError, "something wrong in getting session")
    }
```

セッションを取得しています。

本当はリクエストヘッダを見ることでどのセッションを取り出すかを決めています(セッションは各ユーザーに存在するので)

```go=151                          
    if sess.Values["userName"] == nil {
        return c.String(http.StatusForbidden, "please login")
    }
    c.Set("userName", sess.Values["userName"].(string))

    return next(c)
}
```

Login 時の処理を思い出すと、セッションには"userName"をキーとしてユーザーの名前が登録されていました。

従って、ここに名前が入っているならば、今リクエストを送った人は過去にログインをした人と同じということがわかります。逆に、何も入っていなければリクエストを送った人はログインをしていません。

これを利用して、ログインしていない場合には後続に処理を渡すことをせず途中で処理を止めています。

### getCityInfoHandler関数
Param で与えられた cityName と一致する city を返す関数です。
```go=161
cityName := c.Param("cityName")

city := City{}
db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName)
if !city.Name.Valid {
    return c.NoContent(http.StatusNotFound)
}

return c.JSON(http.StatusOK, city)
```

`c.Param`でリクエストに含まれている cityName を取得します。

db には、cityName を持つ city をくださいといっています。出てきたもの中で一番最初にヒットしたものが city の中に格納されます。


# データベースの準備
users テーブルを作成します。
```
mysql> CREATE TABLE `users` ( `Username` VARCHAR(30) NOT NULL , `HashedPass` VARCHAR(200) NOT NULL , PRIMARY KEY (`Username`)) ENGINE = InnoDB;
```
`mysql > SHOW TABLES;`で作成できているか確認します。
<img src="https://md.trap.jp/uploads/upload_c0f4bf39ee2d2b5b6fd1e2e053a8d39e.png" width=40%>



# 検証
:::warning
全て Postman での検証です
`go run main.go`でサーバーを起動した状態で行ってください
:::

http://133.130.109.224:<ポート番号>/cities/Tokyo へ
初めに普通にアクセスするとダメです
![](https://md.trapti.tech/uploads/upload_96a03d609e761150a2136963dd34006a.png)

ユーザーを作成します。
上手く作成できれば Status 201 が返ってくるはずです。
![](https://md.trapti.tech/uploads/upload_4d891187b392debc9732aeff7ecaca08.png)

そのままパスを変えてログインリクエストを送ります。
![](https://md.trapti.tech/uploads/upload_7b21cf42397801806bab12f5180ce888.png)

ログインに成功したら、レスポンスの方の Cookies を開いて value の中身をコピーします
![](https://md.trapti.tech/uploads/upload_13756985fc7e93bd6d032083d340ea6b.png)

リクエストの方の Headers で Cookie をセットします。

Key に`Cookie`を
Value に`sessions=(コピーした値);`をセットします(既に自動で入っている場合もあります、その場合は追加しなくて大丈夫です)。


もう一度 http://133.130.109.224:<ポート番号>/cities/Tokyo にアクセスすると正常に API が取れるようになりました。
![](https://md.trapti.tech/uploads/upload_59c6c86e127d982f511946d2a183d0a6.png)

ここで、作成されたユーザーがデータベースに保存されていることを確認してみましょう。
`mysql > SELECT * FROM users;`
![](https://md.trap.jp/uploads/upload_f713b7da16df6729729a25ca2b5a6816.png)
ユーザー名とハッシュ化されたパスワードが確認できますね。
![](https://md.trap.jp/uploads/upload_7f007d73bd0ff508dff12246546b1a5b.png)
ちょっと分かりにくい表示ですが、セッションもしっかり確認できます。

:::success
jikkyo チャンネルで URL を貼って他の人にユーザーを作成してもらうと他の人のデータも見ることができます。やってみましょう。
:::

:::success
`checkLogin`関数の中で`c.Set("userName", ~~)`としている部分があります。
それを使って、ログインしているユーザーの名前を表示する API を作ってみましょう。
:::

:::info
# 難
TODO リストのサーバーとしての API を考えて作ってみましょう
:::

:::info
ここまでで Web サービスを作るコードの知識として必要な要素は全て網羅したつもりです。
みなさんはもう何でも作れるわけです!!!!

今回の目標は Twitter クローンの作成なので、早速作り始めても OK です!!
もし作り始める場合は https://github.com/tohutohu/naro-portal から fork して作業をして、Pull Request を出してもらえると講師や TA やその他暇な人が勝手にレビューをします！　全部完成していなくてもここまでできたので見てくださいとかでも構いません！
ぜひ皆さん使ってください！

fork とか Pull Request とかがわからない人は TA に言ってください(3 分くらいで終わるので)
:::
