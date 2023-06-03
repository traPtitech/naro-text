# サーバーからデータベースを扱う

Echo を使い、データベースからデータを取得するサーバーアプリケーションを作りましょう。

<<< @/chapter1/section4/src/server.go

都市が見つかったら`200`を、見つからなかったら`404`を返しています。
Postman からリクエストを送ってみましょう。

![](assets/postman.png)

写真のように返ってきたら成功です。

## 基本問題

都市を追加する API を追加してみましょう。

:::details ヒント1
都市を追加するということはクライアントから情報を受け取る必要があります。このようなときは、どのメソッドを使えばいいでしょうか。
:::
:::details ヒント2
メソッドは`POST`を使いましょう。リクエストボディには JSON を使いましょう。どのようにすれば JSON を扱えたでしょうか。
:::

:::details 答え

- `main`関数内部

<<< @/chapter1/section4/src/practice_server.go#echo

- `postCityHandler`関数を定義

<<< @/chapter1/section4/src/practice_server.go#func

:::

## 応用問題

さまざまな API を作ってみましょう。

- 例
  - 国の情報を取得する
  - 都市をすべて取得する
  - 既にある都市や国の情報を変更する
