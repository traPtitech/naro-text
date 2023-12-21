# 初めに

第四部では、仮想的に接続された機器の設定を通して、インターネットが動くのに必要な要素技術を実用的な形で学習します。

1969 年にその原型となった ARPANET が誕生して以来、インターネットは今や生活のあらゆる場面に登場し、社会インフラの一部となりました。
この講習会では、インターネットを支える技術のなかでも識別子の扱いや経路制御、名前解決などに注目して掘り下げていきます。(前提知識は特に要求しません)

各講義は作業の背景知識となる**Lesson**と、実際に手を動かしてネットワークの設定をする**Assignment**で構成されています。

長年かけて発展してきたインターネットの技術を網羅的に扱う都合上、Lesson にはあまり詳しい内容は書かないので各自で積極的に調べるようにしてください。

この資料は過去に[インターネット構築講習会](https://github.com/hijiki51/InternetArchLecture)として開催されたものを移植したものです。

## 環境構築
- 動作環境
  - Ubuntu 20.04 LTS
  - GCP e2-micro
    - vCPU x2
    - memory 1GB

Exec this script:  
`curl https://raw.githubusercontent.com/hijiki51/InternetArchLecture/main/setup/setup.sh | bash`


### 再起動時の設定
現状 docker と bridge の接続がコンテナの停止によって切断されてしまう。
そのため、コンテナを再起動したときには以下の手順で復旧作業を行う。

#### サーバー
```sh
$ ovs-vsctl list-ports br-r4-server | xargs -IXXX ovs-vsctl del-port br-r4-server XXX
$ ovs-vsctl list-ports br-rEX-server | xargs -IXXX ovs-vsctl del-port br-rEX-server XXX
$ nic_full_reset
$ seq 1 3 | xargs -IXXX ovs-docker add-port br-r4-server ens4 sXXX
$ ovs-docker add-port br-rEX-server ens4 sEX
```

#### 各ルーター
```sh
$ config
$ load
$ commit
$ save
$ exit
```
