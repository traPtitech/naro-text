# Chapter8: DNS 1

この章ではDNS(Domain Name System)に関する基礎知識とその設定を学びます。

[[toc]]

## Lesson

### DNS
DNS(Domain Name System)はIPアドレスなどのリソースに対して別名をつけるアプリケーション層プロトコルです。

DNSはドメイン名空間(Domain Name Space)を管理するネームサーバーと名前解決を行うリゾルバからなります。

ドメイン名空間は一つのネームサーバーで一括管理されているのではなく、木構造のように複数ネームサーバーによって分散管理されており、名前解決を行う際は木構造の親から子へと再帰的に問い合わせを行います。

また、一番初めに問い合わせを行う、全体の木構造の親に当たるネームサーバーのことをルートネームサーバーと呼びます。

現在1600を超えるルートネームサーバーが起動しています。

https://root-servers.org/
## Assignment

::: info
nsはbind9導入済みUbuntuのインスタンスです。
:::
### 1. ネームサーバーを構築してみよう
ネームサーバーに`server.{あなたの traQ ID}`のレコードを登録してみましょう。

後々使うのでs1~s3のどれかのIPにしておくといいです。



:::details ヒント1
扱うゾーンは`{あなたの traQ ID}`になるでしょう。(TLDです)
:::

:::details ヒント2
今回はIPv4を用いているので設定するのは`A`レコードです。
:::

:::details ヒント3
「bind9 Aレコード 設定」などで検索してみるといいでしょう
:::

:::details 答え
NICには事前にIPアドレスを割り当てておきます。

まずはサーバーに割り振られたIPアドレスを確認します。（DHCPであるため手動確認）
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
1.ができたら相手のHTTPサーバのIPアドレスを名前解決することで取得してみましょう。
ルートネームサーバーのIPは講師から共有されます。


:::details ヒント
DNSへのリクエストには`dig`コマンドを用います。
:::

:::details 答え
`dig @localhost server.{あなたのtraQ ID}`
:::