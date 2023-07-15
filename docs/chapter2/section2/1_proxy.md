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

<<<@/chapter2/section2/src/1/vite.config.ts{typescript:line-numbers}

:::tip
VSCode の Settings から`Format on Save`にチェックを入れると、自動できれいなコードに直してくれます。
:::

# 初期状態の確認

:::warning
PC にインストールされているセキュリティソフトによっては開発ページにアクセスできないことがあるようです。
その場合は、セキュリティソフトのファイアウォールを一時的に停止するか、ターミナルから`npm run dev`で起動した後、表示される IP アドレスでの URL にアクセスしてみてください(できない場合は TA に聞いてください)
:::

これまでと同様に`npm run dev`で起動して、以下のような画面が表示されていれば OK です。

<!-- TODO:ここに写真 -->


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
