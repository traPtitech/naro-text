---
title: 第2回 Kubernetesの基本（1）
---

# 第2回 Kubernetes の基本（1）

## 目標と時間配分

Kubernetes に nginx を配置し、通信できるようにします。
Pod を消したときの再作成やレプリカ数の変更を通して、望ましい状態を保つ仕組みを学びます。

解説は、Compose との比較 10 分、クラスタ 10 分、YAML・Namespace・Pod 15 分、Deployment 15 分、通信 10 分です。
実習は、準備 10 分、配置 15 分、通信 15 分、再作成とレプリカ数の変更 15 分、確認 5 分を目安にします。

## 解説

### ComposeからKubernetesへ

Compose で、アプリの構成をファイルへまとめられました。
サービスが増えると、複数のマシンへ分散して配置したり、プロセスが停止しても必要な数を動かし続けたりしたくなります。
Kubernetes は、こうした配置や状態を管理する仕組みです。

Compose で学んだ、イメージ・通信・マウントの知識は引き続き使います。
ただし、Compose の設定をそのまま Kubernetes へ渡せるわけではありません。
Kubernetes では、実行・通信・設定などを、役割ごとのリソースとして記述します。

| Compose で学んだこと   | Kubernetes で対応する考え方                 |
| ---------------------- | ------------------------------------------- |
| サービスの実行設定     | Pod の実行設定と、それを管理する Deployment |
| サービス名による通信   | Service とクラスタ内 DNS                    |
| 環境変数・設定ファイル | Pod の環境変数、ConfigMap・Secret           |
| マウント               | Volume と、その種類ごとの保存先             |

Compose の`service`と Kubernetes の`Service`は、名前が似ていますが役割が異なります。
Kubernetes の Service は通信先をまとめるリソースです。

### クラスタと内部コンポーネント

Kubernetes が管理するマシンの集合をクラスタ、各マシンを Node と呼びます。
クラスタの状態を管理する部分が制御プレーンです。

```text
kubectl → API server ← controller / scheduler
               ↓
              etcd：クラスタの状態を保存
               ↑
Node の kubelet → コンテナランタイム → Pod 内のコンテナ
```

| コンポーネント     | 役割                                                  |
| ------------------ | ----------------------------------------------------- |
| API server         | リソースを読み書きする API の入口                     |
| etcd               | クラスタの設定と状態を保存する                        |
| scheduler          | まだ配置先のない Pod を、条件に合う Node へ割り当てる |
| controller         | 望ましい状態と現在の状態を比較し、必要な変更を行う    |
| kubelet            | Node 上で Pod のコンテナを動かし、状態を報告する      |
| コンテナランタイム | イメージからコンテナを作り、実行する                  |

kubectl は API server へリクエストを送るクライアントです。
ネットワークには、Pod の通信を実装する CNI plugin や、Service の転送を実装する kube-proxy なども関わります。
今回使う kind は Docker コンテナを Node として作り、その中で Kubernetes を動かします。
本番クラスタでも各リソースの役割は共通ですが、Node の数やネットワークの実装は異なります。

