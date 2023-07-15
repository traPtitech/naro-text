# プロジェクトのセットアップ

## 環境準備

今回の演習は、[(第一部)サーバーからデータベースを扱う](../../chapter1/section4/4_server_and_db) の状態から開始します。

もしファイルを削除してしまった場合は、以下の手順でセットアップしましょう。

1. [データベースを扱う準備](../../chapter1/section4/0_prepare) からプロジェクトをセットアップしましょう。

2. `.env` ファイルを作成し、以下のように編集しましょう。

```sh
export DB_USERNAME="root"
export DB_PASSWORD="password"
export DB_HOSTNAME="localhost"
export DB_PORT="3306"
export DB_DATABASE="world"
```

3. `source .env` を実行しましょう。

4. `go mod tidy` を実行しましょう。

以上でセットアップはできているはずです。

## ファイルの分割

このまま演習を初めてしまうとファイルが長くなりすぎてしまうので、ファイルを分割します。
各エンドポイントでの処理はハンドラーと呼ばれますが、それを `handler.go` に移動してみましょう。

## main.go の作成

main.go を以下のようにしましょう。

<<<@/chapter2/section1/src/0/main.go{go:line-numbers}

ファイルを編集したら、`go mod tidy` を実行しましょう。

## handler.go の作成

1. 同じディレクトリに新しく `handler.go` というファイルを作成する
2. `handler.go` を以下のように記述する

<<<@/chapter2/section1/src/0/handler.go{go:line-numbers}

ファイルを編集したら、`go mod tidy` を実行しましょう。

## 準備完了

今回は `main.go` 以外に `handler.go` も存在するので、どちらも指定して `go run *.go` を実行しましょう。

![](images/0/echo.png)

無事起動が確認できたら、 `Ctrl+C` で一旦止めましょう。