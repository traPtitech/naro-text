# アカウント機能の実装

## 本日の目的

`main.go` の handler の設定部分を見てみましょう。

<<<@/chapter2/section1/src/0/main_handler.go#handler{go:line-numbers}

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

### ハッシュソルトの設定

まず、`main.go`の`var`の箇所に、 `salt` を追加します。

<<<@/chapter2/section1/src/0/main_handler.go#var

ここで新しく定義した`salt`という変数は、パスワード等をハッシュ値へと変換する際に、パスワード等の末尾に付与するランダムな文字列のことです。

全員が同じハッシュ処理を使用していたとすると、他のサービスで保存されているハッシュされたパスワードと自分のサービスで保存されるものが一致したときにパスワードが特定できてしまいます（レインボーテーブル攻撃）。

このような処理をすることで、それらの可能性を防ぐことができます。

`salt` は他と一致しない独自の文字列にする必要があります。環境変数に隠しているので、環境変数を更新しましょう。

ハッシュ値は 32 バイト, 64 バイト, 128 バイトにする事が推奨されているようです。

<<<@/chapter2/section1/src/0/.env

### テーブルの作成

更に、アカウントを管理するテーブル `users` を作成します。

<<<@/chapter2/section1/src/0/handlers.go#setup_table

## signUpHandler の実装

続いて、アカウントを作成するハンドラーである `signUpHandler` を `handler.go` に実装していきましょう。

```go
func signUpHandler(c echo.Context) error {
}
```

この `signUpHandler` に以下のものを順番に実装していきます。最悪コピー&ペーストでも動くはず。

### 1. リクエストの受け取り

<<<@/chapter2/section1/src/0/final/code.go#request

まず初めに、 req 変数に requestBody の json 情報を格納します。`LoginRequestBody` 型を見れば分かる通り、ここには UserName と
Password が格納されています。

### 2. リクエストの検証

<<<@/chapter2/section1/src/0/final/code.go#valid

ここでは、UserName と Password が正しく入っているのかをチェック（バリデーションといいます）します。
入っていない場合は、与えられた入力が正しくない間違った形式なので、 400 (Bad Request) をレスポンスします。

### 3. アカウントの存在チェック

<<<@/chapter2/section1/src/0/final/code.go#check_user

`"SELECT COUNT(*) FROM users WHERE Username=?", req.Username` で、指定された UserName を持つユーザーの数を見ます。

結果は `count` 変数に格納されます。

もしすでに居た場合は、そのユーザーが存在しているので処理は受け付けず、 409 (Conflict) をレスポンスします。

### 4. パスワードのハッシュ化

ここまでは「リクエストを実行しても本当に問題がないか」を検証していました。
ユーザーはまだ存在していなくて、パスワードとユーザー名がある事まで確認できれば、リクエストを処理できます。なのでここから処理を行っていきます。

<<<@/chapter2/section1/src/0/final/code.go#hash

まずはパスワードのハッシュ化です。 **パスワードは平文で保存してはいけません！** パスワードを DB に保管するときは、必ずハッシュ化をしましょう。
ソルトは前節で説明しました。

`bcrypt`というのはいい感じにハッシュ化してくれるライブラリです。セキュリティに関わるものは自分で実装すると穴だらけになりやすいので、積極的にライブラリに頼りましょう。

### 5. ユーザーの作成

<<<@/chapter2/section1/src/0/final/code.go#add_user

`Username`, `HashedPassword` を持つ User レコードをデータベースに追加しましょう。

何かしらのエラーによって生成できなかった場合は err にその内容が詰め込まれます。
ユーザーのリクエストは問題なく、ここでエラーが発生した場合はサーバー側で何かが発生したということなので、
500 (InternalServer Error) をレスポンスします。

ここで、どんなエラーが発生したかをユーザーに直接伝えるのはセキュリティの観点から △ です。
ログで出力するだけにして、ユーザー側には 500 という情報だけ渡しましょう。

もし err がなければ、それはうまく成功したということです。 201 (Created) をレスポンスしましょう！

これで実装は終わりです。すべてを実装すると以下のようになります。

<<<@/chapter2/section1/src/0/signUpHandler.go

最後に、`main.go` に、先ほど書いたハンドラーを追加しましょう。

<<<@/chapter2/section1/src/0/handlers.go#signup

:::warning
このコードは後々の回で使用するので、エンドポイントのパス (`/signup` など) は変更しないでください！

エンドポイントの追加は問題ないので、試したい場合は新しくハンドラーを実装しましょう。
:::

ここまでできたら、実行して、Postman 等を用いて正しく実装できているかデバッグしてみましょう。

正しく API を叩いたあとに、テーブルに意図したユーザー名と、ハッシュ化されたパスワードが入っていますか？

:::details 確認に使う SQL クエリ

```sql
USE
world;
SELECT *
FROM users;
```

:::
