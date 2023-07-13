# 2022年度Webエンジニアになろう講習会 第6回 実習編

## 座学編メモ
### ブラウザセキュリティ入門
 - ブラウザユーザーは開発者よりもリテラシーが低い
 - それでも新しい機能は追加したい
 - 様々なセキュリュティ強化機能が搭載される
     - もともとは性善説的に

#### 攻撃例
 - クロスサイトスクリプティング(XSS)
     - 典型的なクライアント側での攻撃
     - 意図していない動作を起こせる
     - セッション ID などを抜き取られるとまずい
 - XSS の対策
     - サニタイジング
         - スクリプトを実行するような入力を拒否・無効化する
     - CSP(Content Security Policy), SRI
         - スクリプト実行範囲の許可を取ったり、実行されるスクリプトが想定されたものか検証する
         - [コンテンツセキュリティポリシー (CSP) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/CSP)
     - Cookie
         - 大事な情報を js で読み取れないようにする(XSS を実行された後でもなんとかなるように)

```js
// おこる
$div.innerHTML = "<script>alert(1)</script>"
// おこらない
$div.textContent = "<script>alert(1)</script>"
```

 - クリックジャッキング
     - `<iframe>`を使って別の URL のサイトを埋め込む
         - 悪意のあるサイトの上に真っ当なサイトを重ねることで意図しない操作をさせる。
     - 対策
         - `X-Frame-Option`
             - HTTP ヘッダー
         - `CSP`
         - `CORS(Cross-Origin Resource Sharing)` / `Same-Origin Policy`
             - 攻撃ではなく防御
             - ブラウザが別ドメインのサーバーへのリクエストを拒否
             - 開発上で一番出会う困るやつ
             - 事前にリクエスト先へアクセス許可を送っておく
             - [オリジン間リソース共有 (CORS) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)
             - `No 'Access-Control-Allow-Origin' Header`などのエラー

## 実習編目標
- API を利用するクライアントを書く
- 複数ページが存在するクライアントを書く

:::tip
第 3 回で作成したサーバーを利用します。先に第 3 回の内容を終わらせるようにしてください。
:::

