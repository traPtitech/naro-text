# Rustでデータベースを扱う

ここからは Rust でプログラムを書いてデータベースを扱っていきます。`task up`を実行してデータベースが立ち上がっていることを確認してください。
まずは VSCode で先ほどクローンしてきたリポジトリを開きましょう。画像のようなファイルが入っているはずです。 main.rs を開いてください。
![](images/files.png)

## データベースに接続する

### Rustのプログラムを書く

サンプルのプログラムが書いてありますが、データベースと接続できるように書き換えます。
Rust でデータベースに接続するためのライブラリは様々ありますが、今回は SQL 文を書く SQLx を使います。

- 参考
  - [launchbadge/sqlx: 🧰 The Rust SQL Toolkit](https://github.com/launchbadge/sqlx)
  - [sqlx - Rust](https://docs.rs/sqlx/latest/sqlx/)

<<< @/chapter1/section4/src/connect_db.rs{rust:line-numbers}

`get_option` 関数により、データベースに接続するための設定を構成し、`MySqlPool::connect_with` でデータベースに接続しています。
`env::var` により、環境変数を読み込んでいます。環境変数を使うことで、プログラムの動作を変えることなく、データベースの接続情報を変更できます。

### 環境変数を設定する

`.env`というファイルを作り、以下の内容を書いてください。

```sh
export DB_USERNAME="root"
export DB_PASSWORD="password"
export DB_HOSTNAME="localhost"
export DB_PORT="3306"
export DB_DATABASE="world"
```

今回は手元で動いているデータベースを使うのでパスワードなどが知られても問題ありませんが、実際には環境変数など外部の人に知られたくない、GitHub などに上げたくないファイルもあります。そのような場合は、`.gitignore`というファイルを使うことで特定のファイルやフォルダを Git 管理の対象外にできます。`.gitignore`ファイルの最後に`.env`を追記しましょう。

```txt
...
# Added by cargo

/target
.env // [!code ++]
```

`.gitignore`ファイルは便利ですが、既に Git の追跡対象になっているファイルを書いても追跡対象から外れないので注意しましょう。

https://docs.github.com/ja/get-started/getting-started-with-git/ignoring-files

最後に環境変数ファイル`.env`を環境変数として読み込むために、ターミナルで`source .env`を実行してください。

:::warning

```sh
$ source .env
```

このコマンドによって読み込んだ環境変数は、コマンドを入力したターミナルを終了すると消えてしまいます。また、コマンドを入力したターミナル以外では環境変数として読み込まれません。新しくターミナルを開きなおした場合などは、もう一度実行してください。
:::

### 実行する

```sh
$ cargo run
```

出力はこのようになります。

```txt
connected
Tokyoの人口は7980230人です
```

`main.rs`を解説してきます。

<<< @/chapter1/section4/src/connect_db.rs#city

`#[derive(sqlx::FromRow)]`を使うことで、SQL 文で取得したレコードを構造体へ変換できるようになります。`#[sqlx(rename_all = "PascalCase")]` によって、データベースのカラム名が`PascalCase`に変換されます。また、`#[sqlx(rename = "ID")]` によって、`ID`というカラム名を`id`というフィールドに変換しています。

<<< @/chapter1/section4/src/connect_db.rs#get

`sqlx::query_as` により、SQL 文を実行して結果を構造体に変換しています。SQL 文中の `?`  に対して、`bind` で値を順番に結び付けることができます。
`fetch_one` により 1 つのレコードを取得しています。

### 基本問題

```sh
$ cargo run {都市の名前}
```

と入力して、同様に人口を表示するようにしましょう。

ヒント：[コマンドライン引数を受け付ける - The Rust Programming Language 日本語版](https://doc.rust-jp.rs/book-ja/ch12-01-accepting-command-line-arguments.html)

:::details 答え
<<< @/chapter1/section4/src/practice_basic1.rs
:::

### 応用問題

基本問題 1 と同様に都市を入力したとき、その都市の人口がその国の人口の何％かを表示してみましょう。

ヒント： 1 回のクエリでも取得できますが、2 回に分けた方が楽に考えられます。

:::details 答え
<<< @/chapter1/section4/src/practice_advanced.rs
:::

## 複数レコードを取得する

`fetch_one`関数の代わりに`fetch_all`関数を使い、第 1 引数を配列のポインタに変えると、複数レコードを取得できます。`main.rs`の`main`関数を以下のように書き換えて実行してみましょう。

<<< @/chapter1/section4/src/select.rs#main{rs:line-numbers}
以下のように日本の都市一覧を取得できます。

```txt
connected
日本の都市一覧
都市名: Tokyo, 人口: 7980230
都市名: Jokohama [Yokohama], 人口: 3339594
都市名: Osaka, 人口: 2595674
都市名: Nagoya, 人口: 2154376
都市名: Sapporo, 人口: 1790886
都市名: Kioto, 人口: 1461974
...省略
```

日本の都市一覧を取得出来たら、スクリーンショットを講習会用チャンネルに投稿しましょう。

## レコードを書き換える

`INSERT`や`UPDATE`、`DELETE`を実行したい場合は、`query`関数を使うことができます。

```rs
let result = sqlx::query("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)")
    .bind(city.name)
    .bind(city.country_code)
    .bind(city.district)
    .bind(city.population)
    .execute(&pool)
    .await?;
```

例えば`INSERT`ならば、このように使うことができます。return で返ってくる`result`には、`INSERT`で何件のレコードが追加されたかなどの情報が入っています。

:::info 詳しく知りたい人向け
**なぜSQL文で「`?`」を使うのか**

sqlx で変数を含む SQL を使いたいときは「`?`」を使わなくてはいけません。これはセキュリティ上の問題です。例として、国のコードからその国の都市の情報一覧を取得することを考えましょう。`format!`を使って SQL 文を作成すると以下のようになります。

```rs
sqlx::query_as::<_, City>(
    format!("SELECT * FROM city WHERE CountryCode = '{}'", code).as_str(),
)
```

`code`に入っている値がただの国名コードなら問題はないのですが、`JPN' OR 'A' = 'A`という値が入っていたらどうなるでしょうか。データベースで実行されるとき、SQL 文は下のようになります。

```sql
SELECT * FROM city WHERE CountryCode = 'JPN' OR 'A' = 'A'
```

`OR`でつなげた条件文のうち、「`'A' = 'A'`」は常に成り立つので、`WHERE`句の条件は常に真です。よって、この SQL を実行すると、作成者が意図しない方法で全ての都市が取得できてしまいます。このような攻撃は「SQL インジェクション」と呼ばれます。

sqlx ではこれを防ぐために`?`を使うことができ、SQL 文が意図しない動きをしないようになっています。
:::
