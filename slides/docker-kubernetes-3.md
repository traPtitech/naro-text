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

# Kubernetes の基本（2）

Docker・Kubernetes 講習会 第3回

解説 60 分 / 実習 60 分

---

# 今日の到達点

- 実行・設定・保存・公開の役割を区別する。
- ConfigMap で HTML を変更する。
- Service の selector と通信先の関係を確かめる。

本文：[第3回の教材](https://traptitech.github.io/naro-text/docker_kubernetes/4_configuration.html)

---

# イメージと設定を分ける

<div class="flow"><div class="node">ConfigMap<br>HTML</div><span>→</span><div class="node">Volume<br>ディレクトリへ</div><span>→</span><div class="node">Nginx<br>配信</div></div>

イメージはプログラム、ConfigMap は値やファイルの内容です。
今回は HTML を渡して、変化を観察します。

---

# ConfigMap の内容

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: web-content
  namespace: container-workshop
data:
  index.html: |
    <h1>Hello from ConfigMap</h1>
```

キーがファイル名、値がその内容になります。

---

# 渡し方で反映方法が変わる

| 渡し方                   | 変更後                         |
| ------------------------ | ------------------------------ |
| 環境変数                 | Pod の作り直しが必要           |
| ディレクトリへのマウント | 時間を置いてファイルに反映     |
| subPath へのマウント     | 既存のマウントには反映されない |

ファイルの更新と、アプリの再読み込みも別です。
今回はリクエスト時に HTML を読みます。

---

# Secret は機密情報を扱う

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-credentials
stringData:
  password: example-only
```

base64 は符号化であり、暗号化ではありません。
実際の値を Git へ平文で保存しません。

---

# 保存先の寿命を区別する

| 保存先                     | 寿命                         |
| -------------------------- | ---------------------------- |
| コンテナの書き込み可能な層 | コンテナの作り直しで失われる |
| emptyDir                   | Pod の削除で失われる         |
| PVC で利用するストレージ   | Pod とは別に管理する         |

ConfigMap と Secret は、設定の読み取りに使います。

---

# PV・PVC・StorageClass

<div class="flow"><div class="node">Pod</div><span>→</span><div class="node">PVC<br>容量などの要求</div><span>→</span><div class="node">PV<br>ストレージ</div></div>

StorageClass はストレージの種類や作り方を指定します。
provisioner が PV を動的に用意する構成もあります。

---

# PVC の例

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: web-data
spec:
  accessModes: [ReadWriteOnce]
  resources:
    requests:
      storage: 1Gi
```

利用可能なストレージがなければ `Pending` になります。

---

# Service の種類

| 種類         | 入口                         |
| ------------ | ---------------------------- |
| ClusterIP    | クラスタ内部                 |
| NodePort     | Node のポート                |
| LoadBalancer | 外部ロードバランサーとの連携 |

外部へ届くかは、クラスタとネットワークの構成によります。
kind で種類を変えるだけでは公開 IP は付きません。

---

# Ingress は HTTP の振り分けルール

<div class="flow"><div class="node">外部の入口</div><span>→</span><div class="node">Ingress controller</div><span>→</span><div class="node">Service</div><span>→</span><div class="node">Pod</div></div>

ホスト名やパスを条件に Service へ転送します。
controller、DNS、外部からの接続経路が必要です。
Ingress API は機能追加が凍結され、Gateway API も使われます。

---

# Pod の管理方法を選ぶ

| リソース    | 用途                              |
| ----------- | --------------------------------- |
| Deployment  | 同じ役割の Pod を入れ替えて動かす |
| StatefulSet | Pod ごとの識別子や保存先を保つ    |
| DaemonSet   | 条件に合う各 Node へ配置する      |

例：Web サーバー / DB のクラスタ / ログ収集

StatefulSet だけで DB の複製やバックアップは設定されません。

---

<!-- _class: section-head -->

# 実習 60 分

---

# 実習の順序

| 内容                             | 目安  |
| -------------------------------- | ----- |
| 前回の構成から再開               | 5 分  |
| ConfigMap の作成・マウント・変更 | 20 分 |
| Service の selector の確認       | 15 分 |
| 最終課題                         | 15 分 |
| 後片付け                         | 5 分  |

---

# ConfigMap の演習で観察すること

1. イメージの HTML が ConfigMap の HTML に置き換わる。
2. ConfigMap だけ変えても、Pod は作り直されない。
3. 時間を置くと、マウントされたファイルが変わる。
4. Nginx が変更後の HTML を返す。

Nginx の設定ファイルを変える場合は、再読み込みも必要です。

---

# selector の演習で観察すること

Service の selector を `app: missing` に変えます。

- Pod が動いていても、通信先に選ばれません。
- EndpointSlice の接続先が空になります。
- client からの HTTP リクエストが失敗します。

`app: web` に戻して、通信の復旧を確認します。

---

# 最終課題と後片付け

ConfigMap の見出しと Deployment のレプリカ数を変更します。
変更したリソースと、通信経路を説明してください。

終了後は実習用 Namespace と kind クラスタを削除します。
具体的なコマンドは本文の「後片付け」を参照してください。

---

# 参考資料と次の学習

- [ConfigMap](https://kubernetes.io/docs/concepts/configuration/configmap/)
- [永続ストレージ](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
- [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)

更新、probe、リソース制限、Kustomize、GitOps は
[講習後の補足資料](https://traptitech.github.io/naro-text/docker_kubernetes/5_extra.html)へ。
