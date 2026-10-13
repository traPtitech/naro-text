---
title: 第3回 Kubernetesの基本（2）
---

# 第3回 Kubernetes の基本（2）

## 目標と時間配分

アプリの実行設定、保存先、外部からの入口を区別します。
実習では ConfigMap から HTML を渡し、イメージをビルドせずに表示内容を変更します。

解説は、ConfigMap・Secret 20 分、保存 20 分、公開 10 分、ワークロードの種類 10 分です。
実習は、再開 5 分、ConfigMap 20 分、selector の確認 15 分、最終課題 15 分、後片付け 5 分を目安にします。

## 解説

### ConfigMapで設定を渡す

アプリのプログラムが同じでも、接続先や表示内容は環境によって変わります。
ConfigMap は、こうした値やファイルの内容を、イメージとは別に管理するリソースです。
キーと値を保存し、Pod へ環境変数やファイルとして渡せます。
機密情報を入れる用途には使いません。

```text
イメージ → Nginx のプログラム
ConfigMap の index.html → Volume → Nginx が配信するディレクトリ
```

今回は設定の受け渡しを観察しやすいよう、HTML の内容を ConfigMap に入れます。
通常のアプリでは、接続先の設定ファイルなどを渡す場合にも使います。
大きなファイルの配布先としては使わず、ConfigMap のデータは 1 MiB を超えられません。

| 渡し方                            | 変更した後の動作                               |
| --------------------------------- | ---------------------------------------------- |
| 環境変数                          | 起動時の値のまま。Pod の作り直しが必要         |
| ディレクトリとしてマウント        | 時間を置いてファイルへ反映される               |
| `subPath`で単一ファイルをマウント | ConfigMap の変更は既存のマウントへ反映されない |

