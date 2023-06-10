# サーバーアプリケーションを作ってみよう

## ファイルの作成

今回は、Go と、Go の有名な web フレームワークである [Echo](https://echo.labstack.com/) を使ってサーバーアプリケーションを作っていきます。

`~/develop/hello-server`というディレクトリを作成し、そのディレクトリを開きます。
```bash
# ディレクトリ ~/develop/hello-server を作成し、そのディレクトリを開く。

$ mkdir -p ~/develop/hello-server
$ cd ~/develop/hello-server
$ code .
```
:::tip
先ほどのコマンドの`mkdir`は、`make directories`の略で`-p`というオプションを付けると、階層の深いディレクトリを 1 回で作ることができます。このような説明は、man コマンド(マニュアルの man )を使うことで調べることができます。ググる前に使うと良いです。(q キーを押すことで抜けられます。)
```bash
$ man mkdir
```
![](assets/mannual.png)
:::

:::warning
作ったディレクトリやファイルの名前が違うと、上手く実行できない場合があります。
:::

作ったディレクトリの中に`main.go`を作成し、以下のプログラムを書き込みます。

<<<@/chapter1/section3/src/1-1_hello-server.go

Echo は、[Go の標準ライブラリ](https://pkg.go.dev/std)に入っていない外部ライブラリなので、外部からダウンロードしなければなりません。しかし、Go にはそれを自動でやってくれる [Go module](https://go.dev/doc/tutorial/create-module) という便利な機能があるので使ってみましょう。以下を VSCode 内のターミナルで実行してください。(他のターミナルでも可)

:::tip
**ターミナルの開き方**

ツールバー > Terminal > New Terminal でその時開いているディレクトリでターミナルが開きます。
もしくは`Ctrl` + `@`でも。
:::

```bash
# Go module を初期化して、足りない物をインストールし、使われてない物を削除する。

$ go mod init develop
$ go mod tidy
```

:::warning
本来この `develop` の所にはリポジトリ名を入れることが多いです。詳しくは[公式ドキュメント](https://go.dev/doc/modules/managing-dependencies#naming_module)を参照してください。
:::

続けて、`main.go` を実行してサーバーを立てましょう。
```bash
# 先ほど書いたファイルを実行して、サーバーを立てる

$ go run main.go
```

以下のような画面が出れば起動できています。

止めるときはターミナル上にて`Ctrl+C`で終了できます。
:::warning
**止めないと次に起動するときにポート番号を変えないとエラーが出てしまうので、使い終わったら`Ctrl+C`で止めるようにしましょう。**
:::

![](assets/hello_server.png)


## アクセスしてみる

まずはコマンドライン(ローカル)でサーバーにアクセスしてみましょう。

コマンドラインでサーバーにアクセスするには、[curl](https://curl.se/)というコマンドを使います。

ターミナルパネルの上にあるツールバーのプラスボタンを押すと、新たにターミナルを開くことができます。
![](assets/plus_button.png)

新しくターミナルを開いて、以下のコマンドを実行してみましょう。

```bash
$ curl localhost:8080/hello
```

すると、レスポンスとして Hello, World が返ってきていることがわかります。
![](assets/hello_server_success.png)

## 更に詳しくリクエストを見る

curl コマンドのオプションとして、リクエストなどの情報を詳しく見る`-vvv`があります。

```bash
$ curl localhost:8080/hello -vvv
```

とすると
![](assets/hello_server_detail.png)

先程座学でやったような、リクエスト・レスポンスが送られていることがわかります。

## ブラウザからアクセスする

localhost は自分自身を表すドメインなので、自分のブラウザからのみアクセスが可能です。

ブラウザで、http://localhost:8080/hello にアクセスしてみましょう。
![](assets/hello_server_localhost.png)

## 基本問題
エンドポイントとして自分の traQ ID のものを生やして自己紹介を返すようにしてみましょう。

`main.go`に`/{自分の traQ ID}`処理を追加して作ってください。

:::warning
この章では、この`main.go`に処理を追加していきます。

以降のコードではすでに作ったエンドポイントを省略していますが、作ったエンドポイントは消さずに、新しいエンドポイントを追加していくようにしてください。
:::

作り終わったら、変更を反映させるために、`go run main.go`を実行したターミナル上で`Ctrl+C`を押してサーバーを止めた後、また`go run main.go`してサーバーを立て直しましょう。

今後`main.go`を書き換えたらこの工程を行うようにして下さい。

サーバーの立て直しができたら、ブラウザで`http://localhost:8080/{自分の traQ ID}` にアクセスするか、以下のコマンドを実行して、上手く出来ていることを確認しましょう。
```bash
$ curl http://localhost:8080/{自分の traQ ID}
```

例:
![](assets/hello_server_me.png)

ここまで完成したら、講習会の実況用チャンネルに↑のようなスクリーンショットを投稿しましょう。
