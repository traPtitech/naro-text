# Chapter8: DNS 1

この章では DNS(Domain Name System)に関する基礎知識とその設定を学びます。

[[toc]]

## Lesson

### DNS
DNS(Domain Name System)は IP アドレスなどのリソースに対して別名をつけるアプリケーション層プロトコルです。

DNS はドメイン名空間(Domain Name Space)を管理するネームサーバーと名前解決をするリゾルバからなります。

ドメイン名空間は 1 つのネームサーバーで一括管理されているのではなく、木構造のように複数ネームサーバーによって分散管理されており、名前解決をする際は木構造の親から子へと再帰的に問い合わせします。

また、一番初めに問い合わせをし、全体の木構造の親に当たるネームサーバーのことをルートネームサーバーと呼びます。

現在 1600 を超えるルートネームサーバーが起動しています。

https://root-servers.org/
## Assignment

::: info
ns は bind9 導入済み Ubuntu のインスタンスです。
:::
### 1. ネームサーバーを構築してみよう
ネームサーバーに`server.{あなたの traQ ID}`のレコードを登録してみましょう。

後々使うので s1~s3 のどれかの IP にしておくといいです。



:::details ヒント1
扱うゾーンは`{あなたの traQ ID}`になるでしょう。(TLD です)
:::

:::details ヒント2
今回は IPv4 を用いているので設定するのは`A`レコードです。
:::

:::details ヒント3
「bind9 A レコード設定」などで検索してみるといいでしょう
:::

:::details 答え
NIC には事前に IP アドレスを割り当てておきます。

まずはサーバーに割り振られた IP アドレスを確認します（DHCP であるため手動確認）
```sh
root@s1:/# ip address
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
61: ens4@if62: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 7a:53:1d:45:76:56 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 192.168.0.129/28 brd 192.168.0.143 scope global dynamic ens4
       valid_lft 86332sec preferred_lft 86332sec
```

ゾーンの設定をします。\
`/etc/bind/named.conf.local`
```
//
// Do any local configuration here
//

// Consider adding the 1918 zones here, if they are not used in your
// organization
//include "/etc/bind/zones.rfc1918";

zone "hijiki51" IN {
        type master;
        file "/etc/bind/rr/hijiki51";
};
```


各種リソースレコードの設定をします。\
`/etc/bind/rr/hijiki51`

```
$TTL 60
@       IN      SOA ns.{自分のtraQ ID}. root.{自分のtraQ ID}. (
                        1;
                        600;
                        600;
                        600;
                        600;
                );
        IN      NS ns.{自分のtraQ ID}.

ns      IN      A 192.168.0.38
server  IN      A 192.168.0.129
```

:::


<!-- 講師用
```
$TTL 60
.               IN      SOA     ns.root. ns.root. (
                                3;
                                600;
                                600;
                                600;
                                600;
                        );
.               IN      NS      ns.root.
ns.root.        IN      A       {ルートネームサーバーのGlobal IP}
hijiki51.       IN      NS      ns.hijiki51.
ns.hijiki51.    IN      A       {受講者のGlobal IP}
``` -->


### 2. 名前解決をしてみよう
1.ができたら相手の HTTP サーバの IP アドレスを名前解決することで取得してみましょう。
ルートネームサーバーの IP は講師から共有されます。


:::details ヒント
DNS へのリクエストには`dig`コマンドを用います。
:::

:::details 答え
`dig @localhost server.{あなたのtraQ ID}`
:::