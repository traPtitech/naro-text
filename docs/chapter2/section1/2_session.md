# セッション管理機構の実装

## セッションストアを設定する
`repository.rs`に以下を追加しましょう。
```rs
use async_sqlx_session::MySqlSessionStore; // [!code ++]
use sqlx::mysql::MySqlConnectOptions;
use sqlx::mysql::MySqlPool;
use std::env;

pub mod country;
pub mod users;

#[derive(Clone)]
pub struct Repository {
    pool: MySqlPool,
    session_store: MySqlSessionStore, // [!code ++]
}

impl Repository {
    pub async fn connect() -> anyhow::Result<Self> {
        let options = get_options()?;
        let pool = sqlx::MySqlPool::connect_with(options).await?;

        let session_store = // [!code ++]
            MySqlSessionStore::from_client(pool.clone()).with_table_name("user_sessions"); // [!code ++]

        Ok(Self {
            pool,
            session_store, // [!code ++]
        })
    }

    pub async fn migrate(&self) -> anyhow::Result<()> {
        sqlx::migrate!("./migrations").run(&self.pool).await?;
        Ok(())
    }
}
...(省略)
```

これらはセッションストアの設定です。
セッションの情報を記憶するための場所をデータベース上に設定して、`session_store` からアクセスできるようにしています。

## `login` ハンドラの実装
続いて、`login` ハンドラを `handler/auth.rs` に実装していきましょう。

```rs
pub async fn login( // [!code ++]
    State(state): State<Repository>, // [!code ++]
    Json(body): Json<Login>, // [!code ++]
) -> Result<impl IntoResponse, StatusCode> { // [!code ++]
} // [!code ++]
```

`login` ハンドラの外に以下の構造体を追加します。
```rs
#[derive(Deserialize)] // [!code ++]
pub struct Login { // [!code ++]
    pub username: String, // [!code ++]
    pub password: String, // [!code ++]
} // [!code ++]
```

`login` ハンドラの中身を実装する前に、必要になるデータベース操作のメソッドを追加します。ここで必要になるのは以下の 2 つです。

- `username` から `id` を取得するメソッド
- `id` と `password` の組が登録されているものと一致するかを確認するメソッド

この 2 つを `repository/users.rs` に追加します。
```rs
use super::Repository;

impl Repository {
    pub async fn is_exist_username(&self, username: String) -> sqlx::Result<bool> {
        ...(省略)
    }

    pub async fn create_user(&self, username: String) -> sqlx::Result<u64> {
        ...(省略)
    }

    pub async fn get_user_id_by_name(&self, username: String) -> sqlx::Result<u64> { // [!code ++]
        let result = sqlx::query_scalar("SELECT id FROM users WHERE username = ?") // [!code ++]
            .bind(&username) // [!code ++]
            .fetch_one(&self.pool) // [!code ++]
            .await?; // [!code ++]
        Ok(result) // [!code ++]
    } // [!code ++]

    pub async fn save_user_password(&self, id: i32, password: String) -> anyhow::Result<()> {
        ...(省略)
    }

    pub async fn verify_user_password(&self, id: u64, password: String) -> anyhow::Result<bool> { // [!code ++]
        let hash = // [!code ++]
            sqlx::query_scalar::<_, String>("SELECT hashed_pass FROM user_passwords WHERE id = ?") // [!code ++]
                .bind(id) // [!code ++]
                .fetch_one(&self.pool) // [!code ++]
                .await?; // [!code ++]

        Ok(bcrypt::verify(password, &hash)?) // [!code ++]
    } // [!code ++]
}
```

データベースに保存されているパスワードはハッシュ化されています。

ハッシュ化は不可逆な処理なので、ハッシュ化されたものから原文を調べることはできません。確認する際はもらったパスワードをハッシュ化することで行います。
`bcrypt::verify` によってパスワードの検証ができます。

`handler/auth.rs` に戻り、`login` ハンドラを実装していきます。

```rs
pub async fn login( // [!code ++]
    State(state): State<Repository>, // [!code ++]
    Json(body): Json<Login>, // [!code ++]
) -> Result<impl IntoResponse, StatusCode> { // [!code ++]
    // バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す) // [!code ++]
    if body.username.is_empty() || body.password.is_empty() { // [!code ++]
        return Err(StatusCode::BAD_REQUEST); // [!code ++]
    } // [!code ++]

    // データベースからユーザーを取得する // [!code ++]
    let id = state // [!code ++]
        .get_user_id_by_name(body.username.clone()) // [!code ++]
        .await // [!code ++]
        .map_err(|e| match e { // [!code ++]
            sqlx::Error::RowNotFound => StatusCode::UNAUTHORIZED, // [!code ++]
            _ => StatusCode::INTERNAL_SERVER_ERROR, // [!code ++]
        })?; // [!code ++]
} // [!code ++]
```

