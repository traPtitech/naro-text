# Chapter6: IP Tunneling

この章では、IP Tunnelingの基礎知識とその設定を学びます。

[[toc]]
## Lesson

### Tunneling

トンネリングとは、間に複数のネットワーク機器が挟まれた状態のネットワーク機器同士をあたかも直結しているかのように見せる手法のことです。
GREやIPIP、SITなどいくつか方法がありますがここではIP in IP Tunnelingを使用します。

### IP in IP Tunneling

IPIPでは、IPパケットの上に更にIPヘッダを付与しカプセル化することでトンネリングを実現します。

IPIPにはPoint to Pointインタフェースであるtunnelインターフェースが必要で、このインターフェースでカプセル化を解除します。
`Outer IP Header`にはtunnelインターフェースのエンドポイントを識別する情報が記載されています。

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
tunnelインターフェース上でOSPFを有効化し、sourceやdestinationのIPをtunnelインターフェース経由で学習してしまわないように注意してください。\
OSPFはネットワークで指定するため、IPアドレスが同じ`ens4`とtunnelインタフェースの両方でOSPFが有効になり、外側にあるtunnelインタフェースで経路を学習するようになります。\
この場合、tunnel接続のためのNextHopがtunnelインタフェース自身になり`recursive routing`状態に陥ることで、tunnelインタフェースがUpDownを繰り返し、その結果OSPFのネイバーもUpDownを繰り返します。
:::
## Assignment

### 1. 他の人のサーバーとIPIPトンネリングを張って、トンネルに指定したIPに対してpingが通ることを確認してください。

::: warning
トンネリングに使うIPアドレスは172.16.以下を利用してください。\
tunnelインターフェースはホスト側に作成し、トンネリング設定をVyOS内で行なってください。\
TTLを必ず指定してください(255程度)。\
Chapter7以降に進むために必須になります。
:::

::: details ヒント1
ホストでの操作は`ip`コマンドを使用すると良いでしょう。
:::

::: details ヒント2

自分のグローバルIPアドレスと相手のグローバルIPアドレスが必要です。
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