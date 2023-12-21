# Chapter2: NAT

この章では、NAT(Network Address Translation)の基礎知識とその設定を学びます。

[[toc]]
## Lesson

### NAT(Network Address Translation)

NAT はその名の通り IP アドレスを変換するための技術です。
一般的にはプライベート IP アドレスとグローバル IP アドレスを相互変換するために用いられます。

プライベートネットワーク内からグローバルネットワークに存在するホストを指定してパケットを送信するとき、つまりインターネットに接続して通信するとき、使用する IP アドレスはユニークなもの(=グローバル IP アドレス)でなければならないため、送信元 IP アドレスをグローバル IP アドレスに変換する必要があります。

また、プライベートネットワーク内のホストを宛先にパケットを送信したい場合、NAT サーバーが受け取った後適切に宛先 IP アドレスをプライベート IP アドレスに書き換える必要があります。

::: info
現実では、NAPT(Network Address Port Translation, IP masquerade)と呼ばれる技術を用いてポート番号を動的に変更することで、1 つの IP アドレスに複数のホストを割り当てられるようにしています。
::: 
### Gateway

ゲートウェイは、通信プロトコルの異なるネットワーク同士が通信する際に中継する役割を備えた機器やそれに関するソフトウェアのことです。

**デフォルトゲートウェイ**はゲートウェイの一種で、内部ネットワークと外部ネットワークを接続するためのノードです。
一般に、経路の判明していないパケットはデフォルトゲートウェイに送られます。


## Assignment

### 1. r1から`8.8.8.8`に`ping`が通るようにしてみよう

::: details ヒント1

`8.8.8.8`は Google の運用するパブリック DNS サーバーです。
プライベートネットワークの外と通信するためにはグローバル IP アドレスを使用する必要があります。
:::

::: details ヒント2
初期状態では、r1 は`8.8.8.8`にパケットを送るための経路を使うべきではない言葉なので修正してくださいのでパケットを破棄します。
r1 に rEX がデフォルトゲートウェイであることを設定する必要があります。
:::

::: details ヒント3

インターネットに接続している rEX で、r1 から来たパケットの送信元アドレスをグローバル IP アドレスに書き換える必要があります。
NAT 時に必要な情報は、「どの範囲から来たパケットに NAT を適用するか」です。
:::

::: details 答え

**[rEX]**
```sh
root@hijiki51-60000:/# attach rEX
minion@rEX:/$ config
[edit]
minion@rEX# set nat source rule 1 outbound-interface ens4 ;グローバルネットワークとの接続点
minion@rEX# set nat source rule 1 source address 192.168.XXX.0/24 ;NATを適用する送信元ネットワークの範囲

minion@rEX# set nat source rule 1 translation address masquerade

minion@rEX# commit
minion@rEX# save
[edit]
minion@rEX# exit
exit
```

**[r1]**
```sh
root@hijiki51-60000:/# attach r1
minion@r1:/$ config
[edit]
minion@r1# set protocols static route 0.0.0.0/0 next-hop 192.168.XXX.1 ;送信先ネットワークに応じて次のノードを指定

minion@r1# commit
minion@r1# save
minion@r1# exit
minion@r1:/$ exit
root@hijiki51-60000:/# ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_req=1 ttl=55 time=1.59 ms
64 bytes from 8.8.8.8: icmp_req=2 ttl=55 time=1.06 ms
^C
--- 8.8.8.8 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1001ms
rtt min/avg/max/mdev = 1.068/1.332/1.596/0.264 ms
```

:::