# Vue.js で TodoList

`~/vue`ディレクトリの中にテンプレートリポジトリをクローンしてプログラムを書きます。

## Vue テンプレートのクローン

// todo: ts 対応したのを traPtitech に別で作る

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

//todo: 色やばい

テンプレートは初期状態でビルド&配信できるようになっているので、以下のコマンドを実行してブラウザで確認してみましょう。

`$ npm run dev`

```bash
mehm8128@DESKTOP-6F4C0KI ~/naro-lecture/todolist (main)$ npm run dev

> todolist@0.0.0 dev
> vite


  VITE v4.3.8  ready in 611 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h to show help

```

この状態で、ブラウザから http://localhost:5173/ にアクセスすると、以下のような画面が表示されるはずです。

:::warning
WSL2 を使う人は、最初に書いた WSL2 の設定についてを実行しないとアクセスできません。

//todo リンク
:::

![](assets/03.png)

止めるときは`Ctrl + C`(Mac の人は`⌘+C`)で止めてください。

## Vue.js 入門

### Vue.js とは

以下のリンクから公式ドキュメントに飛ぶことができます。
[Vue.js](https://v3.ja.vuejs.org/)

traP では、Web フロントフレームワークとして最も多く使われているフレームワークで、traQ、traPortal、Showcase、anke-to、knoQ などで使われています。

### `.vue`ファイルについて

Vue.js では`.vue`という拡張子で単一ファイルコンポーネント(SFC, Single File Component)を作ることができます。

なろう講習会の言葉で言うと、Vue では、**一つの同じファイルに構造(HTML)・ロジック(JavaScript)・スタイル(CSS)**を記述することができます。それぞれを別の巨大なファイルに書くのではなく、**見た目に対応した要素ごとに分割して書く**ことで、それぞれの責任範囲をより直感的な形式で分けることができるわけです。このように分けられた要素をコンポーネントといいます。

//todo: .html でも全部書けるのでは？

### Vue.js の書き方

`<script>`タグ内にロジック
`<template>`タグ内に構造
`<style>`タグ内にスタイル
を記述します。

#### Sample.vue

<<< @/chapter1/section2/src/Sample.vue

#### 使用例

traQ で一つ例を挙げると、メッセージの表示部分はコンポーネントとして定義されています。
メッセージも複数のコンポーネントから構成されています。

https://github.com/traPtitech/traQ_S-UI/blob/master/src/components/Main/MainView/MessageElement/MessageElement.vue

- メッセージ情報の表示(ユーザー名、画像、ファイル等)の構造(HTML)
- メッセージにスタンプをつける等のロジック(TypeScript)
- それらのスタイルを記述するスタイル(Scss)

が一つのファイルに纏められていることがわかります(このファイルはたった 130 行程度ですが、traQ にはコンポーネントが 300 個以上あります。 これがそれぞれの HTML、 CSS、 JavaScript のファイルに書かれていると想像してみると…)。

### プロジェクト構造

```
.
├── index.html              // 最初にブラウザに読まれるHTMLファイル
├── node_modules/           // 依存ライブラリの保存先
├── package-lock.json       // 依存ライブラリ(詳細)
├── package.json            // 依存ライブラリ・タスク・各種設定
├── public
│   └── favicon.ico         // 静的ファイル(ビルドされない)
├── src
│   ├── App.vue             // main.jsから読まれる.vueファイル(Vueの処理開始点)
│   ├── assets              // Vueで使用したい画像など
│   │　　└── logo.svg
│   ├── components          // 各種コンポーネント
│   │　　└── HelloWorld.vue
│   └── main.ts　　　　　　　　// index.htmlから読まれるscript(TSの処理開始点)
└── vite.config.ts
```

#### `index.html`

Vite はここから参照されているファイルをたどってビルドを進めていきます。
ここに必要なものを書き加えることもありますが、基本的には書き換えません
マウント用の<div id="app"></div>と main.js の読み込みが書かれています。

#### `node_modules`

`npm install` でインストールされる依存ライブラリが保存されるディレクトリです。
中を見ることは殆ど無いです。
.gitignore に指定されています。(`package.json`, `package-lock.json` があれば `npm install` で再現できるためです)

#### `package.json` `package-lock.json`

プロジェクトに関するメタ情報や、`npm run ~~`で実行できるタスク、依存ライブラリの情報が記述されています。

#### `src/main.ts`

`index.html` で読み込まれている ts ファイルです。
ここでは Vue インスタンスを生成し、`index.html` の`<div id="app"></div>`部分にマウントしています。

#### `src/App.vue`

Vue.js としてのエントリーポイントです。
HelloWorld コンポーネントを読み込み → 登録 → 描画しています。

#### `src/components/HelloWorld.vue`

ここと似たようなものをどんどん書いていきます。
App.vue で呼び出されています。  
※`components`内に他にも色々なコンポーネントがありますが、重要ではないので無視します。

## Vue.js を書く準備

まず、以下の拡張機能をインストールしてください。

#### Vue Dev tool

Chrome Devtool に Vue.js 向けのデバッグ機能を追加してくれます。  
[Vue.js devtools - Chrome ウェブストア](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd?hl=ja)

//todo: ほとんど使ったことないけどどう使うんだっけ

#### ESLint

前回の「きれいなソースコードを保つために」で紹介した VSCode の拡張機能。  
保存時にフォーマットする設定にしておくと精神衛生が保たれます。  
[ESLint - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)

#### Prettier

同じく前回の「きれいなソースコードを保つために」で紹介した VSCode の拡張機能。  
保存時にフォーマットする設定にしておくと精神衛生が保たれます。  
[Prettier - Code formatter - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)

#### Volar

VSCode の Vue3 向けの統合プラグイン。  
[Vue Language Features (Volar) - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vue.volar)

インストールが終わったら、反映させるために VSCode を一度閉じて開きなおしてください。

### ソースコードの書き進め方

`npm run dev`で起動していれば、ファイルの変更を自動で検知して表示が更新されます。

:::tip
ちゃんと保存しましょう。
![](assets/04.png)

画面上部のタブのファイル名の横に ● がついているときは保存できていません。

![](assets/05.png)

設定で自動保存されるようにしておくと便利です。  
参考： [自動保存するように設定する](https://www.javadrive.jp/vscode/setting/index2.html)
:::

## Vue.js を書く

開発基礎講習会で書いたカウンターのソースコードを再掲します。

#### index.html(一部抜粋)

<<< @/chapter1/section2/src/index.html

#### counter.js

<<< @/chapter1/section2/src/counter.js

### Vue.js を書く

:::info
先にコードを書いてから解説を書いています。  
意味がわからなくてもとりあえずコピペ or 写経しましょう。
:::

#### ファイルの作成

`components`ディレクトリ内に`ClickCounter.vue`というファイルを作成します。

![](assets/06.png)

#### ソースコードの変更

#### src/App.vue

`style`タグを丸ごと消します

<<< @/chapter1/section2/src/App.vue

##### src/components/HelloWorld.vue

`script`タグ内で`ClickCounter.vue`を読み込み、`template`タグ内にカウンターを配置します。

<<< @/chapter1/section2/src/HelloWorld.vue

##### src/components/ClickCounter.vue

<<< @/chapter1/section2/src/ClickCounter.vue

以下のように動けば OK です。

![](assets/01.gif)

### ソースコード解説

#### src/components/HelloWorld.vue

##### 9~14 行目

テンプレート部分です。  
Vue のコンポーネントは一つのタグの中に収まっている必要があります。  
そのため、多くの場合 div タグで囲まれています。(`ClickCounter.vue`も)

##### 2 行目

```ts
import ClickCounter from "./ClickCounter.vue"
```

`ClickCounter` コンポーネントを読み込む部分です。

##### n 行目

```ts
defineProps<{
	msg: string
}>()
```

`msg`props を`string`型で定義してる部分です。  
今回だと`App.vue`で `<HelloWorld msg="Hello Vue 3 + Vite" />`のような形で`msg`に値を指定することで、コンポーネントを使う側から値を渡しています。 JavaScript でいう関数の引数のようなものです。

参考: [プロパティ | Vue.js](https://ja.vuejs.org/guide/components/props.html)

##### 12 行目

```html
<ClickCounter />
```

読み込んだコンポーネントを利用しています。

#### src/components/ClickCounter.vue

##### 4 行目

コンポーネント内で利用する変数をこのように書きます。  
ここでは`count`という名前の変数を`number`型で定義しています(実はこの程度なら TypeScript の型推論というものが効いて、初期値の`0`から自動で`count`変数は`number`型だと推論してくれます)。

```ts
const count = ref<number>(0)
```

参考: [ref によるリアクティブな変数 | Vue.js](https://v3.ja.vuejs.org/guide/composition-api-introduction.html#ref-%E3%81%AB%E3%82%88%E3%82%8B%E3%83%AA%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%95%E3%82%99%E3%81%AA%E5%A4%89%E6%95%B0)

##### 11・12 行目

ボタンが押されたイベントに対する処理を書いています。  
`@click`の中には直接 JavaScript を書くこと(今回)も`<script setup>`内で定義した関数を指定することもできます。

```html
<button @click="count++">クリック！</button>
<button @click="count = 0">リセット！</button>
```

参考: [イベントへの入門 - ウェブ開発を学ぶ | MDN](https://developer.mozilla.org/ja/docs/Learn/JavaScript/Building_blocks/Events)  
参考: [イベントハンドリング | Vue.js](https://v3.ja.vuejs.org/guide/events.html)

:::tip
v-on:click のショートハンドとして@click という書き方ができます(推奨)
:::

##### 5 行目・10 行目

`computed`という機能を使って、表示するメッセージを生成します。

参考: [テンプレート構文 | Vue.js](https://v3.ja.vuejs.org/guide/template-syntax.html#%E3%83%86%E3%83%B3%E3%83%95%E3%82%9A%E3%83%AC%E3%83%BC%E3%83%88%E6%A7%8B%E6%96%87)  
参考: [算出プロパティ | Vue.js](https://v3.ja.vuejs.org/guide/reactivity-computed-watchers.html#%E7%AE%97%E5%87%BA%E3%83%95%E3%82%9A%E3%83%AD%E3%83%8F%E3%82%9A%E3%83%86%E3%82%A3)

### Vue.js の嬉しさを実感する

いかがでしょうか？  
この規模だと、ソースコードの行数としては生の HTML・JS で書いたほうが少ないですが、書く量を増やしてでも Vue を使う嬉しさがあります。

#### 変数を操作するだけで表示が変わる

生 JS ではボタンのクリックごとに表示を変更する必要があります。  
Vue ではコンポーネントの定義の時に表示と変数を関連付ける(5 行目・10 行目)ことで、変数を操作するだけで表示が自動で切り替わります。これが Vue.js の提供するリアクティビティです(コンポーネントが内部に保持している状態を変更したとき、その変更が Vue によって検知されて自動で HTML に反映されます)。

今回は、変数を操作する箇所が少ないのでまだ追えますが、この変数が様々なところで変更されるものだった場合を考えてみましょう。  
生 HTML・JS ではその全ての場所で書き換えるコードを忘れることなく書かなければなりません。  
複数のプログラマでコードを書いた場合、それを忘れないようにするということはかなりのコストになってしまうので辛いです。  
Vue.js ではそれがなくて嬉しいです。

### 一度コンポーネントを登録すれば使い回せる

カウンターを 2 つ作りたくなった場合を考えます。  
生 HTML・JS が書ける人はちょっとチャレンジしてみてください。関数名や変数名をかぶらないようにしたり、セレクタの名前を変更したりと結構めんどくさいです。  
Vue.js ならば、`HelloWorld.vue` の`<ClickCounter />`をコピーして増やすだけで OK です。  
これは traQ のように同じ要素を沢山利用するような Web アプリで大きな利点となります。

## さらに Vue.js を書く

商品リストをテーマに、Todo リストに必要な Vue.js に機能をピックアップしていきます。

こんな感じのを作っていきます。
![](assets/03.gif)

### 必要な要素を考える

上の Gif のようなアプリを実現するためには何が必要か考えてみましょう。

- 商品リストのコンポーネントを作る
- 商品のリストデータを保存する
- 商品のリストデータを表示する
- 商品を追加できる
- 商品の値段が 500 円以上だったら赤くする
- 商品の値段が 1000 円以上だったら「高額商品」と表示する

こんな感じでしょうか。  
それでは上から順番に実装していきましょう。

### 商品リストのコンポーネントを作る

`components`ディレクトリに`ItemList.vue`というファイルを作成します。
![](assets/07.png)

#### src/components/ItemList.vue

中身はコンポーネントに最低限必要な部分だけ書きます。

<<< @/chapter1/section2/src/ItemList.vue

#### HelloWorld.vue

<<< @/chapter1/section2/src/HelloWorld2.vue

表示されました。
こうすることで、後は`ItemList.vue`の中身を書き変えればよくなります。

![](assets/08.png)

### 商品のリストデータを保存する

商品リストのデータを保存するのに適当な変数の型は何でしょうか？  
商品「リスト」なので配列がよさそうです。  
というわけで、配列を使ってデータを保持することにします。  
今は商品の追加ができないので、とりあえずダミーデータを入れておきます。

参考: [Array | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)  
//todo: オブジェクトのリンクもほしいかも

<<< @/chapter1/section2/src/ItemList2.vue

4~7 行目は TypeScript の記法で、`Item`という型を`interface`を用いて定義しています。  
そして ref のジェネリクスに`Item[]`を渡すことで、`items`変数を`Item`型の配列の`ref`として扱えるようにしています。

//todo: interface とジェネリクスの参考リンク

### 商品のリストデータを表示する

先ほど定義したリストの情報を表示していきます。  
Vue.js ではリストデータを`template`タグ内で for 文のように書く `v-for` という構文があります。  
`v-for` を使うときには`:key`を設定しなければいけません(理由(やや難): [key | Vue.js](https://v3.ja.vuejs.org/api/special-attributes.html#key))。

参考: [リストレンダリング | Vue.js](https://v3.ja.vuejs.org/guide/list.html#%E3%83%AA%E3%82%B9%E3%83%88%E3%83%AC%E3%83%B3%E3%82%BF%E3%82%99%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%99)

これを使ってデータを表示してみます。

```tsx
<template>
  <div>
    <div>ItemList</div>
    <ul>
      <li v-for="item in items" :key="item.name">
        <div>名前: {{ item.name }}</div>
        <div>{{ item.price }} 円</div>
      </li>
    </ul>
  </div>
</template>
```

表示できました。

![](assets/09.png)

### 商品を追加する

Vue.js では入力欄に入力された文字列とコンポーネントの変数を結びつけることができます。  
参考： [フォーム入力バインディング | Vue.js](https://v3.ja.vuejs.org/guide/forms.html#form-input-bindings)

これを使って商品を追加できるようにしてみます。

<<< @/chapter1/section2/src/ItemList3.vue

参考: [アロー関数式 | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Functions/Arrow_functions)

できました！

![](assets/02.gif)

:::info
このままだとボタンを連打して商品の追加ができてしまいます。

- ボタンを押したら入力欄を空にする機能
- 入力欄が空だったらボタンを押しても追加されないようにする機能

を追加してみましょう。
:::

### 商品の値段が 500 円以上だったら赤くする

Vue.js では、ある特定の条件が満たされた時に class を追加するという機構を持たせることできます。  
これを使って、条件が満たされたときだけ CSS を当てるといったことができます。

参考: [CSS の基本 | MDN](https://developer.mozilla.org/ja/docs/Learn/Getting_started_with_the_web/CSS_basics)  
参考: [オブジェクト構文 | Vue.js](https://v3.ja.vuejs.org/guide/class-and-style.html#%E3%82%AA%E3%83%95%E3%82%99%E3%82%B7%E3%82%99%E3%82%A7%E3%82%AF%E3%83%88%E6%A7%8B%E6%96%87)

```vue
<template>
	<div>
		<div>ItemList</div>
		<ul>
			<li
				v-for="item in items"
				:key="item.name"
				:class="{ over500: item.price >= 500 }"
			>
				<div>名前: {{ item.name }}</div>
				<div>{{ item.price }} 円</div>
			</li>
		</ul>
		<div>
			<label>
				名前
				<input v-model="newItemName" type="text" />
			</label>
			<label>
				価格
				<input v-model="newItemPrice" type="number" />
			</label>
			<button @click="addItem">add</button>
		</div>
	</div>
</template>

<style>
.over500 {
	color: red;
}
</style>
```

![](assets/10.png)

### 商品の値段が 10000 円以上だったら「高額商品」と表示する

Vue.js では、ある特定の条件を満たした場合のみ、対象コンポーネントを表示するという機能を`v-if`という構文を使って実現することができます。

参考: [条件付きレンダリング | Vue.js](https://v3.ja.vuejs.org/guide/conditional.html#%E6%9D%A1%E4%BB%B6%E4%BB%98%E3%81%8D%E3%83%AC%E3%83%B3%E3%82%BF%E3%82%99%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%99)

これを使って商品の値段が 10000 円以上だったら「高額商品」と表示するという機能を実現してみましょう。

```vue
<template>
	<div>
		<div>ItemList</div>
		<ul>
			<li
				v-for="item in items"
				:key="item.name"
				:class="{ over500: item.price >= 500 }"
			>
				<div>名前: {{ item.name }}</div>
				<div>{{ item.price }} 円</div>
				<div v-if="item.price >= 10000">高額商品</div>
			</li>
		</ul>
		==略==
	</div>
</template>
```

![](assets/11.png)

これで商品リストが完成しました！

## Todo リストを作る

ここまで紹介してきた機能を使うことで Todo リストが作れるはずです。
頑張りましょう！

:::info
Todo リストを作りましょう。

必要な機能は以下の通りです。

- タスクは未完または完了済みの状態を持つ
- タスクはタスク名を持つ
- 未完タスクのリストと完了済みタスクのリストが表示される
- タスクを完了させることができる
- タスクを追加することができる

以上の機能が実現されていれば後は自由です。
スタイルが気になる人は CSS なども書きましょう。
:::

## ビルドする

`npm run build`でビルドを行うことができます。

Vue.js などのフレームワークの記法に従って書かれたコードは、そのままではブラウザ上で動きません。  
Vite のようなバンドラーによって、

- 依存関係の解決
- HTML/CSS/JS への変換
- 圧縮
- etc...

など様々な処理を加えられた後、いい感じにブラウザ上で動作する生の HTML/CSS/JS として出力されます。

ビルドによる成果物は`dist`ディレクトリの中に生成されています。
//todo: 画像貼る

## 公開する

それでは公開しましょう！

早速 push して公開したいところですが、`.gitignore`をみてみると`dist`ディレクトリが ignore されていることがわかります。
//todo: 画像貼る

これは GitHub にはソースコードだけをアップロードし、そのソースコードから再現できるものは極力アップロードしない(Git のパフォーマンスに影響するため)という考えから来ているものです。

`node_modules`も同じような理由で ignore されていることがわかると思います。

ビルド済みの成果物を GitHub Pages などを用いて公開してもいいのですが、上記の考えに従い、今回は違ったサービスで公開しようと思います。

### Vercel で公開する

[Vercel: Develop. Preview. Ship. For the best frontend teams](https://vercel.com/)

Vercel を使うと、ビルドが必要なサイトも簡単に公開することができるので便利です。

// todo:アカウント作るのめんどくさいので後回し

1. Signup から以下の画面に進み、GitHub アカウントと連携してください。
   ![](https://md.trap.jp/uploads/upload_639f66f4a91154672e52ea41f770bc54.png)

2. [新規作成画面](https://vercel.com/new)から==Add GitHub Account==を選択してください。
   ![](https://md.trap.jp/uploads/upload_2ca44b1f8fe1577bde24865b48b80f3c.png)
   ![](https://md.trap.jp/uploads/upload_2fa1180866e85cac7268a40c25786ce5.png)

3. Install Vercel で自分のアカウントを選択し、==Only select repositories==から今回のリポジトリを連携(Install)してください。
   ![](https://md.trap.jp/uploads/upload_ea8c1ea5466ccaa56e072780f2b171ed.png)

4. 連携できたら==PERSONAL ACCOUNT==を選択し、各種設定画面に進みます。
   ![](https://md.trap.jp/uploads/upload_5bd36633e407fa7ac8cf496d4ededfc3.png)
   ![](https://md.trap.jp/uploads/upload_b60e9f0a523d4c7afb91701a0124c3e0.png)

5. 画像の通り設定を進めて、==Deploy==します。
   ![](https://md.trap.jp/uploads/upload_16e092e8dfb72310897aeff26223e85e.png)
   ![](https://md.trap.jp/uploads/upload_6b8fecd35ed5e6b5575ebcebf633005a.png)

6. ログが流れる画面に遷移するので処理が終わるのを待ちます。ビルドが成功すると以下のような画面に遷移します。
   ![](https://md.trap.jp/uploads/upload_fdea23fe0fa0f94656d6f12490629773.png)

見本

- サイト：https://trap.jp
- ソースコード：https://github.com/traPtitech
  //todo: 作ってリンク貼る