ユーザーが存在しなかった場合は `sqlx::Error::RowNotFound` というエラーが返ってきます。
もしそのエラーなら 401 (Unauthorized)、そうでなければ 500 (Internal Server Error) です。
もし 404 (Not Found) とすると、「このユーザーはパスワードが違うのではなく存在しないんだ」という事がわかってしまい（このユーザーは存在していてパスワードは違う事も分かります）、セキュリティ上のリスクに繋がります。

```rs
pub async fn login(
    State(state): State<Repository>,
    Json(body): Json<Login>,
) -> Result<impl IntoResponse, StatusCode> {
    // バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す)
    if body.username.is_empty() || body.password.is_empty() {
        return Err(StatusCode::BAD_REQUEST);
    }

    // データベースからユーザーを取得する
    let id = state
        .get_user_id_by_name(body.username.clone())
        .await
        .map_err(|e| match e {
            sqlx::Error::RowNotFound => StatusCode::UNAUTHORIZED,
            _ => StatusCode::INTERNAL_SERVER_ERROR,
        })?;

    // パスワードが一致しているかを確かめる // [!code ++]
    if !state // [!code ++]
        .verify_user_password(id, body.password.clone()) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)? // [!code ++]
    { // [!code ++]
        return Err(StatusCode::UNAUTHORIZED); // [!code ++]
    } // [!code ++]
}
```

データベースでエラーが起きた場合や、検証の操作に失敗した場合は 500 (Internal Server Error), パスワードが間違っていた場合 401 (Unauthorized) を返却しています。

```rs
pub async fn login( 
    State(state): State<Repository>, 
    Json(body): Json<Login>,
) -> Result<impl IntoResponse, StatusCode> {
    ...(省略)
    
    // パスワードが一致しているかを確かめる
    if !state
        .verify_user_password(id, body.password.clone())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?
    {
        return Err(StatusCode::UNAUTHORIZED);
    } 

    // セッションストアに登録する // [!code ++]
    let session_id = state // [!code ++]
        .create_user_session(id.to_string()) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?; // [!code ++]
    
}
```

`id` をセッションストアに登録して、セッション id を取得します。

ここで用いる、セッションストアに登録するメソッド `create_user_session` を実装していきます。

ファイル `repository/users_session.rs` を作成し、以下を記述してください。

```rs
use anyhow::Context; // [!code ++]
use async_session::{Session, SessionStore}; // [!code ++]

use super::Repository; // [!code ++]

impl Repository { // [!code ++]
    pub async fn create_user_session(&self, user_id: String) -> anyhow::Result<String> { // [!code ++]
        let mut session = Session::new(); // [!code ++]

        session // [!code ++]
            .insert("user_id", user_id) // [!code ++]
            .with_context(|| "Failed to insert user_id")?; // [!code ++]

        let session_id = self // [!code ++]
            .session_store // [!code ++]
            .store_session(session) // [!code ++]
            .await // [!code ++]
            .with_context(|| "Failed to store session")? // [!code ++]
            .with_context(|| "Failed to create session")?; // [!code ++]

        Ok(session_id) // [!code ++]
    } // [!code ++]
} // [!code ++]
```

セッションに `user_id` を登録し、セッションストアに保存します。
セッション id を返却します。

`handler/auth.rs` に戻り、ヘッダーにセッション id を設定する処理を追加します。

```rs
pub async fn login(
    State(state): State<Repository>,
    Json(body): Json<Login>,
) -> Result<impl IntoResponse, StatusCode> {
    ...(省略)

    // セッションストアに登録する
    let session_id = state
        .create_user_session(id.to_string())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    // クッキーをセットする // [!code ++]
    let mut headers = header::HeaderMap::new(); // [!code ++]

    headers.insert( // [!code ++]
        header::SET_COOKIE, // [!code ++]
        format!("session_id={}; HttpOnly; SameSite=Strict", session_id) // [!code ++]
            .parse() // [!code ++]
            .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?, // [!code ++]
    ); // [!code ++]

    Ok((StatusCode::OK, headers)) // [!code ++]
}
```

ここまで書いたら、 `login` ハンドラを使えるようにしましょう。
`handler.rs` に以下を追加してください。

```rs
pub fn make_router(app_state: Repository) -> Router {
    let city_router = Router::new()
        .route("/cities/:city_name", get(country::get_city_handler))
        .route("/cities", post(country::post_city_handler));

    let auth_router = Router::new()
        .route("/signup", post(auth::sign_up))
        .route("/login", post(auth::login)); // [!code ++]

    Router::new()
        .nest("/", city_router)
        .nest("/", auth_router)
        .with_state(app_state)
}
```

