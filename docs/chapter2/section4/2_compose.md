# Docker Composeを使う

Docker Compose を使うことで、複数のコンテナをまとめてあつかったり、コンテナを接続したりすることが簡単になります。

https://docs.docker.jp/compose/toc.html

## Docker Composeの使い方

`compose.yaml`というファイルに設定を書き、`docker compose {操作} [オプション]`のようなコマンドを入力することで使えます。

下のファイルは前ページの`docker run`コマンドを再現する設定ファイルです。

<<< @/chapter2/section4/src/docker-compose-1.yml

yaml ファイルは、インデントでオブジェクト(ひとまとまりの情報)を表現します。

`services`以下を増やしていくことで、複数のコンテナを一度に制御できます。

`build`オプションでは、どのディレクトリの Dockerfile を使ってコンテナを起動するか指定できます。また、既にあるイメージからも起動できます。

## 起動する

`compose.yaml`が存在するディレクトリで`docker compose up`とすることで、コンテナを一括で起動できます。

```sh
docker compose up
```

`Ctrl+C`でコンテナを停止できますが、削除は行われません。

`-d`オプションを追加することで、でタッチモードで起動でき、バックグラウンドでコンテナを実行できます。実際のサーバーなどで運用する場合はバックグラウンドで起動することになります。

コンテナを一括で停止・削除するためには、`down`を実行します。

```sh
docker compose down
```

## 他の Docker Compose のコマンド

`up`、`down`以外によく使う Docker Compose のコマンドを紹介します。

### `docker compose logs [オプション] [サービス名]`

コンテナからのログを出力して確認できます。サービスは`compose.yaml`の`services`以下に書かれているものを指します。サービス名を指定しない場合は全てのサービスからのログを出力します。

### `docker compose exec [オプション] {サービス名} {コマンド}`

起動しているコンテナの中でコマンドを実行できます。例えば、`docker compose exec app ls`とすると、`app`コンテナの中で`ls`コマンドを実行し、その結果を出力できます。また、コマンドで`sh`や`bash`を指定すると、コンテナの中に入って普段のターミナルのようにコマンドを実行できます。

---

他にも様々なコマンドが存在します。`docker compose`と実行するとコマンド一覧を確認できます。

## 基本演習(複数コンテナ)

下の条件を満たすように`compose.yaml`を書きかえてみましょう。

- 2 つのコンテナを起動する。
- どちらも前のページで作った `Dockerfile`を用いる。
- 1 つ目のコンテナは`greeting1`という名前で、<a href="http://localhost:3000/greeting">localhost:3000/greeting</a>にアクセスすると「こんにちは」と表示される。
- 2 つ目のコンテナは`greeting2`という名前で、<a href="http://localhost:3001/greeting">localhost:3001/greeting</a>にアクセスすると「Hello」と表示される。

試してみる前に`docker compose down`を実行して既存のコンテナを削除しておきましょう。

:::details 答え
`compose.yaml`

<<< @/chapter2/section4/src/ex1-docker-compose.yml

コンテナ側のポート番号は、2 つのコンテナで異なっていればよいです。
:::

## バインドマウントを使う

バインドマウントという機能を使ってホストマシンのファイルやフォルダをコンテナにマウントする（ホストとコンテナでファイルやフォルダを結び付ける）ことができます。

https://docs.docker.jp/storage/bind-mounts.html

nginx の設定ファイルをバインドマウントして、リバースプロキシしてみましょう。

リバースプロキシとは、サーバーへのアクセスを受け取り、サーバーアプリケーションに振り分けて中継することです。リクエスト内容に応じて異なるサーバーアプリケーションにリクエストを振り分けることができるので、負荷の軽減につながります。

DockerHub にある nginx の公式イメージを使います。
https://hub.docker.com/_/nginx/

### 設定ファイルを書く

nginx の設定ファイルとして、**`./nginx/conf.d/greeting.cnf`** を下のように書きます。

<<< @/chapter2/section4/src/greeting.conf{txt}

今はファイルの意味がわからなくてもよいですが、 nginx に来たリクエストのヘッダーの`Host`が`hello.local`のときに`greeting`コンテナの 8080 番ポートにリクエストを飛ばしているということだけ理解してください。

### Docker Compose の設定を追加する

`compose.yaml`を`nginx`を使えるように書き換えます。

<<< @/chapter2/section4/src/docker-compose-nginx.yml

`reverse_proxy`コンテナで`nginx`イメージを使って、`volumes`から設定ファイルのディレクトリをマウントしています。`volumes`は`{ホスト側のパス}:{コンテナ側のパス}`のように、コロンで区切って指定します。

`greeting`コンテナは nginx 経由のアクセス経路を作ったので、ポートの開放設定を消しています。`reverse_proxy`コンテナのコンテナ側ポートが`80`なのは、http プロトコルのデフォルトのポートが`80`であり、 nginx もそれに従っているからです。

ここまでやったらコンテナを立ち上げて`curl`コマンドでリクエストを送ってみましょう。

```sh
curl -H "Host:hello.local" http://localhost:3000/greeting
```

本来は`Host`ヘッダーにはドメイン名がのりますが、今回は都合により直接`Host`ヘッダーを指定しています。
（DNS が解決できないので、サーバーに到達できないため。`/etc/hosts` とかに書いてやるのでも良かったのですが、結構概念が難しいので今回は見送りました）

```txt
ikura-hamu@Laptop-hk:~/naro_server$ curl -H "Host:hello.local" http://localhost:3000/greeting
こんにちはikura-hamu@Laptop-hk:~/naro_server$
```

このように挨拶が表示されたら成功です。

## 基本演習(リバースプロキシ)

下の条件を満たすように設定しましょう。`compose.yaml`を編集し、 nginx の設定ファイルを追加する必要があります。

- `http://localhost:3000`で nginx がリクエストを受け付ける。
- `Host`ヘッダーが`hello1.local`であれば、`greeting1`コンテナにリクエストを飛ばし、`/greeting`に対して「こんにちは」と返ってくる。
- `Host`ヘッダーが`hello2.local`であれば、`greeting2`コンテナにリクエストを飛ばし、`/greeting`に対して「Hello」と返ってくる。

::: code-group

```sh [こんにちは]
curl -H "Host:hello1.local" http://localhost:3000/greeting
```

```sh [Hello]
curl -H "Host:hello2.local" http://localhost:3000/greeting
```

:::

:::details  答え

`naro_server/nginx/conf.d/greeting1.conf`

<<< @/chapter2/section4/src/ex2-greeting1.conf{txt}

`naro_server/nginx/conf.d/greeting2.conf`

<<< @/chapter2/section4/src/ex2-greeting2.conf{txt}

`naro_server/compose.yaml`

<<< @/chapter2/section4/src/ex2-docker-compose.yml

:::
