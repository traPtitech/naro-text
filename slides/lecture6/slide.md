---
marp: true
theme: SysAd
---

<!--
_class: title
-->

# セキュリティ入門

Webエンジニアになろう講習会 第6回

---

# 自己紹介

<div class="columns">
  <div>
 <img src="assets/icon.png" alt="Kentaro1043のアイコン" />
  </div>
  <div>
 <h2>Kentaro1043</h2>
 <div>数理・計算科学系</div>
 <div>インフラが好きです</div>
  </div>
</div>

---

# 前回のおさらい

// TODO

---

# 今日の内容は大事

外部にサービスを公開するときに、
絶対に気をつけて欲しいことです。
難しい内容ではないので、しっかり理解してください。

---

# 目次

- 座学
  - Webサービスセキュリティ入門
  - ブラウザセキュリティ入門
- 実習
  - vue-routerの設定
  - プロキシの設定
  - サーバーとの通信

---

# 目次

- 座学
  - <span style="color:red">Webサービスセキュリティ入門</span>
  - ブラウザセキュリティ入門
- 実習
  - vue-routerの設定
  - プロキシの設定
  - サーバーとの通信

---

# Webサービスセキュリティ入門

- サーバー内の情報が流出しないように気を付ける
  - 認証情報を守る
  - ポートを閉じる
  - ライブラリに頼る
- たとえ流出してもわからないようにする
  - ハッシュ化する
  - パスワードの痕跡を残さない
- 他の人の攻撃を手伝わないようにする
  - 不要なポートは閉じる
  - パッチを当てる

---

# 脆弱性は生まれるもの

- どれだけ注意しても少なからず脆弱性は存在する
- クリティカルな脆弱性を作らないことが重要
- どれだけプロトコルやライブラリがセキュア（堅牢）なものでも、開発者が理解していなければ脆弱性は生まれる

---

# Webサービスセキュリティ入門

- <span style="color:red">サーバー内の情報が流出しないように気を付ける</span>
  - ログイン情報を漏らさない
- たとえ流出してもわからないようにする
  - ハッシュ化する
- 他の人の攻撃を手伝わないようにする

---

# OWASP Top 10:2021

- アクセス制御の不備
- 不適切な暗号化
- インジェクション
- 設計の不備
- セキュリティ上の設定ミス
- 脆弱なコンポーネントの使用
- 識別・認証の不備
- ソフトウェアとデータの整合性の問題
- セキュリティログとモニタリングの失敗
- サーバーサイドリクエストフォージェリ

---

# ログイン情報の流出

- サーバーはインターネット上に公開されている
  - 常に攻撃の可能性にさらされている
- ログインされたら何でもできる
- 対策
  - ファイヤーウォールの導入
  - 公開鍵認証方式の導入
  - Fail2banの導入
  - 適切なロギング・モニタリング
  - sudoerの制限

---

# ログイン情報の流出

```log
Jun 16 11:40:26 m011 sudo: hijiki51 : TTY=pts/1 ; PWD=/home/hijiki51 ; USER=root ; COMMAND=/usr/bin/cat /var/log/auth.log
Jun 16 11:40:26 m011 sudo: pam_unix(sudo:session): session opened for user root by hijiki51(uid=0)
Jun 16 11:40:27 m011 sudo: pam_unix(sudo:session): session closed for user root
Jun 16 11:40:50 m011 sshd: Invalid user tickets from 27.101.120.70 port 51846
Jun 16 11:40:51 m011 sshd: Received disconnect from 27.101.120.70 port 51846:11: Bye Bye [preauth]
Jun 16 11:40:51 m011 sshd: Disconnected from invalid user tickets 27.101.120.70 port 51846 [preauth]
Jun 16 11:41:05 m011 sshd: Received disconnect from 201.149.55.226 port 36193:11: Bye Bye [preauth]
Jun 16 11:41:05 m011 sshd: Disconnected from authenticating user root 201.149.55.226 port 36193 [preauth]
Jun 16 11:41:07 m011 sshd: Invalid user qt from 43.156.60.74 port 55154
Jun 16 11:41:07 m011 sshd: Received disconnect from 43.156.60.74 port 55154:11: Bye Bye [preauth]
Jun 16 11:41:07 m011 sshd: Disconnected from invalid user qt 43.156.60.74 port 55154 [preauth]
Jun 16 11:41:09 m011 sshd: Received disconnect from 154.117.199.12 port 40538:11: Bye Bye [preauth]
Jun 16 11:41:09 m011 sshd: Disconnected from authenticating user root 154.117.199.12 port 40538 [preauth]
Jun 16 11:41:48 m011 sudo: hijiki51 : TTY=pts/1 ; PWD=/home/hijiki51 ; USER=root ; COMMAND=/usr/bin/less /var/log/auth.log
Jun 16 11:41:48 m011 sudo: pam_unix(sudo:session): session opened for user root by hijiki51(uid=0)
```