[Kubernetes の内部コンポーネント](https://kubernetes.io/docs/concepts/overview/components/)も参照してください。

### マニフェストを読む

Kubernetes のリソースを YAML などで記述したファイルをマニフェストと呼びます。
`kubectl apply -f`で、その内容を API server に送ります。
送った直後にすべてが起動するとは限らないので、状態も確認します。

| フィールド   | 意味                                 |
| ------------ | ------------------------------------ |
| `apiVersion` | 使用する API のグループとバージョン  |
| `kind`       | リソースの種類。Pod、Deployment など |
| `metadata`   | 名前、Namespace、label など          |
| `spec`       | 望ましい状態や実行設定               |

YAML ではインデントが構造を表します。タブを使わず、教材ではスペース 2 個でそろえます。
`-`はリストの項目です。
リソースによって構造は異なり、ConfigMap のように`spec`ではなく`data`を持つものもあります。
クラスタから取得した`status`は、Kubernetes が報告する現在の状態です。

### NamespaceとPod

Namespace は、クラスタ内のリソースを名前の範囲で分ける仕組みです。
実習では`container-workshop`を使います。
Namespace を分けるだけで通信まで隔離されるわけではありません。

Pod は、Kubernetes が配置する最小単位です。
1 つ以上のコンテナをまとめ、同じ Pod のコンテナはネットワークを共有します。
そのため、同じ Pod 内では`localhost`で通信できます。
別の Pod の`localhost`は別の場所です。
教材の Web サーバーは、1 つの Pod に 1 つの nginx コンテナを置きます。

Pod の IP アドレスは、再作成で変わることがあります。
Pod を手作業で作るだけでは、削除時に代わりを作る仕組みはありません。
継続して動かすアプリには、次の Deployment を使います。

### Deployment・ReplicaSet・label

Deployment は、Pod のテンプレートと必要な数を管理します。
ReplicaSet は、そのテンプレートの Pod を指定した数だけ存在させます。
通常は Deployment を作成し、ReplicaSet は Deployment に管理させます。

```text
Deployment：テンプレートと更新を管理
  └ ReplicaSet：必要な Pod 数を維持
      ├ Pod：Nginx
      └ Pod：Nginx
```

label は、リソースに付けるキーと値です。
selector は、label を条件にリソースを選ぶ指定です。
Deployment の selector は、自分が管理する Pod を選びます。
`spec.selector.matchLabels`と`spec.template.metadata.labels`を対応させます。
Deployment の selector は作成後に変更できないため、用途が分かる値を決めます。

### ServiceとDNS

Service は、Pod の IP が変わっても使える通信先を用意します。
selector に合う Pod を選び、指定したポートへ通信を転送します。
Service の`port`が入口、`targetPort`が Pod 側の接続先です。

```text
client Pod → web:80（Service）→ app=web の Pod:80 → Nginx
```

同じ Namespace からは Service 名`web`で接続できます。
ほかの Namespace からは`web.container-workshop`のように Namespace も指定します。
完全な名前は`web.container-workshop.svc.cluster.local`です。末尾はクラスタの設定によって変わります。

デフォルトの Service は、クラスタ内部向けの`ClusterIP`です。
PC のブラウザから確認するため、今回は`kubectl port-forward`を使います。
これは選ばれた Pod への一時的な転送で、一般公開や Service の負荷分散を確認するための機能ではありません。

## 実習

### 1. 実習用クラスタを作る

`docs/docker_kubernetes/samples`で実行します。
第1回を飛ばした場合も、次のビルドから始められます。

```sh
docker build -t naro-nginx:1 ./docker
kind create cluster --name naro-containers --image kindest/node:v1.34.0@sha256:7416a61b42b1662ca6ca89f02028ac133a309a2a30ba309614e8ec94d976dc5a --wait 120s
kind load docker-image naro-nginx:1 curlimages/curl:8.14.1 --name naro-containers
kubectl --context kind-naro-containers get nodes
kubectl --context kind-naro-containers apply -f kubernetes/namespace.yaml
```

Docker が持つイメージは、kind 内のコンテナランタイムには自動では渡りません。
`kind load docker-image`で実習用イメージを Node に読み込ませます。
`Ready`の Node と、Namespace の作成が確認できれば次へ進みます。
クラスタがすでに存在する場合は、`kind get clusters`で名前を確認し、作成コマンドを省略します。

### 2. nginxを配置する

完成形の Deployment を開きます。

<<< ./samples/kubernetes/basic/deployment.yaml{yaml}

`replicas: 1`は必要な Pod 数です。
`template`以下が Pod のテンプレートで、使うイメージやコンテナのポートを指定します。
`containerPort`はポートの情報であり、それだけで PC に公開されるわけではありません。
`imagePullPolicy: IfNotPresent`は、Node にイメージがあればそれを使う指定です。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/basic/deployment.yaml
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
kubectl --context kind-naro-containers -n container-workshop get deployment,replicaset,pod
kubectl --context kind-naro-containers -n container-workshop logs deployment/web
```

`rollout status`は配置の完了を待ちます。更新の詳しい仕組みは補足資料で扱います。
Pod の`READY`が`1/1`、`STATUS`が`Running`になることを確認してください。
Pod の名前には自動生成された末尾が付きます。

### 3. Serviceで接続する

<<< ./samples/kubernetes/basic/service.yaml{yaml}

`app: web`の Pod を選び、Service の`80`番から Pod の`80`番へ転送します。
HTTP クライアント用の Pod も作成します。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/basic/service.yaml -f kubernetes/basic/client.yaml
kubectl --context kind-naro-containers -n container-workshop wait --for=condition=Ready pod/client --timeout=120s
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS http://web
```

`Hello from Docker`が返れば、Service 名で通信できています。
次に、PC からの一時的なアクセスを用意します。

```sh
kubectl --context kind-naro-containers -n container-workshop port-forward service/web 8088:80 --address 127.0.0.1
```

このターミナルは転送中のままにし、ブラウザで`http://localhost:8088`を開きます。
確認できたら`Ctrl+C`で転送を終了してください。
Pod が入れ替わると転送も終了する場合があります。その場合は再実行します。

### 4. Podの再作成を見る

```sh
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web
kubectl --context kind-naro-containers -n container-workshop delete pod -l app=web
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web -w
```

`-l app=web`は、その label の Pod を選ぶ指定です。
`-w`は変更を続けて表示します。新しい名前の Pod が`Running`になったら`Ctrl+C`で終了します。
Pod を削除しても、必要な数を指定した Deployment と ReplicaSet が残っているため、代わりの Pod が作られます。

### 5. レプリカ数を変更する

`kubernetes/basic/deployment.yaml`の`replicas`を`2`に変更して適用します。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/basic/deployment.yaml
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web --show-labels
kubectl --context kind-naro-containers -n container-workshop get endpointslice -l kubernetes.io/service-name=web
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS http://web
```

Web 用の Pod が 2 つになり、EndpointSlice に接続先が登録されます。
EndpointSlice は、Service の実際の接続先を記録するリソースです。
再び Service を経由して HTML が返ることを確認します。

確認：Pod を作り直す操作と、Deployment の必要数を変更する操作は、何が違いますか。

## 確認問題

1. Deployment と Service は、それぞれ何を管理しますか。
2. Pod を削除したのに、新しい Pod ができたのはなぜですか。
3. 別の Pod の nginx に、`localhost`で接続できますか。
4. Namespace を作るだけで、ほかの Namespace との通信は遮断されますか。

::: details 解答

1. Deployment は Pod のテンプレートと必要数、Service は通信先を管理します。
2. ReplicaSet が必要な Pod 数を保つためです。
3. できません。Pod のネットワークは別です。今回は Service 名を使います。
4. 遮断されません。通信の制御には別の設定が必要です。
   :::

## よくある失敗

| 症状                         | 確認と対応                                                           |
| ---------------------------- | -------------------------------------------------------------------- |
| `ImagePullBackOff`           | イメージ名・タグ・pull policy を確認し、kind へイメージをロードする  |
| Pod が`Pending`              | `describe pod`の Events で、メモリ不足など配置できない理由を確認する |
| Service へ接続できない       | Pod の label、Service の selector とポート、EndpointSlice を確認する |
| `port-forward`を開始できない | Pod の状態と、PC の`8088`番を使っているプロセスを確認する            |

Pod の詳細は、一覧で確認した名前を指定して調べます。`POD_NAME`は実際の名前に置き換えてください。

```sh
kubectl --context kind-naro-containers -n container-workshop describe pod POD_NAME
kubectl --context kind-naro-containers -n container-workshop get events --sort-by=.metadata.creationTimestamp
```

## 終了・再開

第 3 回でもクラスタを使います。`port-forward`と`get -w`は`Ctrl+C`で終了してください。
PC の資源を空けたい場合は、次のコマンドで実習用クラスタだけを削除できます。

```sh
kind delete cluster --name naro-containers
```

次回は「実習用クラスタを作る」から再開し、続けて`kubernetes/basic`のファイルを適用します。
必要なマニフェストとイメージがあれば、同じ構成を作り直せます。

次は[Kubernetes の基本（2）](./4_configuration)です。
