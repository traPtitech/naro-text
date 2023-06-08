# JSONレスポンスを返す

レスポンスとして JSON を返すようにしましょう
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

```go=
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type jsonData struct {
	Number int
	String string
	Bool bool
}

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	e.GET("/json", jsonHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func jsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool: false,
	}

	return c.JSON(http.StatusOK, &res)
}
```

![](https://md.trapti.tech/uploads/upload_c68e48ba6a4aaca07ccdac42bdd9b04f.png)

タグを追加することで構造体のフィールドに対応する、JSON のキー名を指定できます。Go の構造体のフィールドはパスカルケースですが、json のフィールドはキャメルケース / スネークケースになることが多いため、変換した方が良いですね。

```go=
type jsonData struct {
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}
```

参考: [encoding/json#Marshal](https://pkg.go.dev/encoding/json#Marshal)

# Postmanでリクエスト

Postman は HTTP リクエストを行えるアプリケーションです。

![](https://md.trapti.tech/uploads/upload_2e78634e7efe50ac34114a0305c473f8.png)

`[GET] Enter Request URL`とあるところで HTTP method と URL を指定できます。

Postman を使って GET を自分のサーバーに送ってみましょう。


::: warning
HTTP method: GET
URL: http://localhost:8080/hello
:::

![](https://md.trapti.tech/uploads/upload_585a8f0d3bbf31dd4b6c3e5a1d5d7c5e.png)
↑ポート番号が 4000 番の例。

```
Hello, World.
```
と表示されれば成功です。

次に POST リクエストを使ってみましょう。

POST ではサーバーにデータを送ることができます。

**PostmanでBodyタブを選択
ラジオボタンの`raw`を選択
右に出てくるプルダウンから`JSON(application/json)`を選択します**

POST で渡せるデータの型は複数あり、上記の操作で JSON を使うということを明示しています。

以下のように自分の traQ ID を POST してみましょう。

::: warning
HTTP method: POST
URL: https://eo6mn2b7rlihmgg.m.pipedream.net
:::

```json
{
    "traq_id": "asari"
}
```

![](https://md.trapti.tech/uploads/upload_c4d30574a4385ca97f15a3814700815f.png)

`traq_id`が`asari`の例。
```
{
    ...
    "body": {
        "traq_id": "asari"
    }
}
```
が返ってきたら成功です。

<!--
inspectある?
から自分のtraQ IDがあるか確認してみましょう
-->

## 自分のサーバーでPOSTを受け取る

### 基本問題
POST で JSON を受け取って、内容をそのまま返すサーバーを作ってみましょう。

`e.GET`と同じように、`e.POST`と書くことで POST を受け取ることができます。
POST のハンドラは、受け取りたい JSON を示す空の変数を先に用意し、`Context`の`Bind`に渡すことで送られてきたデータを取り出すことができます。
データが存在しなかったりした場合には、返り値の`err`にエラーが入ります。
逆にエラーがないときは`err`に`nil`が返ってくるので、`if`で条件分岐をします。

```go=
package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)

type jsonData struct {
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}

func main() {
	e := echo.New()

	e.POST("/post", postHandler)
	e.Logger.Fatal(e.Start(":<ポート番号>"))
}

func postHandler (c echo.Context) error {
	data := &jsonData{}
	err := c.Bind(data)

	if err != nil { // エラーが発生した際
		// fmt.Sprintf("%+v", data): dataをstringに変換
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", data))
	}
	return c.JSON(http.StatusOK, data)
}
```

Postman を使って実際に受け取れている / 送り返せているか確認してみましょう。
※omitempty を指定していると false, 0, 空文字("")は返ってきません。
