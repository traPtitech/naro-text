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

# Docker・Compose

Docker・Kubernetes 講習会 第1回

解説 60 分 / 実習 60 分

---

# 今日の到達点

- イメージとコンテナを区別する。
- Nginx を動かし、PC から接続する。
- ファイルを渡す方法を選ぶ。
- Compose で構成と通信を記述する。

本文：[第1回の教材](https://traptitech.github.io/naro-text/docker_kubernetes/2_docker_compose.html)

---

# コンテナは隔離されたプロセス

アプリと必要なファイルをそろえ、分離した環境で実行します。

| 仮想マシン            | Linux コンテナ               |
| --------------------- | ---------------------------- |
| OS のカーネルも動かす | 実行環境のカーネルを共有する |

Docker Desktop では Linux 仮想マシンの中で動きます。

<!-- namespace と cgroup は概念の紹介にとどめます。 -->

---

# イメージからコンテナを作る

<div class="flow"><div class="node">イメージ<br>ファイルと起動設定</div><span>→</span><div class="node">コンテナ A<br>実行環境</div><span>→</span><div class="node">Nginx<br>リクエストを待つ</div></div>

同じイメージから、コンテナ B も作れます。

メインプロセスが終了すると、コンテナも停止します。

---

# ファイル変更の保存先

イメージは読み取り専用です。
コンテナで変更したファイルは、そのコンテナの層へ保存されます。

- イメージへ変更が戻るわけではありません。
- コンテナを削除すると、その層の変更は失われます。
- 残したいデータにはマウントを使います。

---

# Dockerfile でイメージを作る

```dockerfile
FROM nginx:1.28.0-alpine
COPY html/ /usr/share/nginx/html/
```

```sh
docker build -t naro-nginx:1 .
```

`FROM` は元のイメージ、`COPY` はビルド時のコピーです。
最後の `.` はビルドへ渡すファイルの範囲です。

---

# タグとレジストリ

`nginx:1.28.0-alpine` = イメージ名 + タグ

- レジストリ：イメージを保存・配布する場所。
- `pull`：イメージを取得する。
- `push`：イメージを送る。

タグは同じ名前でも内容が変わる場合があります。
内容まで固定する場合は digest を使います。

---

# ポートの入口を指定する

<div class="flow"><div class="node">ブラウザ</div><span>→</span><div class="node">PC<br>127.0.0.1:8088</div><span>→</span><div class="node">コンテナ<br>80 番</div></div>

```sh
docker run -d --name naro-web \
  -p 127.0.0.1:8088:80 nginx:1.28.0-alpine
```

`EXPOSE` を書くだけでは、PC へ公開されません。

---

# ファイルを渡す方法

| 方法            | 主な用途                       |
| --------------- | ------------------------------ |
| イメージへ COPY | プログラムや静的ファイルの配布 |
| bind mount      | ホストのソースや設定を渡す     |
| named volume    | DB などのデータを保持する      |

マウント先に元からあるファイルは、マウント中は隠れます。

---

# Compose で構成を記述する

```yaml
services:
  web:
    image: nginx:1.28.0-alpine
    ports:
      - '127.0.0.1:8088:80'
    volumes:
      - ./html:/usr/share/nginx/html:ro
```

サービス、イメージ、ポート、マウントをまとめます。

---

# コンテナ間の通信

<div class="flow"><div class="node">client</div><span>→</span><div class="node">web:80</div><span>→</span><div class="node">Nginx</div></div>

同じ Compose ネットワークではサービス名で接続します。

client の `localhost` は client 自身です。
PC 向けの `8088` 番とは区別します。

---

# 起動と準備完了は別

Compose の `depends_on` だけでは、相手がリクエストを受けられるまで待てません。

必要に応じて healthcheck やリトライを使います。
教材の client はリトライ付きで実行します。

---

<!-- _class: section-head -->

# 実習 60 分

---

# 実習の順序

| 内容                           | 目安  |
| ------------------------------ | ----- |
| Nginx 起動・ログ・停止・削除   | 15 分 |
| HTML 入りイメージのビルド      | 10 分 |
| bind mount で内容変更          | 10 分 |
| Compose とサービス名による通信 | 20 分 |
| 確認問題・後片付け             | 5 分  |

操作は本文の「実習」に沿って進めてください。

---

# 確認すること

- ホストの HTML を変えて、ビルド済みイメージは変わるか。
- bind mount では、何を読み取っているか。
- `localhost:8088` と `web:80` はどこを指すか。

Compose は `down` で終了します。
`naro-nginx:1` は次回も使います。

---

# 参考資料

- [Docker のコンテナ解説](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-container/)
- [Compose のネットワーク](https://docs.docker.com/compose/how-tos/networking/)
- [教材の事前準備](https://traptitech.github.io/naro-text/docker_kubernetes/1_prepare.html)

次回は、Nginx を Kubernetes に配置します。
