# Vue.js で TodoList

`~/vue`ディレクトリの中にテンプレートリポジトリをクローンしてプログラムを書きます。

## Vue テンプレートのクローン

予め設定等が準備されたテンプレートリポジトリを用いて TodoList を作っていきます。

https://github.com/hijiki51/naro-template にアクセスし、「Use this template」→「Create a new repository」をクリックしてください。

![](assets/01.png)

「Repository name」にリポジトリ名を入力、公開状態は TA が見れるように「Public」 にしてください。

![](assets/02.png)

「Create repository from template」でリポジトリを作成したら手元にクローンしてください。  
`cd <リポジトリ名>`でプロジェクトのディレクトリに移動し、`code .`で VSCode を開きます。

:::warning
`command not found`と出る場合にはパスが通っていないので、TA に聞いてください。

// todo: リンクかなんか貼ってもよさそう
:::

開いたプロジェクトの中に入っている`package.json`というファイルには npm に関する様々な設定が書かれています。
この中には依存するパッケージ一覧も含まれており、以下のコマンドでそれらをインストールすることができます。

`$ npm i`  
もしくは  
`$ npm install`

```bash
mehm8128@DESKTOP-6F4C0KI ~/naro-lecture/naro-template (main)$ npm i

added 130 packages, and audited 131 packages in 2s

20 packages are looking for funding
  run `npm fund` for details

1 moderate severity vulnerability

To address all issues, run:
  npm audit fix

Run `npm audit` for details.
```

テンプレートは初期状態でビルド&配信できるようになっているんどえ、以下のコマンドを実行してブラウザで確認してみましょう。

`$ npm run dev`

```bash
mehm8128@DESKTOP-6F4C0KI ~/naro-lecture/naro-template (main)$ npm run dev

> todolist@0.0.0 dev
> vite


  vite v2.9.9 dev server running at:

  > Local: http://localhost:3000/
  > Network: use `--host` to expose

  ready in 158ms.
```

この状態で、ブラウザから http://localhost:3000/ にアクセスすると、以下のような画面が表示されるはずです。

// todo: warning

:::warning
WSL2 を使う人は、最初に書いた WSL2 の設定についてを実行しないとアクセスできません。
:::

:::warning
それでもアクセスできない場合は package.json の該当部分を以下のように変更してください。

```diff
"scripts": {
-    "dev": "vite",
+    "dev": "vite --host",
     "build": "vite build",
     "serve": "vite preview"
  }
```

その後、同様に起動し、今度は Network: に表示されるリンクを使ってアクセスしてみてください。
:::

![](assets/03.png)

止めるときは`Ctrl + C`(Mac の人は`⌘+C`)で止めてください。

## Vue.js 入門