:::details ここまでの全体像
::: code-group
<<<@/chapter2/section1/src/2_session/auth.rs{rs:line-numbers}[auth.rs]
<<<@/chapter2/section1/src/2_session/users.rs{rs:line-numbers}[users.rs]
<<<@/chapter2/section1/src/2_session/users_session.rs{rs:line-numbers}[users_session.rs]
<<<@/chapter2/section1/src/2_session/repository.rs{rs:line-numbers}[repository.rs]
:::

## Middleware の実装

続いて、`auth_middleware` を実装します。
まず、これは Handler ではなく Middleware と呼ばれます。

送られてくるリクエストは、Middleware を経由して、 Handler に流れていきます。

Middleware から次の Middleware/Handler を呼び出す際は `next.run(req)` と記述します。 

以下を`handler/auth.rs`に追加してください。

```rs
pub async fn auth_middleware( // [!code ++]
    State(state): State<Repository>, // [!code ++]
    TypedHeader(cookie): TypedHeader<Cookie>, // [!code ++]
    mut req: Request, // [!code ++]
    next: Next, // [!code ++]
) -> Result<impl IntoResponse, StatusCode> { // [!code ++]

    // セッションIDを取得する // [!code ++]
    let session_id = cookie // [!code ++]
        .get("session_id") // [!code ++]
        .ok_or(StatusCode::UNAUTHORIZED)? // [!code ++]
        .to_string(); // [!code ++]

    // セッションストアからユーザーIDを取得する // [!code ++]
    let user_id = state // [!code ++]
        .get_user_id_by_session_id(&session_id) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)? // [!code ++]
        .ok_or(StatusCode::UNAUTHORIZED)?; // [!code ++]

    // リクエストにユーザーIDを追加する // [!code ++]
    req.extensions_mut().insert(user_id); // [!code ++]

    // 次のミドルウェアを呼び出す // [!code ++]
    Ok(next.run(req).await) // [!code ++]
} // [!code ++]
```

この Middleware はリクエストを送ったユーザーがログインしているのかをチェックし、
ログインしているならリクエスト(`req`) に `user_id` を追加します。

Cookie からセッション id を取得し、セッションストアからユーザー id を取得します。
ここで、セッション id がなかった場合や、不正なセッション id だった場合は 401 (Unauthorized) を返却します。
正しくログインされていれば、次の Middleware/Handler を呼び出します。

ここで使用した、 `get_user_id_by_session_id` メソッドを `repository/users_session.rs` に追加します。

```rs
pub async fn get_user_id_by_session_id( // [!code ++]
    &self, // [!code ++]
    session_id: &String, // [!code ++]
) -> anyhow::Result<Option<String>> { // [!code ++]
    let session = self // [!code ++]
        .session_store // [!code ++]
        .load_session(session_id.clone()) // [!code ++]
        .await // [!code ++]
        .with_context(|| "Failed to load session")?; // [!code ++]

    Ok(session.and_then(|s| s.get::<String>("user_id"))) // [!code ++]
} // [!code ++]
```

最後に、Middleware を設定しましょう。
ログインが必要なエンドポイントを `with_auth_router` でまとめ、Middleware を適用します。

`handler.rs` に以下を追加してください。

```rs
use axum::{
    middleware::from_fn_with_state, // [!code ++]
    routing::{get, post},
    Router,
};

use crate::repository::Repository;

mod auth;
mod country;

pub fn make_router(app_state: Repository) -> Router {
    let city_router = Router::new() // [!code --]
    let with_auth_router = Router::new() // [!code ++]
        .route("/cities/:city_name", get(country::get_city_handler))
        .route("/cities", post(country::post_city_handler));
        .route_layer(from_fn_with_state(app_state.clone(), auth::auth_middleware)); // [!code ++]

    let auth_router = Router::new()
        .route("/signup", post(auth::sign_up))
        .route("/login", post(auth::login)); 

    Router::new()
        .nest("/", city_router) // [!code --]
        .nest("/", with_auth_router) // [!code ++]
        .nest("/", auth_router)
        .nest("/", ping_router)
        .with_state(app_state)
}
```

これで、この章の目標である「ログインしないと利用できないようにする」が達成されました。

## logout ハンドラの実装

ログアウト機能をまだ実装していなかったので、 `logout` ハンドラを実装していきます。

まず、`handler/auth.rs` に以下を追加してください。

