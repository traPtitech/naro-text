---
title: 事前準備
---

# 事前準備

インストールとダウンロードは、講習会の前に済ませてください。
必要なのは Docker、Compose、kubectl、kind、Git、curl と、コードを編集するエディターです。

## 作業環境

コマンドは Bash で実行します。
Windows では WSL2 の Ubuntu、macOS では Bash または zsh、Linux では Bash を使ってください。
PowerShell へそのまま貼り付けることは想定していません。

Docker に割り当てる目安は CPU 4 コア、メモリ 6 GiB 以上、空きディスク 10 GB 以上です。
これは講習用の目安で、Kubernetes 自体の最低要件ではありません。
VPN やプロキシを使っている場合は、イメージを取得できるか事前に確認してください。

### Windows

[WSL のインストール手順](https://learn.microsoft.com/ja-jp/windows/wsl/install)に従って WSL2 を用意します。
[Docker Desktop](https://docs.docker.com/desktop/setup/install/windows-install/)をインストールし、WSL integration で使用する Ubuntu を有効にしてください。
Git や kubectl、kind と、教材のファイルは WSL 側にそろえます。
作業場所は`/mnt/c`配下より、WSL 内のホームディレクトリ配下を勧めます。

### macOS

[Docker Desktop](https://docs.docker.com/desktop/setup/install/mac-install/)をインストールして起動します。
Apple Silicon では arm64 用、Intel Mac では amd64 用のツールを選んでください。

### Linux

[Docker Engine](https://docs.docker.com/engine/install/)と[Compose plugin](https://docs.docker.com/compose/install/linux/)をインストールします。
一般ユーザーで`docker`を使う設定は、[インストール後の手順](https://docs.docker.com/engine/install/linux-postinstall/)を参照してください。
Docker の操作権限はホストを管理できるほど強いため、権限設定の説明も確認します。

## ツールの確認

Docker Desktop の Kubernetes 機能は使いません。kind が Docker 上に別のクラスタを作ります。
kind は Kubernetes in Docker の略です。

[kind の公式手順](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)と
[kubectl の公式手順](https://kubernetes.io/docs/tasks/tools/)から、OS と CPU に合うものを導入してください。
教材と同じ構成にする場合は、kind v0.30.0 と kubectl v1.34 系を使用します。
kubectl は接続する Kubernetes と同じ minor version、または前後 1 つの minor version を使います。

```sh
docker version
docker compose version
kubectl version --client
kind version
git --version
curl --version
```

`docker version`に Client と Server の両方が表示されれば、Docker の実行環境に接続できています。
Server が表示されない場合は、Docker Desktop または Docker Engine を起動してください。
Compose は`docker compose`というコマンドを使います。

## 教材を取得する

```sh
git clone https://github.com/traPtitech/naro-text.git
cd naro-text/docs/docker_kubernetes/samples
```

公開前の教材を使う場合は、clone 後に講師が指定するブランチへ切り替えてください。
この後のコマンドは、特に指定がなければ`samples`を基準に実行します。

| ディレクトリ        | 内容                       |
| ------------------- | -------------------------- |
| `docker`            | Dockerfile と HTML         |
| `compose`           | Compose の完成形           |
| `kubernetes/basic`  | 第2回の完成形              |
| `kubernetes/config` | 第3回の完成形              |
| `reference`         | 保存・公開などの解説用の例 |
| `extra`             | 講習後の Kustomize の例    |

## イメージを取得する

```sh
docker pull nginx:1.28.0-alpine
docker pull curlimages/curl:8.14.1
docker pull kindest/node:v1.34.0@sha256:7416a61b42b1662ca6ca89f02028ac133a309a2a30ba309614e8ec94d976dc5a
docker run --rm nginx:1.28.0-alpine nginx -v
```

最後のコマンドで nginx のバージョンが表示されれば、コンテナの実行も確認できています。
kind の Node イメージは大きいため、時間に余裕を持って取得します。

## kind の動作確認

事前確認用のクラスタを作り、削除します。既存のクラスタを操作しないよう、名前を指定します。

```sh
kind create cluster --name naro-preflight --image kindest/node:v1.34.0@sha256:7416a61b42b1662ca6ca89f02028ac133a309a2a30ba309614e8ec94d976dc5a --wait 120s
kubectl --context kind-naro-preflight get nodes
kind delete cluster --name naro-preflight
```

Node の状態が`Ready`なら準備完了です。
クラスタ作成に失敗した場合は、Docker のメモリ割り当てやディスク容量を確認してください。
削除後は`kubectl config get-contexts`で context の一覧を確認できます。

::: warning 操作先を指定する
実習の kubectl コマンドには`--context kind-naro-containers`を付けます。
自分が普段使っているクラスタへ誤って適用しないため、そのまま入力してください。
:::

## 教材のバージョン

2026 年 10 月 14 日に、公式の配布先と仕様を確認しました。
再現しやすいようにバージョンを固定しています。最新版を使うことが目的ではありません。

| 対象              | 教材の指定                                     |
| ----------------- | ---------------------------------------------- |
| nginx             | `nginx:1.28.0-alpine`                          |
| HTTP クライアント | `curlimages/curl:8.14.1`                       |
| kind              | v0.30.0                                        |
| Kubernetes        | v1.34.0。Node イメージの digest も固定         |
| kubectl           | v1.34 系を推奨                                 |
| Docker・Compose   | 上記イメージと Compose plugin を実行できる環境 |

Node イメージと対応する CPU は[kind v0.30.0 のリリース情報](https://github.com/kubernetes-sigs/kind/releases/tag/v0.30.0)で確認できます。
nginx と curl のタグは、それぞれ[Nginx の配布先](https://hub.docker.com/layers/library/nginx/1.28.0-alpine/images)と
[curl の配布先](https://hub.docker.com/layers/curlimages/curl/8.14.1/images)で確認できます。
kubectl の互換性は[Version Skew Policy](https://kubernetes.io/releases/version-skew-policy/#kubectl)を参照してください。

## 困ったとき

| 症状                     | 確認すること                                                           |
| ------------------------ | ---------------------------------------------------------------------- |
| Docker に接続できない    | Docker の起動状態、WSL integration、Linux の実行権限                   |
| イメージを取得できない   | インターネット接続、VPN、プロキシ、レジストリの制限                    |
| kind の作成が失敗する    | Docker のリソース割り当て、kind のバージョン                           |
| `exec format error`      | CPU とイメージ・実行ファイルのアーキテクチャが一致しているか           |
| `8088`番ポートを使えない | そのポートを使用しているアプリを停止するか、本文のホスト側ポートを変更 |

準備が終わったら[第1回](./2_docker_compose)へ進んでください。
