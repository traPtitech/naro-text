---
title: 第1回 Docker・Compose
---

# 第1回 Docker・Compose

## 目標と時間配分

イメージとコンテナの違いを説明し、nginx を Docker と Compose で動かします。
コンテナの外から接続する方法と、ファイルを渡す方法も確認します。

解説は、コンテナの仕組み、イメージの作成、通信と保存、Compose を各 15 分です。
実習は、起動 15 分、ビルド 10 分、マウント 10 分、Compose 20 分、確認と後片付け 5 分を目安にします。

## 解説

### コンテナは隔離されたプロセス

アプリを実行するには、プログラムに加えてライブラリや設定ファイルが必要です。
コンテナは、こうしたファイルをそろえ、ほかのプロセスとは分離した環境でアプリを実行する仕組みです。
Linux では、プロセスが見られる範囲を分ける namespace と、CPU やメモリの使用を管理する cgroup などを使います。

仮想マシンは OS のカーネルも含めて動かします。
一方、Linux コンテナは実行環境の Linux カーネルを共有します。
macOS や Windows の Docker Desktop では Linux の仮想マシンが用意され、その中で Linux コンテナが動きます。
隔離されていてもホストの資源は使うため、無制限に負荷をかけてよいわけではありません。

```text
Docker Desktop を使う PC
  └ Linux 仮想マシン
      ├ コンテナ A：Nginx のプロセスとファイル
      └ コンテナ B：HTTP クライアントのプロセスとファイル
          ↑ Linux カーネルを共有
```

[Docker のコンテナ解説](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-container/)も参照できます。

### イメージとコンテナ

イメージは、実行に必要なファイルと起動時の設定をまとめたものです。
コンテナは、そのイメージから作られた実行環境です。
同じイメージから複数のコンテナを作れます。

イメージは読み取り専用の層を重ねて作られます。
コンテナ内でファイルを変更すると、そのコンテナの書き込み可能な層に保存されます。
その変更は元のイメージには反映されず、コンテナを削除すると失われます。

コンテナには起動時のメインプロセスがあります。そのプロセスが終了すると、コンテナも停止します。
短いコマンドを実行するコンテナがすぐ終了するのは、正常な動作です。
Web サーバーでは、リクエストを待つプロセスを動かし続けます。

### Dockerfileとレジストリ

Dockerfile は、イメージを作る手順を記述するファイルです。
教材の例では、既存の nginx イメージに HTML を追加します。

<<< ./samples/docker/Dockerfile{dockerfile}

`FROM`は元にするイメージ、`COPY`はビルド時にコピーするファイルを指定します。
`docker build`の最後の`.`は、ビルドへ渡すファイルの範囲であるビルドコンテキストです。
`.dockerignore`で不要なファイルを除外できます。

`nginx:1.28.0-alpine`は、イメージ名`nginx`とタグ`1.28.0-alpine`の組み合わせです。
タグはバージョンなどを表す名前ですが、同じタグが別の内容を指すこともあります。
内容まで固定する場合は digest を使います。
イメージを保存・配布する場所がレジストリで、Docker Hub や GHCR などがあります。
`docker pull`で取得し、`docker push`で公開します。講習中に公開する必要はありません。

### ポートと保存先

コンテナのポートと PC のポートは区別します。
nginx がコンテナ内の`80`番で待ち受けていても、PC の`80`番へ自動で公開されるわけではありません。
`-p 127.0.0.1:8088:80`とすると、PC の`8088`番からコンテナの`80`番へ接続できます。
`127.0.0.1`を指定し、同じ PC からアクセスする実習にします。
Dockerfile の`EXPOSE`はポートの説明であり、それだけで公開する機能はありません。

```text
PC のブラウザ → 127.0.0.1:8088 → コンテナ:80 → Nginx
```

ファイルを渡したり、コンテナの削除後もデータを残したりするには、マウントを使います。

| 方法             | ファイルの場所           | 主な用途                         |
| ---------------- | ------------------------ | -------------------------------- |
| イメージへ`COPY` | イメージ内               | 配布するプログラムや静的ファイル |
| bind mount       | ホストの指定ディレクトリ | 開発中のソースや設定を渡す       |
| named volume     | Docker が管理する領域    | DB などのデータを保持する        |

マウント先に元からファイルがある場合、マウント中はそのファイルが隠れます。
コンテナを停止すること、削除すること、volume を削除することは、それぞれ別の操作です。

### Composeで構成を記述する

Compose は、サービス、イメージ、ポート、マウントなどを YAML にまとめて実行するツールです。
サービスはアプリを構成する実行単位です。Web アプリと DB のように複数のサービスを記述できます。

同じ Compose のネットワークでは、サービス名で相手を指定できます。
たとえば`web`サービスへは`http://web:80`で接続します。
クライアントのコンテナで`localhost`と書くと、クライアント自身を指します。
ホスト向けに公開した`8088`番ではなく、相手のコンテナ内の`80`番を使います。

