# Chapter8: DNS 2

この章ではリソースレコード及びドメインの移譲について学びます。

[[toc]]

## Lesson

### リソースレコード

リソースレコードは次の形式で設定されます。


|	Owner	|	TTL	|	Class	|	Type	|	RDATA	|
|	---	|	---	|	---	|	---	|	---	|
| example.com.	|	3600	|	IN	|	A	| 192.0.2.1 |

- Owner
  - リソースレコードがおかれているドメイン名
- TTL(Time To Live)
	- リソースレコードの有効時間(秒)
- Class
  - 16 ビットの値で、プロトコルファミリーを示す。
    - IN (インターネットシステム)
    - CH (カオスシステム)
- Type
  - このレコードのリソースの種類を示す。
- RDATA
  - 実際のリソースを記述する。

リソースレコードのタイプには様々なものがありますが、その一部を抜粋して以下の表に示します。

| Type  | 記述内容                                    |
| ----- | ------------------------------------------- |
| A     | ホストアドレス(IPv4)                        |
| AAAA  | ホストアドレス(IPv6)                        |
| CNAME | エイリアスの標準名(参照元)                  |
| TXT   | 任意の文字列                                |
| NS    | ゾーンを管理しているネームサーバー          |
| MX    | {優先度} {配送先メールサーバーのドメイン名} |

## Assignment

### 1. サブドメインを移譲してみよう
`{相手の traQ ID}.{あなたの traQ ID}`のレコードを相手のネームサーバーに移譲してみましょう。
移譲ができたらきちんと名前解決できるか確かめてみましょう。

::: details ヒント1
子側では新しく移譲されたゾーンを設定する必要があります。
:::

::: details ヒント2
親側では子のネームサーバーを登録する必要があります。
:::

::: details ヒント3
「bind9 dns delegation」などで調べるといいでしょう。
:::


### 2. サブドメインを移譲されてみよう
`{あなたの traQ ID}.{相手の traQ ID}`のレコードを相手のネームサーバーから移譲してもらいましょう。
移譲ができたらきちんと名前解決できるか確かめてみましょう。




### 解答例
::: details 答え

#### 親側の設定

```
$TTL 60
... (省略)
{子のtraQ ID}                  IN      A {子のネームサーバーのGlobal IP}

sub                            IN      NS ns.{子のtraQ ID}.{親のtraQ ID}.
ns.{子のtraQ ID}.{親のtraQ ID}. IN      A {子のネームサーバーのGlobal IP}
```
```
zone "{子のtraQ ID}.{親のtraQ ID}" IN {
        type master;
        file "/etc/bind/rr/{子のtraQ ID}";
};
```


#### 子側の設定
```
zone "{子のtraQ ID}.{親のtraQ ID}" IN {
        type master;
        file "/etc/bind/rr/{親のtraQ ID}";
};
```

```
$TTL 60
@       IN      SOA ns.{子のtraQ ID}.{親のtraQ ID}. ns.{子のtraQ ID}.{親のtraQ ID}. (
        2;
        600;
        600;
        600;
        600;
);

        IN      NS ns.{子のtraQ ID}.{親のtraQ ID}.
ns      IN      A {子のネームサーバーのGlobal IP}
```

:::