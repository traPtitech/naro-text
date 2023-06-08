# サーバーアプリケーションを作る

## ファイルの作成

今回は、Go 言語と、Go 言語の有名な web フレームワークである [Echo](https://echo.labstack.com/) を使ってサーバーアプリケーションを作っていきます。

`~/naro-server/hello-server`というディレクトリを作成し、そのディレクトリを開きます。
```
#mkdirは-pを付けると、階層の深いディレクトリを1回で作れる
mkdir -p ~/naro-server/hello-server
cd ~/naro-server/hello-server
code .
```
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

	e.Logger.Fatal(e.Start(":8080"))
}

```

Echo は、[Go言語の標準ライブラリ](https://pkg.go.dev/std)に入っていない外部ライブラリなので、外部からインストールしなければなりません。しかし、Go 言語にはそれを自動でやってくれる [Go module](https://go.dev/doc/tutorial/create-module) という便利な機能があるので使ってみましょう。以下を VSCode 内のターミナルで実行してください。

:::tip
<span style="font-size: 150%;font-weight: bold;"> ターミナルの開き方 </span>

ツールバー > Terminal > New Terminal でその時開いているディレクトリでターミナルが開きます。
もしくは`Ctrl` + `@`でも。
:::

```
#Go moduleを初期化して、足りない物をインストールし、使われてない物を削除する。
go mod init naro-server
go mod tidy
```

:::warning
本来この naro-server の所にはリポジトリ名を入れることが多いです。詳しくは[公式ドキュメント](https://go.dev/doc/modules/managing-dependencies#naming_module)を参照してください。
:::

続けて、main.go を実行してサーバーを立てましょう。
```
#先ほど書いたファイルを実行して、サーバーを立てる
go run main.go
```

以下のような画面が出れば起動できています。
止めるときは`ctrl+c`で終了できます。というか止めないと次に起動するときにポート番号を変えないとエラーが出てしまうので使い終わったら止めるようにしましょう。

![](assets/hello_server.png)

:::tip
作ったディレクトリやファイルの名前が違うと、上手く実行できない場合があります。
適宜読み替えてください。
:::

# アクセスしてみる

まずはコマンドライン(ローカル)でサーバーにアクセスしてみましょう。
コマンドラインでサーバーにアクセスするには、[curl](https://curl.se/)というコマンドを使います。

ターミナルパネルの上にあるツールバーのプラスボタンを押すと、新たにターミナルを開くことができます。
![](assets/plus_button.png)

新しくターミナルを開いて、以下のコマンドを実行してみましょう。

```
curl localhost:8080/hello
```

すると、レスポンスとして Hello, World が返ってきていることがわかります。
![](assets/hello_server_success.png)

## 更に詳しくリクエストを見る

curl コマンドのオプションとして、リクエストなどの情報を詳しく見る`-vvv`があります。

```
curl localhost:8080/hello -vvv
```

とすると
![](assets/hello_server_detail.png)

先程座学でやったような、リクエスト・レスポンスが送られていることがわかります。

# ブラウザからアクセスする

localhost は自分自身を表すドメインなので、自分のブラウザからのみアクセスが可能です。
ブラウザで、`http://localhost:8080/hello`にアクセスしてみましょう。
![](assets/hello_server_localhost.png)

### 基本問題
エンドポイントとして自分の traQ ID のものを生やして自己紹介を返すようにしてみましょう。

例:
![](assets/hello_server_me.png)

完成したら webApp/jikkyo チャンネルにスクリーンショットを投稿してください。