:::warning
今日は下の[最重要課題](#最重要課題)を一番やってほしいです。
途中のコードなどはバシバシコピー&ペーストしてもらっても構いません。

最重要課題はこれまでのコードを上手く組み合わせることで実現できるはずです。
どのような API が必要か、それをどうやって表示すればいいかを自分で考えて作ってみましょう。
:::

# Vueのプロジェクトを作成する

[第二回](https://md.trap.jp/grYPeJzbSxWDfz5qHZfXYQ#%E3%83%AA%E3%83%9D%E3%82%B8%E3%83%88%E3%83%AA%E3%81%AE%E4%BD%9C%E6%88%90)と同じように https://github.com/hijiki51/naro-template のテンプレートリポジトリからリポジトリを作成します。
クローンしてきたディレクトリを VSCode で開いておきましょう。
`npm i`するのを忘れずに。

# Vue Routerの導入

SPA を作る際には、`PATH`に応じたページを描画する Router のような補助ライブラリが使用すると便利です。
今回は、公式 Router である[Vue Router v4.x](https://next.router.vuejs.org/)を使用します。

参考: [Routing | Vue.js](https://v3.vuejs.org/guide/routing.html)

## 1. ライブラリのインストール

`npm install vue-router@4` を実行してください。

## 2. Routerの設定

`PATH`と描画対象の関係である Route を定義します。
`src`以下に、`router.js`を以下の内容で作成してください。

<<<@/chapter2/section1/src/0/router.ts{typescript:line-numbers}

## 3. Vue Routerの使用

Vue Router を読み込むように`src/main.js`を以下のように変更します。

<<<@/chapter2/section1/src/0/main.ts{typescript:line-numbers}

次に、`src/App.vue`を以下のように変更します。

<<<@/chapter2/section1/src/0/App.vue{vue:line-numbers}

## 4. Homeページの作成

`src`直下に`pages`ディレクトリを作成し、`src/pages/HomePage.vue`を以下の内容で作成してください。

<<<@/chapter2/section1/src/0/HomePage.vue{vue:line-numbers}

## 5. NotFoundページの作成

`router.js`に定義した Route の配列は先頭からマッチします。

<<<@/chapter2/section1/src/0/routes.ts{typescript:line-numbers}

この後、皆さんにはいくつかのページとその`PATH`の対応を追加してもらうわけですが、どの`PATH`にもマッチしなかった場合、任意の`PATH`にマッチする`/:path(.*)`がマッチし、NotFound ページが表示されます。

`src/pages/NotFound.vue`を以下の内容で作成してください。

<<<@/chapter2/section1/src/0/NotFound.vue{vue:line-numbers}

# プロキシの設定
API へ接続するためにプロキシを設定します。
`localhost:3000/api/*`へのリクエストを、自分のサーバーの`133.130.109.224/*`へプロキシする設定を書きます。

:::info
# プロキシとは
`http://localhost`にアクセスすると、`npm run dev`で起動している手元のサーバーにリクエストが飛んで、`index.html`などの静的なファイルがレスポンスとして返ってきます。
ログインや街の情報の取得などは、リモートのサーバーに送信したいわけですが、`localhost`から配信されているクライアントから`133.130.109.224`に対してリクエストを送ろうとするとブラウザのセキュリティ機構に制限されたりと面倒なことがあります。
これらを避けるために、`133.130.109.224/*`へのリクエストを`localhost:3000/api/*`として手元のサーバーに送信し、手元のサーバーがブラウザの代わりにリモートのサーバーにリクエストを送信します。手元のサーバーは、リモートのサーバーからのレスポンスを透過的にブラウザに返却するので、ブラウザからはあたかも`localhost:3000/api/*`がレスポンスを返しているように見えます。
このように、何らかの目的のために代理で通信するサーバーをプロキシサーバーと言い、通信を代理させることを「プロキシする」と言います。

参考: [オリジン間リソース共有 (CORS) | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)
参考: [プロキシ | Wikipedia](https://ja.wikipedia.org/wiki/%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7)
:::

プロジェクトルートの`vite.config.js`というファイルの内容を、以下の内容に変更します。
ポート番号は自分がサーバーを起動しているポート番号にしてください。

<<<@/chapter2/section1/src/0/vite.config.ts{typescript:line-numbers}

:::tip
VSCode の Settings から`Format on Save`にチェックを入れると、自動できれいなコードに直してくれます。
:::

# 初期状態の確認

:::warning
PC にインストールされているセキュリティソフトによっては開発ページにアクセスできないことがあるようです。
その場合は、セキュリティソフトのファイアウォールを一時的に停止するか、ターミナルから`npm run dev`で起動した後、表示される IP アドレスでの URL にアクセスしてみてください(できない場合は TA に聞いてください)
:::

これまでと同様に`npm run dev`で起動して、以下のような画面が表示されていれば OK です。

![](images/0/vue_init.png)


# プロジェクト構成

以下のようなプロジェクト構成になっています。

```
.
├── README.md
├── index.html
├── package-lock.json
├── package.json
├── public
│   └── favicon.ico
├── src
│   ├── App.vue
│   ├── assets
│   │   └── logo.png
│   ├── components
│   │   └── HelloWorld.vue
│   ├── main.js
│   ├── pages                // Routerで出し分けるページ
│   │   ├── HomePage.vue
│   │   └── NotFound.vue
│   └── router.js            // Routeを定義するファイル
└── vite.config.js

5 directories, 13 files
```

## src/router.js
Vue Router を定義しているファイルです。

ページを追加したい場合は、`routes`の中にオブジェクトを追加していけば OK です。
参考: [Routes' Matching Syntax | Vue Router](https://next.router.vuejs.org/guide/essentials/route-matching-syntax.html)

`App.vue` に書かれている、`<router-view/>`コンポーネントが、このファイルで指定されたコンポーネントに置き換えられ描画されます。
参考: [`router-link` | Vue Router](https://next.router.vuejs.org/guide/#router-link)
参考: [`router-view` | Vue Router](https://next.router.vuejs.org/guide/#router-view)

参考: [Vue Router](https://next.router.vuejs.org/)

## pages以下

ページを表します。
