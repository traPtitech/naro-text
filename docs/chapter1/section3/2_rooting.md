# ルーティング

## パスパラメーター、クエリパラメーターについて

```
http://example.com/path/param1/param2?query1=param3&query2=param4
```

この URL の、`param1`や`param2`のように、パスに埋め込まれている情報をパスパラメーターと呼びます。

また、パスパラメーターの後に`key=value&key=value&...`の形で埋め込まれる情報をクエリパラメーターと呼びます。この URL の`param3`と`param4`の部分です。

## パスパラメーターを取得してみよう

`"hello/{ユーザー名}"`というパスのリクエストが来たときに、以下を返すサーバーを書いてみましょう。

```
Hello, {ユーザー名}.
```

Echo ではパスに`/:hoge`のようなコロンから始まる文字列を含めると、ハンドラに渡される`Context`の`Param`関数を使うことで取得できます。

考えうる名前全てに対して 1 つずつハンドラを設定するのは不可能なので、パスパラメーターを取得して、それをもとにレスポンスを生成します。

<<<@/chapter1/section3/src/3-1_param-server.go

サーバーを立て直した後、<a href='http://localhost:8080/hello/pikachu' target="_blank" rel="noopener noreferrer">localhost:8080/hello/pikachu</a> にアクセスして実際に機能していることを確かめましょう。

また、URL の `pikachu` の部分を自分の名前や任意の文字列にしても動く事を確認しましょう。

`/hello/:username`とすることで`c.Param("username")`によって`pikachu`をとれることが分かりました。

### 参考
[Echoガイド](https://echo.labstack.com/guide)

[Echoガイド routing](https://echo.labstack.com/guide/routing)

[Echo godoc](https://pkg.go.dev/github.com/labstack/echo/v4)

[Context godoc](https://golang.org/pkg/context/)

## クエリパラメータを取得してみよう
```
/hello/pikachu?page=2&lang=ja
```

パスパラメーターでは`c.Param("param")`を使いましたが、クエリパラメーターは`c.QueryParam("param")`で取得できます。

クエリパラメータは順不同で`?lang=ja&page=2`でも同じ意味になります。
### 基本問題

試しに、`"hello/{ユーザー名}?lang={言語名}&page={ページ数}"`というパスのリクエストが来たときに、以下を返すサーバーを書いてみましょう。
```
Hello, {ユーザー名}.
language: {言語名}
page: {ページ数}
```

書いたらサーバーを立て直した後、<a href='http://localhost:8080/hello/pikachu?page=5&lang=ja' target="_blank" rel="noopener noreferrer">localhost:8080/hello/pikachu?page=5&lang=ja</a> にアクセスして実際に機能していることを確かめましょう。
:::details 解答
自分で書きましたか？

ここは考えれば出来るはずです。解答は自分で書いた後に、確認のために見てください。
:::details 見る
<<<@/chapter1/section3/src/3-2_query-server.go
:::

このクエリパラメータは検索のリクエストを受け取るときに使うことが多いです。

例として、Google 検索だとこんな風になってます([Google検索のパラメータ(URLパラメータ)一覧](http://www13.plala.or.jp/bigdata/google.html))。

### 参考
[Echoでのクエリパラメータの取り方](https://echo.labstack.com/guide/request#query-parameters-1)
