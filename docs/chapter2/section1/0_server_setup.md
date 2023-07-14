# main.go
まず先に各関数の解説をします。`main.go`全体は後ろに載せています。
:::tip 
下に解説が書いてあります
:::

依存ライブラリをインストールするために以下のコマンドを実行します。

```
$ go get -u github.com/labstack/echo-contrib/session
$ go get -u github.com/srinathgs/mysqlstore
```


## 解説
### main関数

<<<@/chapter2/section1/src/0/main.go#setup_session{go:line-numbers}

セッションストアを設定しています。セッションとは、今来た人が次来たとき、同じ人であることを確認するための仕組みです。それらを覚えておくための場所をデータベース上に設定しています。
<!-- TODO: ここ行数が違う -->
54 行目の`e.Use(session.Middleware(store))`が echo にこのセッションストアを使ってと言っている部分です。

それ以降では echo についての設定をしています。具体的にはハンドラの登録やサーバーの起動などを行っています。このあたりは第一部で説明しているはずなので割愛します。

### postSignUpHandler関数
これは User 登録を行なうための関数です。

<<<@/chapter2/section1/src/0/main.go#request{go:line-numbers}

req にリクエスト情報を入れています。ここには UserName と Password が格納されているはずです。

<<<@/chapter2/section1/src/0/main.go#valid{go:line-numbers}

ここでは、本当に UserName と Password が入っているのかをチェックしています。入っていなければ不正なリクエストなので、400(Bad Request)を返しています。

<<<@/chapter2/section1/src/0/main.go#hash{go:line-numbers}

基本的に Password を平文で保存しておくのは危険です！　従って、パスワードを DB に格納するときはハッシュ化を行ってから格納してください。本来はハッシュ化の有効性を高めるためにハッシュソルトというものを使用するべきなのですが、今回は割愛します。興味がある人は調べてみてください。

`bcrypt`というのはハッシュ化をいい感じにやってくれるライブラリです。それを使ってパスワードをハッシュ化しています。

<<<@/chapter2/section1/src/0/main.go#check_user{go:line-numbers}

`db.Get(&count, "SELECT COUNT(*) FROM users WHERE Username=?", req.Username)`という部分は、db に`req.Username`という名前の User は何人いますか？　という問い合わせをしています。

その結果は`count`に格納されています。同名のユーザーがいたら困るので、そういう場合はその名前の User はもういるからダメだよというレスポンスを返します。



<<<@/chapter2/section1/src/0/main.go#add_user{go:line-numbers}


`db.Exec`はクエリを実行する関数です。ここでは Username,HashedPassword を持つ User を生成しようとしてます。

何かしらのエラーによって生成できなかった場合は err にその内容が詰め込まれます。上までの処理から理論上は User を生成できるはずなので、ここで何かエラーが出たとするとそれはサーバー側の問題になります。従ってここで返しているエラーコードは 500 になっています。

### postLoginHandler関数

<<<@/chapter2/section1/src/0/main.go#post_req{go:line-numbers}

req の代入のところは SignUp のところと同じです。

下の部分ではリクエストで送られてきた UserName と Password を持つユーザーは存在するのか？　という問い合わせをしています。存在した場合は`user`にそのユーザーの情報が入ります。

<<<@/chapter2/section1/src/0/main.go#post_hash{go:line-numbers}


SignUp の方にも書きましたが、パスワードを平文で保存するのは良くないということでハッシュ化されています。

送られてきたパスワードが正しいパスワードなら、ユーザー生成時に行った計算と同様のものを行うことで同じハッシュ文字列を得ることができ、それによってこのパスワードが正しいかどうかを確認できます。その処理を行っています。

この関数はハッシュが一致するときに限り`err`が`nil`で返ってきます。一致しない場合のエラーは`bcrypt.ErrMismatchedHashAndPassword`というものです。

従って、これの場合はパスワードが違うよというレスポンスを返し、それ以外のエラーの場合は 500 を返しています。


<<<@/chapter2/section1/src/0/main.go#add_session{go:line-numbers}


セッションに登録する処理です。

ここではセッションにログインした人を登録しているんだなあということが分かれば大丈夫です。

### checklogin関数
まず、この関数は handler ではありません。handler 関数を返す関数です。(middleware と呼ばれています) リクエスト→middleware→handler という順序で処理されます。

middleware から次の handler を呼び出すには`next(c)`と書きます。

このミドルウェアはリクエストを送ったユーザーがログインしているのかをチェックし、ログインしているなら echo の Context にそのユーザーの UserName を登録します。

<<<@/chapter2/section1/src/0/main.go#get_session{go:line-numbers}

セッションを取得しています。

本当はリクエストヘッダを見ることでどのセッションを取り出すかを決めています(セッションは各ユーザーに存在するので)

<<<@/chapter2/section1/src/0/main.go#check_session{go:line-numbers}


Login 時の処理を思い出すと、セッションには"userName"をキーとしてユーザーの名前が登録されていました。

従って、ここに名前が入っているならば、今リクエストを送った人は過去にログインをした人と同じということがわかります。逆に、何も入っていなければリクエストを送った人はログインをしていません。

これを利用して、ログインしていない場合には後続に処理を渡すことをせず途中で処理を止めています。

### getWhoAmIHandler関数

<<<@/chapter2/section1/src/0/main.go#whoami{go:line-numbers}

セッションからアクセスしているユーザーの`userName`を取得して返しています。
ここにアクセスすれば自分がどのアカウントでアクセスしてるか知ることができます。

## 完成形

<details>

<<<@/chapter2/section1/src/0/main.go{go:line-numbers}

</details>

# 検証
:::warning
全て Postman での検証です
`go run main.go`でサーバーを起動した状態で行ってください
:::

http://localhost:8080/cities/Tokyo へ
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


もう一度 http://localhost:8080/cities/Tokyo にアクセスすると正常に API が取れるようになりました。
![](https://md.trapti.tech/uploads/upload_59c6c86e127d982f511946d2a183d0a6.png)

ここで、作成されたユーザーがデータベースに保存されていることを確認してみましょう。
`mysql > SELECT * FROM users;`
![](https://md.trap.jp/uploads/upload_f713b7da16df6729729a25ca2b5a6816.png)
ユーザー名とハッシュ化されたパスワードが確認できますね。
![](https://md.trap.jp/uploads/upload_7f007d73bd0ff508dff12246546b1a5b.png)
ちょっと分かりにくい表示ですが、セッションもしっかり確認できます。


# 難
TODO リストのサーバーとしての API を考えて作ってみましょう。


:::tip
ここまでで Web サービスを作るコードの知識として必要な要素は全て網羅したつもりです。
みなさんはもう何でも作れるわけです!!!!

今回の目標は Twitter クローンの作成なので、早速作り始めても OK です!!
<!-- リポジトリ名変える -->
もし作り始める場合は https://github.com/tohutohu/naro-portal から fork して作業をして、Pull Request を出してもらえると講師や TA やその他暇な人が勝手にレビューをします！　全部完成していなくてもここまでできたので見てくださいとかでも構いません！
ぜひ皆さん使ってください！

fork とか Pull Request とかがわからない人は TA に言ってください(3 分くらいで終わるので)
:::
