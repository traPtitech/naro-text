---
title: Docker・Kubernetes講習会
description: Nginxを動かしながら、コンテナとKubernetesの基本を学ぶ講習会
---

# Docker・Kubernetes 講習会

同じアプリを動かすために、PC ごとにライブラリや設定をそろえるのは手間がかかります。
コンテナを使うと、アプリと実行に必要なファイルをまとめて配布できます。
さらに Kubernetes を使うと、複数のマシンにコンテナを配置し、指定した数のアプリを動かし続けられます。

この講習会では、Web サーバーの nginx で HTML を配信します。
Docker、Compose、Kubernetes の順に動かし方を変え、実行・通信・設定の仕組みを学びます。

## 対象と到達点

なろう講習会の第二部までを履修した人を想定しています。
ターミナル、Git、HTTP の基本は前提としますが、コンテナの知識は最初から説明します。
HTTP は「クライアントがリクエストを送り、サーバーがレスポンスを返す」という関係を思い出せれば十分です。

講習後は、次の操作と説明ができることを目指します。

- イメージとコンテナの違いを説明し、Docker と Compose でアプリを動かす。
- Kubernetes の主要リソースの役割を説明し、アプリを配置する。
- Service で通信先を指定し、ConfigMap でアプリに渡す内容を変更する。
- 状態とログを確認し、書いた設定と実際の動作を対応づける。

## 全体の構成

各回は解説約 1 時間、実習約 1 時間です。ツールのインストールとイメージ取得は事前に済ませます。

| 回                             | 解説                          | 実習                                                 |
| ------------------------------ | ----------------------------- | ---------------------------------------------------- |
| [第1回](./2_docker_compose)    | Docker・Compose               | nginx の起動、イメージ作成、マウント、コンテナ間通信 |
| [第2回](./3_kubernetes_basics) | Kubernetes の実行と通信       | Deployment・Service、Pod の再作成、レプリカ数の変更  |
| [第3回](./4_configuration)     | Kubernetes の設定・保存・公開 | ConfigMap による HTML の変更、label・selector の確認 |

まず[事前準備](./1_prepare)を進めてください。
本文の「解説」が座学、「実習」が手を動かすパートです。設定例は本文と同じ内容のファイルも用意しています。
途中から参加する場合は、各回の再開手順を使えます。

## スライド

スライドには解説の要点と実習の案内を載せています。操作の詳細と確認問題は本文を参照してください。

| 回                     | PDF                                                                                           | HTML                                                                                            |
| ---------------------- | --------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| Docker・Compose        | <a href="../slides/docker-kubernetes-1.pdf" target="_blank" rel="noopener noreferrer">PDF</a> | <a href="../slides/docker-kubernetes-1.html" target="_blank" rel="noopener noreferrer">HTML</a> |
| Kubernetes の基本（1） | <a href="../slides/docker-kubernetes-2.pdf" target="_blank" rel="noopener noreferrer">PDF</a> | <a href="../slides/docker-kubernetes-2.html" target="_blank" rel="noopener noreferrer">HTML</a> |
| Kubernetes の基本（2） | <a href="../slides/docker-kubernetes-3.pdf" target="_blank" rel="noopener noreferrer">PDF</a> | <a href="../slides/docker-kubernetes-3.html" target="_blank" rel="noopener noreferrer">HTML</a> |

## 教材の範囲

実習は各自の PC に作る専用クラスタで完結します。
特定組織の本番環境やリポジトリへのアクセスは不要です。
nginx を題材にしますが、Service や ConfigMap の考え方は、ほかの OSS を起動するときにも使えます。

ServiceAccount、RBAC、Job、CronJob は扱いません。
更新の管理や Kustomize、GitOps は、講習後の[補足資料](./5_extra)で扱います。

## 講師向けの進め方

解説の時間配分は各回の冒頭に記載しています。
実習では、各節の「確認」に到達した人から次へ進めます。
時間が足りない場合は、完成形サンプルから再開し、確認問題を自習に回してください。
保存や公開など、実習で使わないリソースは、用途と構成の説明にとどめます。

事前に講師も演習を一通り実行し、使用するツールのバージョンを確認してください。
教材で指定したバージョン、イメージの参照先、確認日は[事前準備](./1_prepare#教材のバージョン)に記載しています。
