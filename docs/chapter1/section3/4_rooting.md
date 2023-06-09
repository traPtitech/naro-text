# ルーティング

:::tip
`"hello/[ユーザー名]"`というパスのリクエストが来たときに

```
Hello, [ユーザー名].
```

と返すサーバーを書いてみましょう
:::

パスに情報が埋め込まれているのをパスパラメータといいます
Echoではパスに`/:hoge`のようなコロンから始まる文字列を含めると、ハンドラに渡される`Context`の`Param`関数を使うことで取得できます

考えられる名前全てを書くことは不可能なので、パスパラメーターを取得して、それをもとにレスポンスを生成します

<<<@/chapter1/section3/src/3-1_param-server.go

[Echoガイド](https://echo.labstack.com/guide)
[Echoガイド routing](https://echo.labstack.com/guide/routing)
[Echo godoc](https://pkg.go.dev/github.com/labstack/echo/v4)
[Context godoc](https://golang.org/pkg/context/)

## リクエストパスの解析
```
/hello/pikachu?q=golang&lang=ja
```

上のようなところにリクエストが来たとき、`/hello/:username`とすることで`c.Param("username")`で`pikachu`をとれることを知りました

この`?q=golang&lang=ja`はクエリパラメータといって`c.QueryParam("q")`で`golang`をとれます
クエリパラメータは順不同で`?lang=ja&q=golang`でも同じ意味になります
`c.QueryParam("lang")`で`ja`をとれます
このクエリパラメータは検索のリクエストを受け取るときに使うことが多いです
例として、Google検索だとこんな風になってます([Google検索のパラメータ(URLパラメータ)一覧](http://www13.plala.or.jp/bigdata/google.html))

[Echoでのクエリパラメータの取り方](https://echo.labstack.com/guide/request#query-parameters-1)