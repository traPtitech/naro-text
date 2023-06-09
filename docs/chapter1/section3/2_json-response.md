# JSONレスポンスを返す

先ほど作った`hello-server.go`に、レスポンスとして JSON を返すエンドポイントを追加しましょう。
:::warning
JSON について分からない人は
[JSONってなにもの？ | Think IT（シンクイット）](https://thinkit.co.jp/article/70/1)
:::

:::warning
Go 言語の構造体についてわからない人は
https://go-tour-jp.appspot.com/moretypes/2
を見ると良いです。
:::

JSON をレスポンスとして返すためには、`c.JSON`メソッドに構造体を渡します。

<<<@/chapter1/section3/src/2-1_json-server.go

http://localhost:8080/json

![](assets/json_server.png)

タグを追加することで構造体のフィールドに対応する、JSON のキー名を指定できます。Go の構造体のフィールドはパスカルケースですが、JSON のフィールドは普通キャメルケース / スネークケースであるため、構造体を以下のように書き換えましょう。

```go=
type jsonData struct {
    # Numner -> number (omitemptyは、ゼロ値の場合はそのフィールドを出力しないという意味)
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}
```

参考: [encoding/json#Marshal](https://pkg.go.dev/encoding/json#Marshal)

## Postmanでリクエストしてみよう

### Postmanのインストール
Postman は GUI で HTTP リクエストを行えるアプリケーションです。

ダウンロードページ→ https://www.postman.com/downloads/

[Postman | API Development Environment](https://www.getpostman.com/)

![](assets/postman.png)

`[GET] Enter Request URL`とあるところで HTTP method と URL を指定できます。

Postman を使って GET を自分のサーバーに送ってみましょう。


::: tip
HTTP method: GET

URL: http://localhost:8080/hello
:::

![](assets/postman-hello.png)

```
Hello, World.
```
と表示されれば成功です。

## 次に POST リクエストを使ってみましょう

POST ではサーバーにデータを送ることができます。

**PostmanでBodyタブを選択
ラジオボタンの`raw`を選択
右に出てくるプルダウンから`JSON(application/json)`を選択します**

POST で渡せるデータの型は複数あり、上記の操作で JSON を使うということを明示しています。

以下のように自分の traQ ID を POST してみましょう。

::: tip
HTTP method: POST

URL: https://eo6mn2b7rlihmgg.m.pipedream.net

```json
{
    "traq_id": "pikachu"
}
```
:::

![](assets/postman-post.png)

![](assets/postman-response.png)
`traq_id`が`pikachu`の例だと、上の画像のように、以下のような JSON が返ってきます。
```
{
    ...
    "body": {
        "traq_id": "pikachu"
    }
}
```

<!--
inspectある?
から自分のtraQ IDがあるか確認してみましょう
-->

## 自分のサーバーでPOSTを受け取ってみよう

POST で JSON を受け取って、内容をそのまま返すサーバーを作ってみます。

`e.GET`と同じように、`e.POST`と書くことで POST を受け取ることができます。
POST のハンドラは、受け取りたい JSON を示す空の変数を先に用意し、`Context`の`Bind`に渡すことで送られてきたデータを取り出すことができます。
データが存在しなかったりした場合には、返り値の`err`にエラーが入ります。
逆にエラーがないときは`err`に`nil`が返ってくるので、`if`で条件分岐をします。

<<< @/chapter1/section3/src/2-2_echo-server.go

Postman を使って実際に受け取れている / 送り返せているか確認してみましょう。

※omitempty を指定していると false, 0, 空文字("")は返ってきません。

![](assets/postman-echo.png)
