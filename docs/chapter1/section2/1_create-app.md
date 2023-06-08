# アプリを作ってみよう

## 商品リストを作ってみる

みなさんにはこの節の最後に Todo リストを作ってもらうのですが、商品リストをテーマに、Todo リストに必要な Vue の機能をピックアップしていきます。

こんな感じのを作っていきます。
![](images/1/preview.gif)

### 必要な要素を考える

上の gif のようなアプリを実現するためには何が必要か考えてみましょう。

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
![](images/1/directory.png)

#### src/components/ItemList.vue

中身はコンポーネントに最低限必要な部分だけ書きます。

<<< @/chapter1/section2/src/1/ItemListInit.vue{vue:line-numbers}

#### HelloWorld.vue

<<< @/chapter1/section2/src/1/HelloWorld.vue{vue:line-numbers}

表示されました。
こうすることで、後は`ItemList.vue`の中身を書き変えればよくなります。

![](images/1/itemlist-setup.png)

### 商品のリストデータを保存する

商品リストのデータを保存するのに適当な変数の型は何でしょうか？  
商品「リスト」なので配列がよさそうです。  
というわけで、配列を使ってデータを保持することにします。  
今は商品の追加ができないので、とりあえずダミーデータを入れておきます。

参考: [Array | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)  
参考：[JavaScript オブジェクトの基本 - ウェブ開発を学ぶ | MDN](https://developer.mozilla.org/ja/docs/Learn/JavaScript/Objects/Basics)

<<< @/chapter1/section2/src/1/ItemListItems.vue{vue:line-numbers}

4~7 行目は TypeScript の記法で、`Item`という型を`interface`を用いて定義しています。  
そして ref のジェネリクスに`Item[]`を渡すことで、`items`変数を`Item`型の配列の`ref`として扱えるようにしています。

参考：[ジェネリクス (generics) | TypeScript 入門『サバイバル TypeScript』](https://typescriptbook.jp/reference/generics)  
参考：[インターフェース (interface) | TypeScript 入門『サバイバル TypeScript』](https://typescriptbook.jp/reference/object-oriented/interface)

### 商品のリストデータを表示する

先ほど定義したリストの情報を表示していきます。  
Vue ではリストデータを`template`タグ内で for 文のように書く `v-for` という構文があります。  
`v-for` を使うときには`:key`を設定しなければいけません(理由(やや難): [優先度 A: 必須 | Vue](https://ja.vuejs.org/style-guide/rules-essential.html#use-keyed-v-for))。

参考: [リストレンダリング | Vue](https://ja.vuejs.org/guide/essentials/list.html#v-for)

これを使ってデータを表示してみます。

<<< @/chapter1/section2/src/1/ItemListList.vue

表示できました。

![](images/1/itemlist.png)

### 商品を追加する

Vue では入力欄に入力された文字列とコンポーネントの変数を結びつけることができます。  
参考： [フォーム入力バインディング | Vue](https://ja.vuejs.org/guide/essentials/forms.html)

これを使って商品を追加できるようにしてみます。

<<< @/chapter1/section2/src/1/ItemListAdd.vue{vue:line-numbers}

参考: [アロー関数式 | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Functions/Arrow_functions)

できました！

![](images/1/add-item.gif)

#### 練習問題 1：商品リストに機能を追加

このままだとボタンを連打して商品の追加ができてしまいます。

- ボタンを押したら入力欄を空にする機能
- 入力欄が空だったらボタンを押しても追加されないようにする機能

を追加してみましょう。

### 商品の値段が 500 円以上だったら赤くする

Vue では、ある特定の条件が満たされたときに class を追加するという機構を持たせることできます。  
これを使って、条件が満たされたときだけ CSS を当てるといったことができます。

参考: [CSS の基本 | MDN](https://developer.mozilla.org/ja/docs/Learn/Getting_started_with_the_web/CSS_basics)  
参考: [クラスとスタイルのバインディング | Vue](https://ja.vuejs.org/guide/essentials/class-and-style.html#binding-html-classes)

<<< @/chapter1/section2/src/1/ItemListRed.vue

![](images/1/red.png)

### 商品の値段が 10000 円以上だったら「高額商品」と表示する

Vue では、ある特定の条件を満たした場合のみ、対象コンポーネントを表示するという機能を`v-if`という構文を使って実現できます。

参考: [条件付きレンダリング | Vue](https://ja.vuejs.org/guide/essentials/conditional.html)

これを使って商品の値段が 10000 円以上だったら「高額商品」と表示するという機能を実現してみましょう。

<<< @/chapter1/section2/src/1/ItemListExpensive.vue

![](images/1/expensive.png)

これで商品リストが完成しました！

今回の商品リストの全体像は以下のブランチに入っているので、参考にしてみてください。  
[traPtitech/naro-template-frontend at example/itemlist](https://github.com/traPtitech/naro-template-frontend/tree/example/itemlist)

## Todo リストを作る

ここまで紹介してきた機能を使うことで Todo リストが作れるはずです。
頑張りましょう！

#### 練習問題 2：Todo リストを作る

Todo リストを作りましょう。

必要な機能は以下の通りです。

- タスクは未完または完了済みの状態を持つ。
- タスクはタスク名を持つ。
- 未完タスクのリストと完了済みタスクのリストが表示される。
- タスクを完了させることができる。
- タスクの追加ができる。

以上の機能が実現されていれば後は自由です。
スタイルが気になる人は CSS なども書きましょう。

一応作成例は以下のブランチに作ってみましたが、できるだけ自力で頑張ってみてください。分からないことなどあれば遠慮なく TA に質問してください。  
[traPtitech/naro-template-frontend at example/todolist](https://github.com/traPtitech/naro-template-frontend/tree/example/todolist)