---
title: 講習後に学ぶこと
---

# 講習後に学ぶこと

このページは講習時間外の自習用です。
nginx を配置し、Service と ConfigMap を使えた後に、更新や設定管理を学ぶための入口を載せています。

## 更新とrollback

Deployment の Pod テンプレートを変更すると、新しい ReplicaSet が作られます。
デフォルトの RollingUpdate では、新しい Pod の起動と古い Pod の削除を進めて入れ替えます。
`rollout status`はその進行を確認するコマンドです。

イメージ更新を試す場合は、HTML を変えて`naro-nginx:2`をビルドし、kind にロードします。
第2回の ConfigMap をマウントしない構成で、Deployment のイメージを変更してください。
ConfigMap で HTML を上書きした構成では、イメージ内の HTML を変えても表示に現れません。

```sh
kubectl --context kind-naro-containers -n container-workshop rollout history deployment/web
kubectl --context kind-naro-containers -n container-workshop rollout undo deployment/web
kubectl --context kind-naro-containers -n container-workshop rollout status deployment/web --timeout=120s
```

`undo`は Deployment の Pod テンプレートを以前の revision へ戻します。
ConfigMap の内容や DB のデータまで戻るものではありません。
マニフェストの変更も元へ戻さないと、次の apply で再び新しい設定になります。
[Deployment の公式説明](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)を参照してください。

## アプリの状態を伝えるprobe

プロセスが動いていても、アプリがリクエストに応答できるとは限りません。
probe は、その状態を Kubernetes へ伝えるための検査です。

| 種類            | 用途                 | 失敗したとき                           |
| --------------- | -------------------- | -------------------------------------- |
| readiness probe | 通信を受けられるか   | Service の通常の転送先から外れる       |
| liveness probe  | 再起動が必要な状態か | コンテナが再起動される                 |
| startup probe   | 起動が完了したか     | 規定の失敗回数でコンテナが再起動される |

startup probe がある場合、成功するまで liveness・readiness の検査は始まりません。
検査対象や間隔をアプリに合わせます。依存先の一時的な不調だけで再起動を繰り返さないよう、用途を区別してください。

```yaml
# containers の項目に追加する例
readinessProbe:
  httpGet:
    path: /
    port: 80
  periodSeconds: 5
```

[probe の公式説明](https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/)で、それぞれの条件を確認できます。

## requestsとlimits

`requests`は、配置の判断に使う CPU・メモリの要求量です。
`limits`は、使用量の上限です。CPU の上限では実行が抑制されます。
メモリの上限を超えた場合は、プロセスが終了することもあります。

```yaml
# containers の項目に追加する例
resources:
  requests:
    cpu: 100m
    memory: 64Mi
  limits:
    cpu: 500m
    memory: 128Mi
```

CPU の`1000m`は 1 CPU、メモリの`Mi`は MiB を表します。
ここに載せた数値は説明用です。実際にはアプリの使用量を観察して決めます。
[リソース管理の公式説明](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)を参照してください。

## Kustomizeで構成を組み立てる

環境ごとにマニフェストを丸ごとコピーすると、共通部分の変更がそろわなくなります。
Kustomize は、元のリソースへ差分を適用して構成を生成するツールです。
教材では第2回の構成を参照し、レプリカ数を 2 に変更します。

<<< ./samples/extra/kustomization.yaml{yaml}

`samples`で次のコマンドを実行してください。

```sh
kubectl kustomize extra
```

これは生成結果を表示するだけで、クラスタを変更しません。
Deployment の`replicas: 2`を確認してください。
クラスタへ適用する場合は、対象を確認してから`kubectl apply -k`を使います。

[Kustomize の公式手順](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/)では、patch や ConfigMap の生成も学べます。

## GitOpsとOSSの起動設定

GitOps では、Git に保存した望ましい状態と、クラスタの状態を同期する仕組みを使います。
Argo CD や Flux がそのためのツールです。
手元から直接 apply する実習とは、変更を反映する経路が異なります。

新しい OSS を起動するときは、公式の起動例から、次の項目を確認すると構成を読みやすくなります。

- イメージとタグ、起動するコマンド。
- 待ち受けるポートと、Service からの通信経路。
- 環境変数や設定ファイル、機密情報の渡し方。
- 一時ファイルと、残す必要があるデータの保存先。
- 外部へ公開する必要があるか、起動時にほかのアプリへ接続するか。

公式のサンプルにも、開発用のパスワードや公開設定を含むものがあります。
学習用の設定をそのまま本番へ持ち込まず、各設定の用途を確認してください。
[Argo CD の概要](https://argo-cd.readthedocs.io/en/stable/)や[Flux の概要](https://fluxcd.io/flux/)から、同期の仕組みを学べます。
