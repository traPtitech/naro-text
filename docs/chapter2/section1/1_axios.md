# axiosの導入

実際にサーバーにリクエストを送るために便利なライブラリとして、axios を導入します。
参考: [Axios Docs](https://axios-http.com/docs/intro)
参考: [axios/axios: Promise based HTTP client for the browser and node.js](https://github.com/axios/axios)


ターミナルでプロジェクトのルートディレクトリに移動し、以下のコマンドを実行します。
```
$ npm install axios
```

# axiosの使い方

## axiosを使うためのページを作成

axios を使うためのページを作成し、`App.vue`にリンクを追加します。

### src/pages/AxiosPage.vue

```htmlmixed=
<template>
  <div>
    <button @click="sendGet">GET</button>
    <button @click="sendPost">POST</button>
    <button @click="sendPut">PUT</button>
    <button @click="sendDelete">DELETE</button>
  </div>
</template>
```

### src/App.vue

```htmlmixed=
<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/">Home</router-link>
        |
        <router-link to="/axios">Axios</router-link>
      </header>

      <router-view />
    </div>
  </main>
</template>
==略==
```

### src/router.js

Axios コンポーネントをインポートし、ルーターに登録します。
```javascript=
import { createRouter, createWebHistory } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import NotFound from "./pages/NotFound.vue";
import AxiosPage from "./pages/AxiosPage.vue";

const routes = [
  { path: "/", name:"home", component: HomePage },
  { path: "/axios", name:"axios", component: AxiosPage },
  { path: "/:path(.*)", component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
```

追加されました
![](https://md.trap.jp/uploads/upload_1f14b208185e842cbd9efc0ef91a3e21.png)


## RequestBinで自分のエンドポイントを作成
以前の講習会でも使った RequestBin を使って API を叩く練習をしてみましょう。

[RequestBin - Reliably inspect and observe any HTTP traffic](https://requestbin.com/)

上のリンクを開きます。
Create a Request Bin を押します。
![](https://md.trapti.tech/uploads/upload_ce9add981796a2cb2b57085b088ad1dc.png)

次のページで表示されたこの URL があなたのエンドポイントです。
![](https://md.trapti.tech/uploads/upload_cd26e791860b6b484e40743591e3d628.png)

## AxiosPage.vueのスクリプトを書く
axios を利用して、リクエストを送るスクリプトを書きます。

:::warning
以下のコードでは自分のエンドポイントに置き換えて、コードを書いてください
:::
### src/pages/AxiosPage.vue

```javascript=
<script setup>
import axios from "axios";
const sendGet = () => axios.get("https://ent5o104rfsmh.x.pipedream.net");
const sendPost = () => axios.post("https://ent5o104rfsmh.x.pipedream.net", {
  traQID: "temma",
  grade: 19,
});
const sendPut = () => axios.put("https://ent5o104rfsmh.x.pipedream.net", {
  testData: { string: "string", number: 1 },
});
const sendDelete = () => axios.delete("https://ent5o104rfsmh.x.pipedream.net", {
  data: { po: 1 },
});

</script>
==略==
```

`axios.post`や`axios.put`では、第二引数としてオブジェクトを渡すと、JSON 形式でサーバーに投げてくれます。

delete は若干特殊で`{data: {}}` とすれば`data`の中身が投げられます。

## 試してみる
実際にボタンを押してリクエストが送られているか試してみましょう。

Chrome Devtool の network タブを開くことでリクエストの様子を見ることができます。

![](https://md.trap.jp/uploads/upload_cd8ea06ad1025c2699c419e3f01b5baf.gif)
:::warning
POST リクエストの前に OPTIONS というリクエストが飛んでいますが、それは Preflight request というものです。
[Preflight request (プリフライトリクエスト) | MDN](https://developer.mozilla.org/ja/docs/Glossary/Preflight_request)
:::

## RequestBinを見てみる
先程生成したエンドポイントを見てみると実際にリクエストが送られてきているのがわかります。
![](https://md.trap.jp/uploads/upload_fdb9900c401319e831d81e980ecf4624.png)


:::warning
同じように、自分の用意したサーバーに対して、リクエストを送ることができます。
できるのですが、異なるオリジンからリクエストを送ろうとするとプロキシの項で言ったようなブラウザのセキュリティ機構に引っかかるので注意が必要です。

参考: [同一オリジンポリシー | MDN](https://developer.mozilla.org/ja/docs/Web/Security/Same-origin_policy)
:::

# ログインページの作成
:::	warning
リモートでのサーバー起動前に以下のコマンドを実行することを忘れないこと。(忘れると DB にアクセスできない)
```
$ source env.sh
```

:::

:::success
ログインページを作成してみましょう！
上と同じようにページを分けて進めていきましょう。

- ユーザー名とパスワードが入力できる
	- input タグを使いましょう
	- v-model とかをうまく使いましょう
		- [v-model | Vue.js](https://v3.vuejs.org/guide/migration/v-model.html)
- ログインボタンを押すと`/api/login`に POST する
	- POST の JSON はサーバーのコードに合わせて書きましょう

<details>
<summary>解答</summary>

### src/pages/LoginPage.vue

新規作成
```htmlmixed=
<script setup>
import axios from "axios";
import { ref } from 'vue';

const username = ref('');
const password = ref('');
const login = () => axios.post("/api/login", { username: username.value, password: password.value });

</script>

<template>
  <div class="login">
    <h1>This is an login page</h1>
    <div>
      <input type="text" v-model="username" />
      <input type="password" v-model="password" />
    </div>
    <div>
      <button @click="login">login</button>
    </div>
  </div>
</template>
```

### src/App.vue

template 部分のみ
```htmlmixed=
<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/">Home</router-link>
        |
        <router-link to="/axios">Axios</router-link>
        |
        <router-link to="/login">Login</router-link>
      </header>

      <router-view />
    </div>
  </main>
</template>
```

### src/router.js

```javascript=
import { createRouter, createWebHistory } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import NotFound from "./pages/NotFound.vue";
import AxiosPage from "./pages/AxiosPage.vue";
import LoginPage from "./pages/LoginPage.vue";

const routes = [
  { path: "/", name:"home", component: HomePage },
  { path: "/axios", name:"axios", component: AxiosPage },
  { path: "/login", name:"login", component: LoginPage },
  { path: "/:path(.*)", component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

```

</details>
:::

# ログイン済みページの作成

昨日作った都市の情報を返す API を表示するページを作成します。

### src/pages/CityPage.vue

新規に作成するファイルです。
```htmlmixed=
<script setup>
import axios from "axios";
import { ref, onMounted } from "vue";
const props = defineProps({
  cityName: {
    type: String,
    required: true,
  },
});
const cityInfo = ref();
onMounted(async () => {
  const res = await axios.get("/api/cities/" + props.cityName);
  cityInfo.value = res.data;
});
</script>

<template>
  <div>
    <h1>
      {{ cityName }}
    </h1>
    <div v-if="cityInfo">{{ cityInfo }}</div>
    <div v-else>街が見つかりませんでした</div>
  </div>
</template>
```

### src/router.js

`CityPage.vue`を読み込み登録します。
echo と同じように、`PATH`に`:`始まりで書くと、PathParameter として値を取得できます。

参考: [Dynamic Route Marching | Vue Router](https://next.router.vuejs.org/guide/essentials/dynamic-matching.html)

```javascript=
import { createRouter, createWebHistory } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import NotFound from "./pages/NotFound.vue";
import AxiosPage from "./pages/AxiosPage.vue";
import LoginPage from "./pages/LoginPage.vue";
import CityPage from "./pages/CityPage.vue";

const routes = [
  { path: "/", name: "home", component: HomePage },
  { path: "/axios", name: "axios", component: AxiosPage },
  { path: "/login", name: "login", component: LoginPage },
  { path: "/city/:cityName", name: "city", component: CityPage, props: true },
  { path: "/:path(.*)", component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

```

### src/App.vue

リンクを追加します。
```htmlmixed=
<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/">Home</router-link>
        |
        <router-link to="/axios">Axios</router-link>
        |
        <router-link to="/city/Tokyo">Tokyo</router-link>
        |
        <router-link to="/login">Login</router-link>
      </header>

      <router-view />
    </div>
  </main>
</template>
```

## 確認

完成するとこんな感じ。

![](https://md.trap.jp/uploads/upload_6870d0b68ea440a6b466f4e1e15135d6.png)


:::info
HomePage.vue に任意の都市について表示できるような仕組みを作ってみましょう。
- input タグで都市名を指定
- 「表示する」のようなボタンを押すことで`/city/{その都市名}`というリンクに飛ばす

参考: [Programmatic Navigation | Vue Router](https://next.router.vuejs.org/guide/essentials/navigation.html)
:::


# ログインしていない場合に、ログインページに遷移させる

Chrome のシークレットウィンドウを起動し、先程の`/city/Tokyo`を開いてみます。

![](https://md.trap.jp/uploads/upload_9aa94d6e853f3efecf87d99696c44b31.png)


本来は上のスクリーンショットのように東京の情報が表示されてほしいですが、表示されません。
![](https://md.trap.jp/uploads/upload_b0f0d7786a00f839edb30b9d3f2ba65a.png)



Chrome Devtool から見てみるとログインしていないため、ダメだということがわかりました。
![](https://md.trap.jp/uploads/upload_46d499b9c5398001ea5e236c116d2145.png)



そこで、ログインしていない場合にはログインページにリダイレクトするように変更してみます。


## whoamiエンドポイントの作成

:::warning
サーバーサイドの変更をします。
:::

上のように何らかのエンドポイントを叩いた結果、403 が返ってきたらリダイレクトするようにしてもいいですが、今回は traQ やその他 traP の Vue での書き方に習って、`whoami`というエンドポイントを使ってログインされているかの確認をします。

このエンドポイントはログインしているユーザー自身の情報を取得するエンドポイントです。なぜこんなエンドポイントが必要かというと、Vue.js 自身は自分が何というユーザーでログインしているかをサーバーに問い合わせることなく知ることができないからです。
traQ でも一番始めに whoami エンドポイントを叩き自分の情報を取得しています。

### main.go
#### ルーティングの追加

```go=62
	withLogin := e.Group("")
	withLogin.Use(checkLogin)
	withLogin.GET("/cities/:cityName", getCityInfoHandler)
	withLogin.GET("/whoami", getWhoAmIHandler)
```

#### 関数の追加

```go=29
type Me struct {
	Username string `json:"username,omitempty"  db:"username"`
}
```

```go=177
func getWhoAmIHandler(c echo.Context) error {
	sess, _ := session.Get("sessions", c)
  
	return c.JSON(http.StatusOK, Me{
		Username: sess.Values["userName"].(string),
	})
}
```

## router.jsでログインの確認を行う

Vue Router の`beforeEach`という機能を使って、各 Routing の前に特定の関数を呼び出すことができます。
このようにログイン状態を確認する方法はパターンとして覚えてしまってもいいと思います。

`beforeEach`に関して詳しくは: [Navigation Guards | Vue Router](https://next.router.vuejs.org/guide/advanced/navigation-guards.html)

### src/router.js

```javascript=
import axios from 'axios'
```
```javascript=22
router.beforeEach(async (to) => {
  try {
    await axios.get("/api/whoami");
  } catch (_) {
    if (to.path === "/login") {
      return true;
    }
    return "/login";
  }
  return true;
});

export default router;
```

これでログインしていない場合には、`/login`へリダイレクトされるようになりました。
しかし、`/login`以外の全てのページへアクセスできません(シークレットウィンドウなどで開いて確認してみましょう)。

## 特定のページだけログイン不要にする

素朴な実装としては、`beforeEach`の中の条件分岐を増やして許可するという方法が思いつきますが、同じような処理を何度も書くのは面倒ですし、読みにくくなります。
ここでは、Vue Router の meta という機能を使って Route にメタ情報を付与し、それを用いてログイン不要かどうかを判断します。

参考: [Route Meta Field | Vue Router](https://next.router.vuejs.org/guide/advanced/meta.html)

### ルーティング設定にmetaを追加

### src/router.js
```javascript=9
const routes = [
  { path: "/", name: "home", component: HomePage, meta: { isPublic: true } },
  { path: "/axios", name: "axios", component: AxiosPage },
  {
    path: "/login",
    name: "login",
    component: LoginPage,
    meta: { isPublic: true },
  },
  { path: "/city/:cityName", name: "city", component: CityPage, props: true },
  { path: "/:path(.*)", component: NotFound, meta: { isPublic: true } },
];
```

### リダイレクト設定の変更

`if (to.path === '/login')`で分岐していたところを`if (to.meta.isPublic)`に置き換えます。

### src/router.js
```javascript=
router.beforeEach(async (to) => {
  try {
    await axios.get("/api/whoami");
  } catch (_) {
    if (to.meta.isPublic) {
      return true;
    }
    return "/login";
  }
  return true;
});
```

クライアントの見本：https://github.com/itt828/naro-client-2022-v2

:::warning
### これでサーバー・クライアント両方でAPIを利用する方法がわかりました。
これからは必要な API を考え、実装していくことになります。
:::


:::success
# 最重要課題

 国一覧を表示するページを作り、その国名をクリックすると、その国の都市一覧が表示され、その都市名をクリックすると都市の情報が表示されるようにしてみましょう。
:::

:::info 
## 発展課題
ログアウト機能を作りましょう。
### ヒント
* サーバープログラムに`/logout`を作る
* API を叩いた人のセッションをセッションストアから破棄する
* クライアントプログラムに`/logout`の API を叩くボタン作る
:::