### /var/log/auth.log などで確認できる

---

# 権限は必要最小限に

ポートも、アクセス権も、できることが増えると
脆弱性は生まれやすくなる

---

# DBに直接アクセスされる

- データベースの認証情報が流出
  - 常に攻撃の可能性にさらされている
- DBの中身を見放題・書き換え放題
  - ユーザーの個人情報が…
  - 各種認証情報も…
  - もし金銭を扱っていたら…
- 対策
  - 適切な認証情報の管理
  - 適切なアクセスコントロール
    - 権限設定をちゃんとする

---

# SQLi, OS Command Injection, etc

- 巧妙なリクエストを送って、不正にデータを取得したり、権限昇格をしたりする
- 対策
  - サニタイジングを行う
  - OS側の機能を利用しない

---

# SQL injection

```go
cities := []City{}
db.Select(&cities, fmt.Sprintf("SELECT * FROM city WHERE CountryCode='%s'", os.Args))

fmt.Println("日本の都市一覧")
for _, city := range cities {
    fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
}
```

### こういうコードを書くと……

---

# SQL injection

**正常な入力**

```bash
$ go run main.go JPN
Connected!
日本の都市一覧
都市名: Tokyo, 人口: 7980230人
都市名: Jokohama [Yokohama], 人口: 3339594人
都市名: Osaka, 人口: 2595674人
都市名: Nagoya, 人口: 2154376人
...
```

**攻撃**

```bash
$ go run main.go "' OR 1 OR ''='"
Connected!
日本の都市一覧
都市名: Kabul, 人口: 1780000人
都市名: Qandahar, 人口: 237500人
都市名: Herat, 人口: 186800人
都市名: Mazar-e-Sharif, 人口: 127800人
...
```

---

# SQL injection

### 正しくはこう

```go
cities := []City{}
// プレースホルダを使う
db.Select(&cities, "SELECT * FROM city WHERE CountryCode=?", os.Args)

fmt.Println("日本の都市一覧")
for _, city := range cities {
    fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
}
```

---

# Webサービスセキュリティ入門

- サーバー内の情報が流出しないように気を付ける
  - 認証情報を守る
  - ポートを閉じる
  - ライブラリに頼る
- <span style="color:red">たとえ流出してもわからないようにする</span>
  - ハッシュ化する
  - パスワードの痕跡を残さない
- 他の人の攻撃を手伝わないようにする
  - 不要なポートは閉じる
  - パッチを当てる

---

# 流出しても、認証情報は漏れないようにする

- どんなに注意しても流出が起きてしまうことはある
  - どんなに大きな企業でもやらかすときはやらかす
- 流出しても機密データは漏らさないように
  - 機密データはサーバーに保存しない

---

# パスワードのハッシュ化

- データベースに平文（そのままの状態）のパスワードを保存するのはNG
  - 漏洩すると大惨事になる
- パスワードを元に一定の手順に従って求めたハッシュ値を代わりに保存する
- 同じパスワードからは同じハッシュ値が得られるので、パスワードを保存せずにパスワードが検証できる
- 必ずハッシュ化する
  - bcrypt, PBKDF2, scrypt, Argon2などがよく用いられる
- 単なるハッシュ化では、事前計算したハッシュの結果と比較することでパスワードが復元できてしまうことがある
  - ソルト（ランダムな値）をパスワードに付加してからハッシュ化することで事前計算を困難に出来る

---

# Webサービスセキュリティ入門

- サーバー内の情報が流出しないように気を付ける
  - 認証情報を守る
  - ポートを閉じる
  - ライブラリに頼る
- たとえ流出してもわからないようにする
  - ハッシュ化する
  - パスワードの痕跡を残さない
- <span style="color:red">他の人の攻撃を手伝わないようにする</span>
  - 不要なポートは閉じる
  - パッチを当てる

---

# 他の人の攻撃を手伝わないために