ファイルが更新されることと、アプリがそれを読み直すことは別です。
今回の nginx はリクエスト時に HTML を読むため、更新後の内容を確認できます。
Nginx の設定ファイルを変える場合は、設定の再読み込みも必要です。
[ConfigMap の公式説明](https://kubernetes.io/docs/concepts/configuration/configmap/)も参照してください。

### Secretで機密情報を区別する

Secret は、パスワードやトークンなどを扱うリソースです。
ConfigMap と同様に、環境変数やファイルとして Pod へ渡せます。
次の例の値は説明用で、実際のパスワードではありません。

<<< ./samples/reference/secret.yaml{yaml}

`stringData`には文字列を書けます。API では、base64 で符号化された`data`として保存されます。
base64 は復元できる表現方法であり、暗号化ではありません。
Secret という種類にするだけで、保存時の暗号化やアクセス制限の設定まで済むわけではありません。
実際の値を Git へ平文で保存せず、利用する環境の管理方法に従います。

ConfigMap と Secret は、Pod が実行される Namespace に用意します。
また、マウントした ConfigMap や Secret をアプリの書き込み先にはできません。

### Volume・PV・PVC・StorageClass

コンテナの書き込み可能な層に置いたデータは、コンテナの作り直しで失われます。
Kubernetes の Volume は、Pod に別の保存先をマウントする仕組みです。
どれだけ長くデータが残るかは、Volume の種類によって異なります。

| 種類                | 主な用途とデータの寿命                                               |
| ------------------- | -------------------------------------------------------------------- |
| `emptyDir`          | Pod 内の一時ファイル。コンテナの再起動では残り、Pod の削除で失われる |
| ConfigMap・Secret   | 設定の読み取り。書き込み用の保存先ではない                           |
| PVC から使う Volume | Pod とは別の寿命でデータを保持する                                   |

PersistentVolume（PV）は、クラスタが扱う永続ストレージのリソースです。
PersistentVolumeClaim（PVC）は、アプリ側が必要な容量や利用方法を要求するリソースです。
StorageClass は、ストレージの種類や動的な用意の仕方を指定します。

```text
Pod の volumeMounts
  ↓ volumes の persistentVolumeClaim
PVC：必要な容量と利用方法を要求
  ↓ bind
PV：使用するストレージ
  ↑ provisioner が StorageClass に従って作成する場合もある
```

<<< ./samples/reference/pvc.yaml{yaml}

この PVC は`1Gi`の容量を要求しています。
`ReadWriteOnce`は単一 Node から読み書きするアクセスモードで、単一 Pod だけを意味するわけではありません。
`storageClassName`を省略すると、デフォルトの StorageClass があればそれを使います。
利用可能な PV や provisioner がなければ、PVC は`Pending`のままになります。

Pod では、次のように PVC を参照します。これはフィールドの抜粋で、単独で適用するマニフェストではありません。

```yaml
volumeMounts:
  - name: data
    mountPath: /data
volumes:
  - name: data
    persistentVolumeClaim:
      claimName: web-data
```

PVC を削除した後にデータが残るかは、PV の回収ポリシーやストレージの構成によります。
永続ストレージを使うことと、バックアップを取ることも別です。
kind のローカルストレージは、クラスタの削除後も残る本番用ストレージとしては扱いません。

### ServiceとIngressで入口を作る

前回の Service はデフォルトの`ClusterIP`で、クラスタ内部から利用しました。
ほかに、Node のポートを使う`NodePort`と、外部のロードバランサー連携を要求する`LoadBalancer`があります。
実際に外部から届くかは、クラスタの実装やネットワーク設定によります。
kind で種類を変更するだけで、クラウドの公開 IP が割り当てられるわけではありません。

Ingress は、HTTP のホスト名やパスを条件に Service へ振り分けるルールです。
ルールを実際に処理する Ingress controller が別途必要です。

```text
ブラウザ → 外部の入口 → Ingress controller
                             ↓ ホスト名・パスを判定
                          Service → Pod
```

<<< ./samples/reference/ingress.yaml{yaml}

この例は`workshop.example.com`宛ての`/`以下を`web` Service へ送るルールです。
`ingressClassName: example`は説明用で、対応する controller は実習クラスタに導入していません。
このファイルを適用するだけでは公開できません。DNS と外部からの接続経路も必要です。

Ingress API は機能追加が凍結されており、新しい公開方法として Gateway API も使われます。
ここでは既存の構成でも見かける Ingress の役割を押さえます。
[Ingress の公式説明](https://kubernetes.io/docs/concepts/services-networking/ingress/)も参照してください。

### Deployment・StatefulSet・DaemonSet

アプリの性質によって、Pod の管理方法を選びます。

| リソース    | 使う場面                          | 例                    |
| ----------- | --------------------------------- | --------------------- |
| Deployment  | 同じ役割の Pod を入れ替えて動かす | Web サーバーや API    |
| StatefulSet | Pod ごとの識別子や保存先を保つ    | DB のクラスタなど     |
| DaemonSet   | 条件に合う各 Node へ Pod を置く   | ログ収集、Node の監視 |

StatefulSet の Pod は`db-0`のような安定した識別子を持ち、`volumeClaimTemplates`で Pod ごとの PVC を要求できます。
DaemonSet は必要な Node ごとに Pod を配置するため、Deployment のようにレプリカ数を指定して管理するものではありません。

```yaml
# StatefulSet の spec の一部
serviceName: db
replicas: 2
volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ReadWriteOnce]
      resources:
        requests:
          storage: 1Gi
```

これは構造を読むための抜粋です。StatefulSet の Pod テンプレートや通信を支える Service などは省略しています。
StatefulSet を選ぶだけで DB の複製やバックアップが自動で設定されるわけではありません。
DB 自体の構成も必要です。

## 実習

### 1. 前回の状態から再開する

`docs/docker_kubernetes/samples`で実行します。
クラスタを削除していた場合は、[第2回の準備](./3_kubernetes_basics#実習)を実行してください。
第2回の途中で終了していても、次のコマンドで完成形を適用できます。

```sh
kubectl --context kind-naro-containers apply -f kubernetes/namespace.yaml
kubectl --context kind-naro-containers -n container-workshop apply -k kubernetes/basic
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
kubectl --context kind-naro-containers -n container-workshop wait --for=condition=Ready pod/client --timeout=120s
```

### 2. ConfigMapを作る

<<< ./samples/kubernetes/config/configmap.yaml{yaml}

`data`のキー `index.html`が、マウント後のファイル名になります。
`|`は後ろの複数行を文字列として記述する YAML の書き方です。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/config/configmap.yaml
kubectl --context kind-naro-containers -n container-workshop get configmap web-content -o yaml
```

出力に HTML があることを確認してください。

### 3. nginxへマウントする

<<< ./samples/kubernetes/config/deployment.yaml{yaml}

`volumes`で`web-content`を参照し、`volumeMounts`でコンテナ内の場所を指定します。
同じ`name: content`で両者を対応させます。
マウントによって、イメージに含まれていた HTML は隠れ、ConfigMap の内容が見えるようになります。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/config/deployment.yaml
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS http://web
```

`Hello from ConfigMap`が返れば成功です。
ブラウザで確認したい場合は、別のターミナルで実行します。

```sh
kubectl --context kind-naro-containers -n container-workshop port-forward service/web 8088:80 --address 127.0.0.1
```

`http://localhost:8088`を開きます。確認後は`Ctrl+C`で転送を止めます。

### 4. HTMLを変更する

`kubernetes/config/configmap.yaml`の見出しを`Updated with ConfigMap`に変更します。
編集前の Pod 名も確認しておきます。

```sh
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/config/configmap.yaml
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS http://web
```

マウントされたファイルは即時に更新されるとは限りません。
少し時間を置いて最後のコマンドを繰り返し、表示が変わることを確認してください。
反映時間は kubelet の同期やキャッシュの設定に依存します。
Pod が複数ある場合、それぞれへの反映時点が異なることもあります。

```sh
kubectl --context kind-naro-containers -n container-workshop exec deployment/web -- cat /usr/share/nginx/html/index.html
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web
```

今回は ConfigMap の内容だけを変更しました。この変更で Pod が作り直されることはありません。
反映を待てない場合は、次のコマンドで Pod を作り直せます。
自然な反映と再作成の違いを確認してから使ってください。

```sh
kubectl --context kind-naro-containers -n container-workshop rollout restart deployment/web
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
```

### 5. Serviceのselectorを確かめる

Pod の label と Service の selector を表示します。

```sh
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web --show-labels
kubectl --context kind-naro-containers -n container-workshop get service web -o yaml
```

`kubernetes/config/service.yaml`の selector を`app: missing`に変更して適用します。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/config/service.yaml
kubectl --context kind-naro-containers -n container-workshop get endpointslice -l kubernetes.io/service-name=web -o yaml
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS --max-time 5 http://web
```

接続先の更新には少し時間がかかる場合もあります。
EndpointSlice の`endpoints`が空になると、HTTP リクエストは失敗します。
nginx が起動していても、selector が一致しなければ Service の接続先になりません。
この確認は`port-forward`ではなく、client Pod から Service へ通信して行います。

selector を`app: web`へ戻し、同じ Service の apply を実行します。
EndpointSlice に接続先が戻り、再度 HTML を取得できることを確認してください。
Deployment 側の selector は変更しません。

### 6. 最終課題

次の条件を満たす構成へ変更してください。

1. ConfigMap で配信する見出しを、自分で決めた文字列にする。
2. `kubernetes/config/deployment.yaml`のレプリカ数を`2`にする。
3. client から Service 名でページを取得する。
4. 変更したリソース、通信経路、表示の確認方法を説明する。

::: details 操作例
ConfigMap の HTML と Deployment の`replicas`を編集します。
Service の selector が`app: web`へ戻っていることも確認します。

```sh
kubectl --context kind-naro-containers -n container-workshop apply -f kubernetes/config
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
kubectl --context kind-naro-containers -n container-workshop get pods -l app=web
kubectl --context kind-naro-containers -n container-workshop exec client -- curl -fsS http://web
```

表示がまだ古ければ、ConfigMap の反映を待ちます。
Pod の数は Deployment、通信先は Service、HTML は ConfigMap が担当しています。
:::

### 7. 後片付け

`port-forward`を`Ctrl+C`で終了してから実行します。

```sh
kubectl --context kind-naro-containers delete namespace container-workshop
kind delete cluster --name naro-containers
docker image rm naro-nginx:1
```

Namespace の削除は、その中のリソースも削除します。
マニフェストやホスト側の HTML は残るので、再実行できます。
ほかのアプリも使う nginx や curl のイメージは、一括削除しません。

## 確認問題

1. ConfigMap を変更すると、Pod は必ず作り直されますか。
2. Secret の base64 表現を Git に保存すれば、秘密を守れますか。
3. PV と PVC の役割は何ですか。
4. Ingress のマニフェストだけで HTTP の公開が完了しますか。
5. Node ごとにログ収集用 Pod を配置するには、どのリソースが合いますか。

::: details 解答

1. 作り直されません。渡し方とアプリの読み方によって、反映の手順が変わります。
2. 守れません。base64 で符号化した値は、元の値へ戻せます。
3. PV はストレージを表します。PVC でアプリから容量や利用方法を要求します。
4. 完了しません。controller と外部からの接続経路、DNS なども必要です。
5. DaemonSet です。
   :::

## よくある失敗

| 症状                            | 確認と対応                                                                 |
| ------------------------------- | -------------------------------------------------------------------------- |
| マウントできず Pod が起動しない | `describe pod`の Events で ConfigMap 名、キー、Namespace を確認する        |
| HTML が古い                     | ConfigMap の適用結果とコンテナ内ファイルを確認し、反映を待つ               |
| Service の通信が戻らない        | selector が`app: web`へ戻っているか、EndpointSlice と Pod の状態を確認する |
| レプリカ数が意図せず戻る        | 適用した Deployment ファイルの`replicas`を確認する                         |

各リソースの役割をつかめたら、必要に応じて[講習後に学ぶこと](./5_extra)へ進んでください。
