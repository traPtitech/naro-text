# Vue.js で TodoList(公開編)

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
//todo: 画像貼る。

## 公開する

それでは公開しましょう！

早速 push して公開したいところですが、`.gitignore`をみてみると`dist`ディレクトリが ignore されていることがわかります。  
//todo: 画像貼る。

これは GitHub にはソースコードだけをアップロードし、そのソースコードから再現できるものは極力アップロードしない(Git のパフォーマンスに影響するため)という考えから来ているものです。

`node_modules`も同じような理由で ignore されていることがわかります。

ビルド済みの成果物を GitHub Pages などを用いて公開してもいいのですが、上記の考えに従い、今回は違ったサービスで公開します。

### Vercel で公開する

[Vercel: Develop. Preview. Ship. For the best frontend teams](https://vercel.com/)

Vercel を使うと、ビルドが必要なサイトも簡単に公開できるので便利です。

// todo:アカウント作るのめんどくさいので後回し。

1. Signup から以下の画面に進み、GitHub アカウントと連携してください。
   ![](https://md.trap.jp/uploads/upload_639f66f4a91154672e52ea41f770bc54.png)

2. [新規作成画面](https://vercel.com/new)から==Add GitHub Account==を選択してください。
   ![](https://md.trap.jp/uploads/upload_2ca44b1f8fe1577bde24865b48b80f3c.png)
   ![](https://md.trap.jp/uploads/upload_2fa1180866e85cac7268a40c25786ce5.png)

3. Install Vercel で自分のアカウントを選択し、==Only select repositories==から今回のリポジトリを連携(Install)してください。
   ![](https://md.trap.jp/uploads/upload_ea8c1ea5466ccaa56e072780f2b171ed.png)

4. 連携できたら==PERSONAL ACCOUNT==を選択し、各設定画面に進みます。
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

//todo: 作ってリンク貼る。