[Compose のネットワーク解説](https://docs.docker.com/compose/how-tos/networking/)に詳しい説明があります。
起動順とアプリがリクエストを受けられる時点も区別します。
`depends_on`を書くだけでは、依存先の準備完了まで待てません。必要な場合は healthcheck などを組み合わせます。

## 実習

[事前準備](./1_prepare)を終え、`docs/docker_kubernetes/samples`で作業します。
HTML を編集するエディターと、ブラウザを用意してください。

### 1. nginxを起動する

```sh
docker run -d --name naro-web -p 127.0.0.1:8088:80 nginx:1.28.0-alpine
docker ps
curl http://127.0.0.1:8088
```

`-d`はバックグラウンドで起動する指定です。
ブラウザで`http://localhost:8088`を開くと、nginx の初期ページが表示されます。

```sh
docker logs naro-web
docker exec naro-web cat /usr/share/nginx/html/index.html
docker stop naro-web
docker ps -a
docker rm naro-web
```

`logs`ではアクセスログなど、プロセスが標準出力・標準エラー出力へ書いた内容を確認します。
`exec`は起動中のコンテナ内で別のコマンドを実行します。
`stop`後は`ps -a`に停止したコンテナが残り、`rm`で削除されます。

確認：初期ページを表示でき、停止と削除の違いを一覧で確認できましたか。

### 2. 自分のイメージを作る

```sh
cd docker
docker build -t naro-nginx:1 .
docker image ls naro-nginx
docker run -d --name naro-web -p 127.0.0.1:8088:80 naro-nginx:1
curl http://127.0.0.1:8088
```

`Hello from Docker`が返ることを確認します。
`html/index.html`の見出しを`Edited on host`に変更し、同じ URL をもう一度開いてください。
この時点では、コンテナ内の HTML は変わりません。ビルド時にコピーしたファイルを配信しているためです。

後の実習でも完成形イメージを使うため、編集した見出しを`Hello from Docker`に戻します。

```sh
docker stop naro-web
docker rm naro-web
```

確認：ホストの HTML を編集しても、ビルド済みイメージが変わらない理由を説明できますか。

### 3. bind mountでHTMLを渡す

`docker`ディレクトリで実行します。

```sh
docker run -d --name naro-web -p 127.0.0.1:8088:80 --mount "type=bind,src=$(pwd)/html,dst=/usr/share/nginx/html,readonly" nginx:1.28.0-alpine
```

`html/index.html`の見出しを`Hello from bind mount`に変え、ブラウザを再読み込みしてください。
今度はビルドなしで表示が変わります。nginx がホスト側のファイルを読んでいるためです。
`readonly`は、コンテナ側からこのマウント内へ書き込めない指定です。

```sh
docker stop naro-web
docker rm naro-web
cd ..
```

確認：イメージへのコピーとマウントの違いを、ファイルの場所から説明できますか。

### 4. Composeで動かす

`compose`へ移動し、完成形を開きます。

```sh
cd compose
```

<<< ./samples/compose/compose.yaml{yaml}

`web`が nginx、`client`が HTTP クライアントです。
`client`の`profiles: [tools]`は、通常の`up`では起動せず、必要なときだけ使う指定です。

```sh
docker compose config
docker compose up -d web
docker compose ps
docker compose run --rm client --retry 5 --retry-connrefused --retry-delay 1 http://web:80
```

`config`は、Compose が解釈した構成を表示します。
`run --rm`はクライアントを一時的に起動し、終了後に削除します。
リトライを指定すると、nginx の起動前に送ったリクエストを再試行できます。
レスポンスに`Hello from Compose`があれば、サービス名で通信できています。

ブラウザからも`http://localhost:8088`を開いてください。
`html/index.html`を編集して再読み込みし、Compose でも bind mount が働くことを確認します。

```sh
docker compose logs web
docker compose exec web cat /usr/share/nginx/html/index.html
```

確認：ブラウザの`localhost:8088`と、クライアントの`web:80`の接続先を説明できますか。

### 5. 後片付け

```sh
docker compose down
cd ..
```

コンテナと Compose のネットワークが削除されます。
一般に named volume は`down`だけでは削除されず、`down --volumes`で削除されます。
教材の Compose では bind mount を使っているため、HTML はホスト側に残ります。
`naro-nginx:1`イメージは第2回でも使うので残してください。

## 確認問題

1. コンテナを削除すると、イメージも削除されますか。
2. `-p 127.0.0.1:8088:80`の`8088`と`80`は何のポートですか。
3. Compose の`client`から`localhost:8088`へ接続すると、なぜ`web`に届きませんか。
4. DB のデータをコンテナの削除後も保持するには、何を使いますか。

::: details 解答

1. イメージは残ります。イメージの削除は別の操作です。
2. `8088`はホスト側、`80`はコンテナ側のポートです。
3. `localhost`はクライアントのコンテナ自身です。`web:80`を使います。
4. named volume などを使います。コンテナの書き込み可能な層だけに保存すると、削除時に失われます。
   :::

## よくある失敗

| 症状                                   | 確認と対応                                                         |
| -------------------------------------- | ------------------------------------------------------------------ |
| `naro-web`がすでに存在する             | `docker ps -a`で状態を確認し、実習用の同名コンテナを停止・削除する |
| ページへ接続できない                   | `docker ps`の起動状態とポート、`docker logs naro-web`を確認する    |
| マウントした HTML が見えない           | `pwd`と`html/index.html`の場所、マウント先を確認する               |
| Compose の client が起動直後に失敗する | `docker compose ps`と`logs web`で nginx が動いているか確認する     |

次は[Kubernetes の基本（1）](./3_kubernetes_basics)です。
