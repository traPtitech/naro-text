# DevToolsについて
この章では，Webアプリ開発において欠かせないChrome DevToolsについて説明します。この章では基本的な内容に絞って説明していますので，DevToolsについてより詳細に知りたい人は[公式のドキュメント](https://developer.chrome.com/docs/devtools?hl=ja)を参照してください。
[[toc]]
## DevToolsって何？
DevToolsとは，Chromeに直接組み込まれたウェブデベロッパーツールのセットです。これを使うことにより，そのページで使われているCSSを把握したり，どのようなエラーが生じているかを確認したりと，Web開発に必要な非常に多くのことを実現することができます。
## DevToolsの開き方
DevToolsの開き方はいくつかありますが，ここではそのうちの2つを紹介します。

最初に，ChromeメニューからDevToolsを開く方法を紹介します。まず，DevToolsを開きたいページでマウスを右クリックします。すると，次のようなメニューが表示されると思います。

![menu](./images/menu.png)

このメニューの「検証」をクリックすると次のような画面が開かれると思います。これがDevToolsです。

![DevTools-main](./images/DevTools-main.png)

併せて，ショートカットを利用してDevToolsを開く方法を紹介します。WindowsまたはLinuxの場合は`Ctrl+Shift+C`で，Macの場合は`Cmd+Option+C`でDevToolsのElement(要素)タブを開くことができます。そのほかにもいくつかのショートカットがあり，詳細は[このページ](https://developer.chrome.com/docs/devtools/open?hl=ja)を参照してください。

## DevToolsの機能の紹介
ここからは，DevToolsの機能のうち基本的かつ重要な内容について取り上げて説明します。この章を読んでいる皆さんも，手元でDevToolsを開いて実際に触りながら学習を進めてください。

具体的な機能の紹介に入る前に，DevToolsの言語設定を変更する方法を紹介しておきます。上記の手順でDevToolsを開いたら，右上の歯車マークをクリックしてください。すると，Setting画面が開かれますので，この画面のLanguageタブからお好みの言語を選択してください。


### Elementについて
まず，Elements(要素)について説明します。
DevToolsを開いてElementsタブを選択すると，次のような画面が表示されると思います。

![DevTools-Elements](./images/DevTools-Elements.png)

このパネルには，DevToolsを開いたページ全体のDOMとCSSが表示されています。もし個別の要素について詳しく調査したければ，DevToolsの左上に表示されているInspect(検査)アイコンをクリックしましょう。この状態でページ上でカーソルを動かすと焦点が当たっている要素についての詳細な情報が表示されます。またInspectモードのまま各要素をクリックすると，パネルのDOMツリーのうちクリックした要素に対応する部分がハイライト表示されます。

![DevTools-Inspect](./images/DevTools-Inspect.png)

さらに下の写真で示すように，ElementのパネルやStyleタブ，LayoutタブなどからCSS等を直接書き換えることもできます。

![DevTools-Styletab](./images/DevTools-Styletab.png)

また，Inspectアイコンの隣にあるアイコンをクリックすると，PC上でスマホなどのモバイル機器での表示を確認することができます。例えば，該当のアイコンをクリックした時に表示されるDimensionsタブからPixel7を選択すると，下の写真のように実際にPixel7で同じページを開いた時と同様の表示がなされます。この機能により，PCを使いながら他のモバイル機器での表示を簡単に確認することができます。

![DevTools-Pixel7](./images/DevTools-Pixel7.png)

### Consoleについて
次に，Consoleについて説明します。
DevToolsを開いてConsoleタブを選択すると，次のような画面が表示されると思います。
<!-- 写真 -->
ここにはアプリケーションのすべてのログやエラーが表示されます。これを見ることにより，開発中に生じたエラーの詳細がわかり，エラーを解決するための手助けとなります。また，ConsoleではJavaScriptのコードを実行することもできます。例えば，次のようなコードをConsoleに直接入力してみてください。
```js
alert("Hello,Console!")
```
うまくいけば次のように表示されると思います。

![DevTools-Console-alert](./images/DevTools-Console2.png)

直接入力する他にコードファイルをドラック&ドロップすることでもコードを実行することができますので，覚えておくと良いかもしれません。

### Networkについて
DevToolsを開いてNetworkタブを選択すると，次のような画面が表示されると思います。

![DevTools-Network](./images/DevTools-Network.png)

ここでは，ページを開いたときのネットワークアクティビティーを知ることができます。jsファイルやcssファイル，画像やフォントなどそのページで読み込む必要があるすべてのファイルについて，読み込み時のステータスや読み込みにどれくらいかかったかを知ることができます。

例えば下の写真を見てみると，analytics.jsというファイルはStatusが200であることから正しく読み込まれていて，また読み込みには11msかかったことを知ることができます。

![DevTools-analytics](./images/DevTools-analytics.png)

もし何らかの理由で正しく読み込まれていないときは，Statusに404(Not Found)などが表示されているでしょう。もし詳細について知りたければ，各ファイルのName欄をクリックすれば良いです。各ファイルのName欄をクリックすると次のような画面が表示されると思います。少し込み入った話になるので省略しますが，例えばHeadersタブを見てHTTPヘッダーを調べることができたり，Timingタブからネットワークアクティビティの内訳を調べることができたりします。

![DevTools-Network-Detail](./images/DevTools-Network-Detail.png)

### Applicationについて
DevToolsを開いてApplicationタブを選択すると，次のような画面が表示されると思います。

![DevTools-Application](./images/DevTools-Application.png)

ここではマニフェスト，Service Worker，ストレージ，キャッシュデータなど，Webアプリのさまざまな要素を検査したりデバッグしたりすることができます。パネルの左側を見るとApplication,Storage,Background services，Framesという4つの大きな区分けがあり，その下に様々な項目があると思います。このすべての項目について説明することが難しいので，区分けごとに簡単な解説をするにとどめます。

Applicationの項目にはマニフェスト(manifest.jsonに含まれる名前やversionなどの情報)，Service Worker(ウェブアプリケーション，ブラウザー，そしてネットワークの間に介在するプロキシサーバーのように振る舞うもの)，ストレージなどアプリに関する全般的な情報が含まれます。

Storageの項目には使用されているさまざまなストレージ方法が含まれており，編集することもできます。

Background servicesの項目では，キャッシュに関するテストを実行したり支払いハンドラのイベントを記録したりと，バックグラウンドサービスの検査，テスト，デバッグをすることができます。

Frameの項目では，Webページを複数に分割し，個別に読み込んで確認することができます。

## 最後に
ここではDevToolsの基本的な内容について説明しました。ぜひこれから，Webアプリを制作していて何かうまくいかないことがあればDevToolsを見る，という習慣をつけていきましょう。
最初にも述べましたが，DevToolsについてより詳細に知りたい人は公式のドキュメントなどを読んでください。

