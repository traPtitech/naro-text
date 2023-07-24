# セッション管理機構の実装

## セッションストアの設定

<<<@/chapter2/section1/src/0/final/code.go#setup_session

セッションストアを設定しましょう。

ここでは、セッションの情報を記憶するための場所をデータベース上に設定しています。

この仕組みを使用するために、 `e.Use(session.Middleware(store))` を含めてセッションストアを使ってね〜、って echo に命令しています。

`e.Use(middleware.Logger())` は文字通りログを取るものです。ついでに入れましょう。

## loginHandler の実装

```go
func signUpHandler(c echo.Context) error {
}
```

つづいて `loginHandler` を実装していきます。これも `handler.go` に実装しましょう。

<<<@/chapter2/section1/src/0/final/code.go#post_req

req への代入は signUpHandler と同じです。UserName と Password が入っているかも確認しましょう。

パスワードの一致チェックをするために、データベースからユーザーを取得してきましょう。

ユーザーが存在しなかった場合は `sql.ErrNoRows` というエラーが返ってきます。
もしそのエラーなら 401 (Unauthorized)、そうでなければ 500 (Internal Server Error) です。
もし 404 (Not Found) とすると、「このユーザーはパスワードが違うのではなく存在しないんだ」という事がわかってしまい（このユーザーは存在していてパスワードは違う事も分かります）、セキュリティ上のリスクに繋がります。

ここで、エラーチェックは `==` を使ってはいけません。 `errors.Is` を使いましょう。 参考: https://pkg.go.dev/errors#Is

<<<@/chapter2/section1/src/0/final/code.go#post_hash

データベースに保存されているパスワードはハッシュ化されています。

ハッシュ化は不可逆な処理なので、ハッシュ化されたものから原文を調べることはできません。確認する際はもらったパスワードをハッシュ化することで行います。

これは `bcrypt.CompareHashAndPassword` が行ってくれるのでそれに乗っかりましょう。

- この関数はハッシュが一致すれば返り値が `nil` となります
- 一致しない場合、 `bcrypt.ErrMismatchedHashAndPassword` が返ってきます
- 処理中にこれ以外の問題が発生した場合は、返り値はエラー型の何かです

従って、これらのエラーの内容に応じて、 500 (Internal Server Error), 401 (Unauthorized) を返却するか、処理を続行するかを選択していきます。

<<<@/chapter2/section1/src/0/final/code.go#add_session

セッションストアに登録します。
セッションの `userName` という値にそのユーザーの名前を格納していることは覚えておきましょう。

ここまで書いたら、 `loginHandler` を使えるようにしましょう。

```go
e.POST("/login", loginHandler) // [!code ++]
e.POST("/signup", signUpHandler)
```

## userAuthMiddleware の実装

続いて、`userAuthMiddleware` を実装します。
まず、これは Handler ではなく Middleware と呼ばれます。

送られてくるリクエストは、Middleware を経由して、 Handler に流れていきます。

Middleware から次の Middleware/Handler を呼び出す際は `next(c)` と記述します。 Middleware の実装は難しいので、なんとなく理解できれば十分です。

<<<@/chapter2/section1/src/0/final/code.go#userAuthMiddleware

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
e.GET("/cities/:cityName", getCityInfoHandler) // [!code --]
e.POST("/cities", postCityHandler) // [!code --]
withAuth := e.Group("") // [!code ++]
withAuth.Use(userAuthMiddleware) // [!code ++]
withAuth.GET("/cities/:cityName", getCityInfoHandler) // [!code ++]
withAuth.POST("/cities", postCityHandler) // [!code ++]
```

これで、この章の目標である「ログインしないと利用できないようにする」が達成されました。

## getMeHandler の実装

最後に、 `getMeHandler` を実装します。叩いたときに自分の情報が返ってくるエンドポイントです。

<<<@/chapter2/section1/src/0/final/code.go#me

アクセスしているユーザーの`userName`をセッションから取得して返しています。
`userAuthMiddleware` を実行したあとなので、`c.Get("userName").(string)` によって userName を取得できます。

`withAuth.GET("/me", getMeHandler)` を忘れずに追加しましょう。
