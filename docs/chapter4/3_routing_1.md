# Chapter3: Routing 1

この章では、静的経路制御(Static Routing)の基礎知識とその設定を学びます。

[[toc]]
## Lesson

### 経路制御

ここまでは1hop(直接接続した機器同士)のパケット送信でしたが、普段インターネットを利用する上では通信先の間には数多くのルーターが存在します。
パケットは送信元から送信先へバケツリレーのように運ばれていくわけですが、この経路は一意ではありません。

様々な経路制御手法によってパケットの送信経路はコントロールされますが、ここでは各ルーターに対して宛先に対応する次のノードを指定することでパケットの経路制御を実現してみましょう。

## Assignment

### 1. 静的ルーティングによってr3からrEX, r5, `8.8.8.8`のすべてに`ping`が通るようにしてみよう

::: details ヒント1

Chapter1と同様に使用する各NICにはIPアドレスを割り当てる必要があります。
::: 

::: details ヒント2
`ping`の応答パケットのルーティングも必要です。
::: 

::: details ヒント3
「VyOS static route set」などで検索してみると良いでしょう。
::: 

::: details 答え
以降、解答ではIPアドレスを以下のように設定し、NICに適切に割り振っています。
各自の環境に合わせて読み替えるようにしてください。

![IP Setting](assets/ip-setting.drawio.svg)

**[rEX]**
```sh
root@hijiki51-60000:/# attach rEX
minion@rEX:/$ config
[edit]
minion@rEX# set protocols static route 192.168.0.0/28 next-hop 192.168.0.2

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
minion@r1# set protocols static route 192.168.0.8/30 next-hop 192.168.0.6
minion@r1# set protocols static route 192.168.0.16/30 next-hop 192.168.0.6
minion@r1# commit
minion@r1# save
[edit]
minion@r1# exit
exit
```

**[r2]**
```sh
root@hijiki51-60000:/# attach r2
minion@r2:/$ config
[edit]
minion@r2# set protocols static route 0.0.0.0/0 next-hop 192.168.0.5
minion@r2# commit
minion@r2# save
[edit]
minion@r2# exit
exit
```

**[r3]**
```sh
root@hijiki51-60000:/# attach r3
minion@r3:/$ config
[edit]
minion@r3# set protocols static route 0.0.0.0/0 next-hop 192.168.0.9

minion@r3# commit
minion@r3# save
[edit]
minion@r3# exit
exit
```

**[r5]**
```sh
root@hijiki51-60000:/# attach r5
minion@r5:/$ config
[edit]
minion@r5# set protocols static route 0.0.0.0/0 next-hop 192.168.0.17
minion@r5# commit
minion@r5# save
[edit]
minion@r5# exit
exit
```
:::