```rs
pub async fn logout( // [!code ++]
    State(state): State<Repository>, // [!code ++]
    TypedHeader(cookie): TypedHeader<Cookie>, // [!code ++]
) -> Result<impl IntoResponse, StatusCode> { // [!code ++]
    // セッションIDを取得する // [!code ++]
    let session_id = cookie // [!code ++]
        .get("session_id") // [!code ++]
        .ok_or(StatusCode::UNAUTHORIZED)? // [!code ++]
        .to_string(); // [!code ++]

    // セッションストアから削除する // [!code ++]
    state // [!code ++]
        .delete_user_session(session_id) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?; // [!code ++]

    // クッキーを削除する  // [!code ++]
    let mut headers = header::HeaderMap::new(); // [!code ++]
    headers.insert( // [!code ++]
        header::SET_COOKIE, // [!code ++]
        "session_id=; HttpOnly; SameSite=Strict; Max-Age=0" // [!code ++]
            .parse() // [!code ++]
            .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?, // [!code ++]
    ); // [!code ++]

    Ok((StatusCode::OK, headers)) // [!code ++]
} // [!code ++]
```

ログアウトするときは、ログインするときとは逆にセッションと Cookie を削除します。

ここで呼び出す `delete_user_session` メソッドを `repository/users_session.rs` に追加します。

```rs
pub async fn delete_user_session(&self, session_id: String) -> anyhow::Result<()> { // [!code ++]
    let session = self // [!code ++]
        .session_store // [!code ++]
        .load_session(session_id.clone()) // [!code ++]
        .await // [!code ++]
        .with_context(|| "Failed to load session")? // [!code ++]
        .with_context(|| "Failed to find session")?; // [!code ++]

    self.session_store // [!code ++]
        .destroy_session(session) // [!code ++]
        .await // [!code ++]
        .with_context(|| "Failed to destroy session")?; // [!code ++]

    Ok(()) // [!code ++]
} // [!code ++]
```

セッション ID からセッションを取得し、セッションストアから削除します。

最後に、`handler.rs` に `logout` ハンドラを追加します。

```rs
let auth_router = Router::new()
    .route("/signup", post(auth::sign_up))
    .route("/login", post(auth::login))
    .route("/logout", post(auth::logout));　// [!code ++]
```


## me ハンドラの実装

最後に、 `me` ハンドラを実装します。叩いたときに自分の情報が返ってくるエンドポイントです。

以下を `handler/auth.rs` に追加してください。

```rs
#[derive(Serialize)] // [!code ++]
pub struct Me { // [!code ++]
    pub username: String, // [!code ++]
}// [!code ++]
```

```rs
pub async fn me(State(state): State<Repository>, req: Request) -> Result<Json<Me>, StatusCode> { // [!code ++]
    // リクエストからユーザーIDを取得する // [!code ++]
    let user_id = req // [!code ++]
        .extensions() // [!code ++]
        .get::<String>() // [!code ++]
        .ok_or(StatusCode::UNAUTHORIZED)? // [!code ++]
        .to_string(); // [!code ++]

    // データベースからユーザー名を取得する // [!code ++]
    let username = state // [!code ++]
        .get_user_name_by_id( // [!code ++]
            user_id // [!code ++]
                .parse() // [!code ++]
                .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?, // [!code ++]
        ) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?; // [!code ++]

    Ok(Json(Me { username })) // [!code ++]
} // [!code ++]
```

リクエストからユーザー ID を取得し、データベースからユーザー名を取得します。

ここで呼び出す `get_user_name_by_id` メソッドを `repository/users.rs` に追加します。

```rs
impl Repository {
    ...(省略)

    pub async fn delete_user_session(&self, session_id: String) -> anyhow::Result<()> { // [!code ++]
        let session = self// [!code ++]
            .session_store// [!code ++]
            .load_session(session_id.clone()) // [!code ++]
            .await// [!code ++]
            .with_context(|| "Failed to load session")? // [!code ++]
            .with_context(|| "Failed to find session")?; // [!code ++]

        self.session_store // [!code ++]
            .destroy_session(session) // [!code ++]
            .await // [!code ++]
            .with_context(|| "Failed to destroy session")?; // [!code ++]

        Ok(()) // [!code ++]
    } // [!code ++]

    ...(省略)
}
```

最後に、`handler.rs` に `me` ハンドラを追加します。

```rs
let with_auth_router = Router::new()
        .route("/cities/:city_name", get(country::get_city_handler))
        .route("/cities", post(country::post_city_handler))
        .route("/me", get(auth::me)) // [!code ++]
        .route_layer(from_fn_with_state(app_state.clone(), auth::auth_middleware));
```
