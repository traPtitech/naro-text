# プロジェクトのセットアップ

## 環境準備

今回の演習は、[(第一部)サーバーからデータベースを扱う](../../chapter1/section4/4_server_and_db) の状態から開始します。

もしファイルを削除してしまった場合は、以下の手順でセットアップしましょう。

1. [データベースを扱う準備](../../chapter1/section4/0_prepare) からプロジェクトをセットアップしましょう。

2. `.env` ファイルを作成し、以下のように編集しましょう。

```sh
DB_USERNAME="root"
DB_PASSWORD="password"
DB_HOSTNAME="localhost"
DB_PORT="3306"
DB_DATABASE="world"
```

3. `source .env` を実行しましょう。

4. 以下のコマンドを実行し、クレートの依存関係を追加しましょう。

```sh
$ cargo add axum anyhow serde serde_json tokio --features tokio/full,serde/derive,axum/macros
$ cargo add sqlx --features mysql,migrate,chrono,runtime-tokio
```

以上でセットアップはできているはずです。

## ファイルの分割

このまま演習を始めてしまうとファイルが長くなりすぎてしまうので、ファイルを別のモジュールとして分割します。
:::tip
パッケージとは、関連する複数のファイルをまとめる単位のことです。  
ディレクトリとパッケージは一対一に対応しています。原則的に、ディレクトリ名とパッケージ名は同じにします。    
パッケージによって、機能を分離でき、変数や関数の公開範囲を最低限にできる等沢山の恩恵が得られます。  
パッケージの外部に公開する変数や関数などのシンボルは、先頭を大文字にする必要があります。  
逆に言えば、先頭が大文字でないシンボルは、パッケージの外部からはアクセスできません。  
詳しくは以下を参照してください。
[The Rust Programming Language 日本語版 - パッケージとクレート](https://doc.rust-jp.rs/book-ja/ch07-01-packages-and-crates.html)
:::

まずは、`src` ディレクトリに
![](images/0/file-tree.png)

この画像のようなディレクトリ構造を作成しましょう。



それぞれのファイルを編集していきます。

#### handler.rs

<<<@/chapter2/section1/src/first/handler.rs{rs:line-numbers}

#### main.rs

<<<@/chapter2/section1/src/first/main.rs{rs:line-numbers}

#### repository.rs

<<<@/chapter2/section1/src/first/repository.rs{rs:line-numbers}

#### handler/country.rs

<<<@/chapter2/section1/src/first/handler/country.rs{rs:line-numbers}

#### repository/country.rs

<<<@/chapter2/section1/src/first/repository/country.rs{rs:line-numbers}

## 変更点の説明

## 準備完了

それでは、`cargo run` で実行してみましょう。

![](images/0/echo.png)

無事起動が出来たら、ターミナルで`task up`を実行してデーターベースを起動し、<a href="http://localhost:8080/cities/Tokyo">localhost:8080/cities/Tokyo</a>にアクセスして実際に動いていることを確認しましょう。

![](images/0/Tokyo.png)

上手く動いていることを確認できたら、 `Ctrl+C` で一旦止めましょう。
