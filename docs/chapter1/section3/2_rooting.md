# ルーティング

:::tip
`"hello/[ユーザー名]"`というパスのリクエストが来たときに、以下を返すサーバーを書いてみましょう。

```
Hello, [ユーザー名].
```

:::

パスに埋め込まれている情報をパスパラメータといいます。
Echo ではパスに`/:hoge`のようなコロンから始まる文字列を含めると、ハンドラに渡される`Context`の`Param`関数を使うことで取得できます。

考えうる名前全てに対して 1 つずつハンドラを設定するのは不可能なので、パスパラメーターを取得して、それをもとにレスポンスを生成します。

<<<@/chapter1/section3/src/3-1_param-server.go

[Echoガイド](https://echo.labstack.com/guide)
[Echoガイド routing](https://echo.labstack.com/guide/routing)
[Echo godoc](https://pkg.go.dev/github.com/labstack/echo/v4)
[Context godoc](https://golang.org/pkg/context/)

## リクエストパスの解析
```
/hello/pikachu?q=golang&lang=ja
```

上のようなところにリクエストが来たとき、`/hello/:username`とすることで`c.Param("username")`によって`pikachu`をとれることを知りました。

この`?q=golang&lang=ja`はクエリパラメータといって`c.QueryParam("q")`で`golang`をとれます。
クエリパラメータは順不同で`?lang=ja&q=golang`でも同じ意味になります。
`c.QueryParam("lang")`で`ja`をとれます。
このクエリパラメータは検索のリクエストを受け取るときに使うことが多いです。
例として、Google 検索だとこんな風になってます([Google検索のパラメータ(URLパラメータ)一覧](http://www13.plala.or.jp/bigdata/google.html))

[Echoでのクエリパラメータの取り方](https://echo.labstack.com/guide/request#query-parameters-1)