- DDoS攻撃
  - 複数人で一斉に攻撃
  - バックドアやBotを仕込まれて知らぬ間に自分が攻撃者になる
- リフレクション攻撃
  - サーバーの設定をちゃんとする
  - NXNSAttack
- 単純に踏み台サーバーとして利用されることも

---

# 参考資料

- [安全なウェブサイトの運用管理に向けての20ヶ条 - IPA](https://www.ipa.go.jp/security/vuln/20point-website-operation.html)
- [安全なウェブサイトの作り方 – IPA](https://www.ipa.go.jp/security/vuln/websecurity.html)
- [脆弱性対策 - IPA](https://www.ipa.go.jp/security/vuln/index.html)

---

# 目次

- 座学
  - Webサービスセキュリティ入門
  - <span style="color:red">ブラウザセキュリティ入門</span>
- 実習
  - クライアントからAPIを呼び出す

---

# ブラウザのセキュリティ

- さっきまではサーバーのセキュリティに関して
  - ブラウザ自体にも重要なデータが保存されている
- 加えてブラウザ自体も各端末で動くネイティブアプリケーション
  - 攻撃者と端末とのインターフェースとなりうる

---

# 保存場所の一例

（ブラウザの開発者ツール `Application` タブのスクリーンショット）

- **Storage**
  - Local Storage
  - Session Storage
  - IndexedDB
  - Web SQL
  - Cookies
- **Cache**
  - Cache Storage

---

# ブラウザの持つ情報

- 悪意のあるWebページが、Local Storageに保存した他のWebページに関する情報を無制限に読み取れるなら…
- 悪意のあるWebページAがiframeを用いて悪意のないWebページBを埋め込んでいる時、WebページAのJavaScriptが無制限にWebページBのDOMにアクセスできるなら…

---
**→ 大切なデータを攻撃者に知られてしまう**

---

# 境界の必要性

- Webページの間には「ほどよい」境界が必要
  - どのような境界が適切か？
- **Origin**: URI中の「スキーム、ホスト、ポート」の組
  - **Scheme**: HTTP, HTTPS
  - **Host**: a.example.com, b.example.com, a.sample.com
  - **Port**: 80, 443, 8080

---

# Same-Origin Policy

- SOP: Originに依拠した最も基本的なセキュリティ機構
- 次の2つのルールに従って、Webページ間の「やりとり」に制限を加える
  - 2つのページの Origin が一致していれば、無制限で「やりとり」を許す
  - 2つのページの Origin が異なっていれば、「やりとり」を原則禁止する
- やりとり：書き込み、埋め込み、読み込み
  - Cross-Originでの書き込み：基本的に許可（条件付きで禁止される）…「単純な」リクエストに限る
  - Cross-Originでの埋め込み：許可
  - Cross-Originでの読み込み：基本的に禁止される

---

# Cross-Origin Resource Sharing

```
Access to XMLHttpRequest at 'https://play.google.com/log?...' from origin
'https://drive.google.com' has been blocked by CORS policy: No
'Access-Control-Allow-Origin' header is present on the requested resource.

Failed to load resource: net::ERR_FAILED
```

- CORS: SOPの緩和
- SOPは時に善良な開発者の障害となる
- HTTPヘッダのやり取りを通じて、リソースの読み出し・書き込みを行ってよいかすり合わせる

---

# CORSの仕組み (1/5)

1. ユーザーのブラウザが `http://a.example.com` のWebページにアクセスします。

![CORS Diagram 1](https://via.placeholder.com/600x300.png/FFFFFF/000000?text=User+->+a.example.com)

---

# CORSの仕組み (2/5)

2. `a.example.com` のページで実行されたスクリプトが、別のオリジンである `http://b.example.com` にリクエストを送信しようとします。

![CORS Diagram 2](https://via.placeholder.com/600x300.png/FFFFFF/000000?text=a.example.com+Script+->+b.example.com)

---

# CORSの仕組み (3/5)

3. しかし、オリジンが異なるため、ブラウザはSame-Origin Policyに基づき、このリクエストをブロックします。

![CORS Diagram 3](https://via.placeholder.com/600x300.png/FFFFFF/000000?text=Browser+BLOCKS+Request)

---

# CORSの仕組み (4/5)

4. これを解決するため、`http://b.example.com` サーバー側で、予め「`http://a.example.com` からのリクエストを許可する」と設定しておきます。（`Access-Control-Allow-Origin` ヘッダなど）

![CORS Diagram 4](https://via.placeholder.com/600x300.png/FFFFFF/000000?text=b.example.com+sets+CORS+policy)

---

# CORSの仕組み (5/5)

5. ブラウザは `b.example.com` の許可設定を確認し、リクエストを許可します。これで異なるオリジン間での通信が安全に実現できます。

![CORS Diagram 5](https://via.placeholder.com/600x300.png/FFFFFF/000000?text=Request+is+now+Allowed!)

---

# Cross-Site Scripting (XSS)

- SOPは同一Origin内のリソースのやり取りに関して何ら制限を加えない
- 攻撃者が何らかの方法で他のユーザーの閲覧ページに任意のScriptを挿入することができれば？
  - 自身のサーバーに読み取った値を送信するような事ができる
- 対策
  - ユーザーが生成したコンテンツは必ずサニタイジングする
  - 危ない関数を迂闊に使わない
    - v-htmlディレクティブ
    - render関数のdomPropsやdomPropsInnerHTML

---

# Content Security Policy

- CSP: XSS脆弱性の水際対策
- 開発者は自身がブラウザ上で実行したいスクリプトを知っている
  - それだけ読み込めれば良い
- Inline Scriptを`.js`や`.css`にまとめて、信頼のできるOriginからのみ配信する
  - それ以外のScriptはContent-Injectionによるもの
- 信頼性担保にはSRI (SubResource Integrity)などが利用される

---

# Cookie

**HttpOnly属性**

- セッションの実現だけに用いられるCookieはブラウザの実行するJavaScriptから参照する必要がない
  - サーバーへの送信時以外参照できないように設定する

**Secure属性**

- HTTPSプロトコル上の暗号化されたリクエストでのみサーバーに送信できるようになる

---

# まとめ

---

# Webサービスセキュリティ入門

- サーバー内の情報が流出しないように気を付ける
  - 認証情報を守る
  - ポートを閉じる
  - ライブラリに頼る
- たとえ流出してもわからないようにする
  - ハッシュ化する
  - パスワードの痕跡を残さない
- 他の人の攻撃を手伝わないようにする
  - 不要なポートは閉じる
  - パッチを当てる

---

# 境界の必要性

- Webページの間には「ほどよい」境界が必要
  - どのような境界が適切か？
- **Origin**: URI中の「スキーム、ホスト、ポート」の組
  - **Scheme**: HTTP, HTTPS
  - **Host**: a.example.com, b.example.com, a.sample.com
  - **Port**: 80, 443, 8080

---

# Same-Origin Policy

- SOP: Originに依拠した最も基本的なセキュリティ機構
- 次の2つのルールに従って、Webページ間の「やりとり」に制限を加える
  - 2つのページの Origin が一致していれば、無制限で「やりとり」を許す
  - 2つのページの Origin が異なっていれば、「やりとり」を原則禁止する
- やりとり：書き込み、埋め込み、読み込み
  - Cross-Originでの書き込み：基本的に許可（条件付きで禁止される）…「単純な」リクエストに限る
  - Cross-Originでの埋め込み：許可
  - Cross-Originでの読み込み：基本的に禁止される

---

# Cross-Origin Resource Sharing

- CORS: SOPの緩和
- SOPは時に善良な開発者の障害となる
- HTTPヘッダのやり取りを通じて、リソースの読み出し・書き込みを行ってよいかすり合わせる

---

# Content Security Policy

- CSP: XSS脆弱性の水際対策
- 開発者は自身がブラウザ上で実行したいスクリプトを知っている
  - それだけ読み込めれば良い
- Inline Scriptを`.js`や`.css`にまとめて、信頼のできるOriginからのみ配信する
  - それ以外のScriptはContent-Injectionによるもの
- 信頼性担保にはSRI (SubResource Integrity)などが利用される

---

# Cookie

**HttpOnly属性**

- セッションの実現だけに用いられるCookieはブラウザの実行するJavaScriptから参照する必要がない
  - サーバーへの送信時以外参照できないように設定する

**Secure属性**

- HTTPSプロトコル上の暗号化されたリクエストでのみサーバーに送信できるようになる

---

# 目次

- 座学
  - Webサービスセキュリティ入門
  - ブラウザセキュリティ入門
- 実習
  - <span style="color:red">クライアントからAPIを呼び出す</span>
