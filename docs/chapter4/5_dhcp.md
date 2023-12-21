# Chapter5: DHCP

この章ではDHCPに関する基礎知識とその設定を学びます。

[[toc]]
## Lesson

### DHCP

DHCP(Dynamic Host Configuration Protocol) は、ネットワーク接続するのに必要なIPアドレスなどの情報を自動的に割り当てるアプリケーション層プロトコルです。

DHCPは情報の割当を行うDHCPサーバーと情報を割り当てられるDHCPクライアントからなります。
最初は、DHCPクライアントは自身のIPアドレスも知らなければ、DHCPサーバのIPアドレスも知らないので全ての宛先(ブロードキャスト)にDHCP Discoverメッセージを送信して、ネットワーク全体に問い合わせを行います。

DHCP Discoverメッセージを受け取ったDHCPサーバは、クライアントに割り当てるIPアドレス設定などをアドレスプールから選択して提案します。
DHCPサーバの仕様によりDHCP Offerはブロードキャストで送信するので、その場合は宛先MACとIPはブロードキャストアドレスになります。

DHCP Offerを受け取ったDHCPクライアントは提案されたIPを使用する事を通知するためにDHCP Requestをブロードキャストします。

最後に、DHCPサーバはDHCPクライアントが使用するIPアドレスなどの設定情報をDHCP Ackで送信します。

これを受け取ることでDHCPクライアントは提案されたIPアドレスを自信のIPアドレスとして利用できるようになります。

DHCPサーバには色々な設定項目がありますが、DHCPクライアントにIPアドレスを割り当てる範囲のアドレスプール、サブネットマスク、デフォルトゲートウェイのアドレス、DNSサーバのIPアドレス、リース期間(IPアドレスの貸し出し期間)などが、基本的な設定項目となります。

![](assets/dhcp.png)

## Assignment

### 1. DHCPを使ってs1~s3がインターネットに接続できるようにしてみよう

::: info
s1~s3は`dhclient`導入済みUbuntu 20.04 LTSです。
:::


::: details ヒント1

ここではr4がDHCPサーバー、s1~s3がDHCPクライアントです。
:::


::: details ヒント2

最初はこれまでと同様にr4の`eth100`に対してIPアドレスやOSPFの設定を行う必要があります。
:::


::: details ヒント3

DHCPサーバーには少なくとも割り当てる範囲のアドレスプール・サブネットマスク、デフォルトゲートウェイのアドレスの3つが必要です。
:::


::: details 答え

**[r4]**
```sh
root@hijiki51-60000:/# attach r4
minion@r4:/$ config
[edit]
minion@r4# set interfaces ethernet eth100 address 192.168.0.142/28

minion@r4# set protocols ospf area 0 network 192.168.0.128/28

minion@r4# set protocols ospf passive-interface eth100

minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 default-router 192.168.0.142 ; 送信先ネットワークに対してデフォルトルート(今回はDHCPホストサーバー)を設定
minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 range 0 start 192.168.0.129
minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 range 0 stop 192.168.0.139 ; DHCPで使用するネットワークとその中で割り振る範囲を設定

minion@r4# commit
minion@r4# save
[edit]
minion@r4# exit
exit

minion@r4/$ show dhcp server statistics

Pool                      Pool size   # Leased    # Avail
----                      ---------   --------    -------
dhcp_scope_01             10          0           10
```

`show dhcp server statistics`で割り振りが行われているか確認可能です。

**[s1~s3]**
```sh
root@s1:~# dhclient ens4
```
でDHCPの再リースが可能です。

割り振り前後で`ip address`コマンドの結果を比較してみると良いでしょう。
:::

