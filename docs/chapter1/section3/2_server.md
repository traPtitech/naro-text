# サーバーアプリケーションを作る

## ファイルの作成
今回は、Go言語と、Go言語の有名なwebフレームワークである[Echo](https://echo.labstack.com/)を使ってサーバーアプリケーションを作っていきます。

`/home/userX/go/src/hello-server`というディレクトリを作成し、そのディレクトリを開きます。
ディレクトリの中に`main.go`を作成し、以下のプログラムを書き込みます。

```go=
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World.\n")
    })
    
    e.Logger.Fatal(e.Start(":<ポート番号>")) 
    // ここを前述の通り自分のポートにすること(例: e.Start(":10100"))
}
```

Echoは、[Go言語の標準ライブラリ](https://pkg.go.dev/std)に入っていない外部ライブラリなので、外部からインストールしなければならないのですが、それを自動でやってくれる[Go module](https://go.dev/doc/tutorial/create-module)という便利な機能があるので、使いましょう。
```
#Go moduleを初期化して、足りない物をインストールし、使われてない物を削除する。
$ go mod init naro-server
$ go mod tidy
```
```
#先ほど書いたファイルを実行
$ go run /home/userX/go/src/hello-server/main.go
```

以下のような画面が出れば起動できています。
止めるときは`ctrl+c`で終了できます。というか止めないと次に起動するときにポート番号を変えないとエラーが出てしまうので使い終わったら止めるようにしましょう。

![](https://md.trapti.tech/uploads/upload_1fd461f85490ef5014ebbafdfa430a0a.png)

:::warning
作ったディレクトリやファイルの名前が違うと、上手く実行できない場合があります。
適宜読み替えてください。
:::

:::warning
## ターミナルの開き方
ツールバー > Terminal > New Terminalでその時開いているディレクトリでターミナルが開きます。
もしくは`Ctrl` + `@`でも(Windowsの場合)。
:::