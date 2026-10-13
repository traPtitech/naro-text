---
marp: true
theme: SysAd
paginate: true
style: |
  section { font-size: 28px; }
  section:not(.title, .section-head, .end) { padding: 125px 64px 65px; }
  section:not(.title, .section-head, .emphasize, .end) h1 { font-size: 42px; }
  section.title h1 { font-size: 66px; }
  p { font-size: 28px; }
  pre { font-size: 21px; }
  table { font-size: 23px; }
  .flow { display: flex; gap: 16px; align-items: center; justify-content: center; margin: 35px 0; }
  .node { border: 3px solid #14a39e; border-radius: 12px; padding: 20px; background: #edf8f7; text-align: center; }
  .note { font-size: 20px; }
---

<!-- _class: title -->

# Kubernetes の基本（1）

Docker・Kubernetes 講習会 第2回

解説 60 分 / 実習 60 分

---

# 今日の到達点

- Pod と Deployment の関係を説明する。
- Service 名で Nginx に接続する。
- Pod の再作成とレプリカ数の変更を確かめる。

本文：[第2回の教材](https://traptitech.github.io/naro-text/docker_kubernetes/3_kubernetes_basics.html)

---

# Compose から Kubernetes へ

構成を記述する知識は、そのまま使えます。

Kubernetes では、複数の Node への配置と、望ましい状態の維持を行います。
実行・通信・設定を、役割ごとのリソースに分けます。

Compose の設定ファイルをそのまま使うものではありません。

---

# クラスタと Node

クラスタは Kubernetes が管理するマシンの集合です。
各マシンを Node と呼びます。

- 制御プレーン：状態を管理し、配置を決める。
- Node：kubelet とコンテナランタイムが Pod を動かす。
- kind：Docker コンテナを Node として使う。

---

# API を通して状態を管理する

<div class="flow"><div class="node">kubectl</div><span>→</span><div class="node">API server</div><span>→</span><div class="node">etcd<br>状態を保存</div></div>

controller と scheduler も API を通して状態を確認します。
kubelet は割り当てられた Pod を動かし、状態を報告します。

---

# 内部コンポーネントの役割

| コンポーネント     | 役割                   |
| ------------------ | ---------------------- |
| API server         | API の入口             |
| scheduler          | Pod の配置先を決める   |
| controller         | 望ましい状態へ近づける |
| kubelet            | Node 上で Pod を動かす |
| コンテナランタイム | コンテナを実行する     |

---

# YAML でリソースを記述する

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
  namespace: container-workshop
spec:
  replicas: 1
  # 以下は省略
```

`spec` は望ましい状態、`status` は報告された状態です。

---

# Namespace と Pod

Namespace は、リソースを名前の範囲で分けます。
分けるだけで通信を遮断するものではありません。

Pod は、配置する最小単位です。
同じ Pod 内のコンテナはネットワークを共有します。

今回は 1 Pod に 1 Nginx コンテナを置きます。

---

# Deployment と ReplicaSet

<div class="flow"><div class="node">Deployment<br>テンプレートと更新</div><span>→</span><div class="node">ReplicaSet<br>Pod 数を維持</div><span>→</span><div class="node">Pod<br>Nginx</div></div>

必要数が 2 で現在の Pod が 1 つなら、もう 1 つ作ります。
Pod を削除しても、必要数の指定は残ります。

---

# label と selector

label はキーと値、selector は選択条件です。

```yaml
selector:
  matchLabels:
    app: web
template:
  metadata:
    labels:
      app: web
```

Deployment の selector と Pod の label を対応させます。

---

# Service で通信先を用意する

<div class="flow"><div class="node">client Pod</div><span>→</span><div class="node">Service<br>web:80</div><span>→</span><div class="node">app=web<br>Pod:80</div></div>

Pod の IP が変わっても、Service 名を使えます。
`port` は Service の入口、`targetPort` は Pod 側です。

---

# PC から確認する方法

今回は一時的な `port-forward` を使います。

```sh
kubectl --context kind-naro-containers \
  -n container-workshop \
  port-forward service/web 8088:80
```

選ばれた Pod への転送です。
Service の負荷分散を確かめる操作とは区別します。

---

<!-- _class: section-head -->

# 実習 60 分

---

# 実習の順序

| 内容                        | 目安  |
| --------------------------- | ----- |
| kind 作成・イメージのロード | 10 分 |
| Deployment の配置と状態確認 | 15 分 |
| Service と通信の確認        | 15 分 |
| Pod 再作成・レプリカ数変更  | 15 分 |
| 確認問題                    | 5 分  |

すべて実習用 context と Namespace を指定します。

---

# イメージを kind へ渡す

Docker のイメージは、kind 内に自動では渡りません。

```sh
kind load docker-image \
  naro-nginx:1 curlimages/curl:8.14.1 \
  --name naro-containers
```

マニフェストでは `imagePullPolicy: IfNotPresent` を指定します。

---

# 観察すること

- Deployment、ReplicaSet、Pod が作られたか。
- Service 経由で HTML が返ったか。
- Pod 削除後に、新しい名前の Pod ができたか。
- `replicas: 2` で Web 用 Pod が 2 つになったか。

クラスタは次回も使います。削除した場合も本文から再開できます。

---

# 参考資料

- [Kubernetes の内部コンポーネント](https://kubernetes.io/docs/concepts/overview/components/)
- [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [Service](https://kubernetes.io/docs/concepts/services-networking/service/)

次回は、設定・保存・公開のリソースを学びます。
