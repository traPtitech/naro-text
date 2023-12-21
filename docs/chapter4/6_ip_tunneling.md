# Chapter6: IP Tunneling

この章では、IP Tunneling の基礎知識とその設定を学びます。

[[toc]]
## Lesson

### Tunneling

トンネリングとは、間に複数のネットワーク機器が挟まれた状態のネットワーク機器同士をあたかも直結しているかのように見せる手法のことです。
GRE や IPIP、SIT などいくつか方法がありますがここでは IP in IP Tunneling を使用します。

### IP in IP Tunneling

IPIP では、IP パケットの上に更に IP ヘッダを付与しカプセル化することでトンネリングを実現します。

IPIP には Point to Point インタフェースである tunnel インタフェースが必要で、このインタフェースでカプセル化を解除します。
`Outer IP Header`には tunnel インタフェースのエンドポイントを識別する情報が記載されています。

```
                                         +---------------------------+
                                         |                           |
                                         |      Outer IP Header      |
                                         |                           |
     +---------------------------+       +---------------------------+  　   +---------------------------+       
     |                           |       |                           |   　  |                           |       
     |         IP Header         |       |         IP Header         |　　　　|         IP Header         |
     |                           |       |                           |    　 |                           |       
     +---------------------------+ ====> +---------------------------+ ====> +---------------------------+
     |                           |       |                           |   　  |                           |
     |                           |       |                           |   　  |                           |
     |         IP Payload        |       |         IP Payload        |   　  |         IP Payload        |
     |                           |       |                           |   　  |                           |
     |                           |       |                           |   　  |                           |
     +---------------------------+       +---------------------------+  　   +---------------------------+
```

::: warning
tunnel インタフェース上で OSPF を有効化し、source や destination の IP を tunnel インタフェース経由で学習してしまわないように注意してください。\
OSPF はネットワークで指定するため、IP アドレスが同じ`ens4`と tunnel インタフェースの両方で OSPF が有効になり、外側にある tunnel インタフェースで経路を学習するようになります。\
この場合、tunnel 接続のための NextHop が tunnel インタフェース自身となり`recursive routing`状態に陥ることで、tunnel インタフェースが UpDown を繰り返し、その結果 OSPF のネイバーも UpDown を繰り返します。
:::
## Assignment

### 1. 他の人のサーバーとIPIPトンネリングを張って、トンネルに指定したIPに対してpingが通ることを確認してください。

::: warning
トンネリングに使う IP アドレスは 172.16.以下を利用してください。\
tunnel インタフェースはホスト側に作成し、トンネリング設定を VyOS 内で行なってください。\
TTL を必ず指定してください(255 程度)。\
Chapter7 以降へ進むために必須となります。
:::

::: details ヒント1
ホストでの操作は`ip`コマンドを使用すると良いでしょう。
:::

::: details ヒント2

自分のグローバル IP アドレスと相手のグローバル IP アドレスが必要です。
:::

::: details ヒント3

`encapsulation`の形式は`ipip`です。
:::

::: details 答え
**[Host]**
```sh
root@150-95-184-195:~# ip tunnel add tun0 mode ipip remote <相手のグローバルIP> local <自分のグローバルIP> ttl 255
root@150-95-184-195:~# ip link set up tun0
```

**[rEX]**
```sh
root@150-95-184-195:~# attach rEX
vyos@rEX:/$ config
[edit]
vyos@rEX# set interfaces tunnel tun0 address 172.16.0.1/30

vyos@rEX# set interfaces tunnel tun0 local-ip <自分のグローバルIP>

vyos@rEX# set interfaces tunnel tun0 remote-ip <相手のグローバルIP>

vyos@rEX# set interfaces tunnel tun0 encapsulation ipip

vyos@rEX# commit
vyos@rEX# save
[edit]
vyos@rEX# exit
exit
```

:::