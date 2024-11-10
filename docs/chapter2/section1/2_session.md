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
) -> Result<StatusCode, StatusCode> { // [!code ++]
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
) -> Result<StatusCode, StatusCode> { // [!code ++]
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
) -> Result<StatusCode, StatusCode> {
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
) -> Result<StatusCode, StatusCode> {
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
    state // [!code ++]
        .create_user_session(id.to_string()) // [!code ++]
        .await // [!code ++]
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?; // [!code ++]

    Ok(StatusCode::OK) // [!code ++]
}
```

`id` をセッションストアに登録します。

ここで用いる、セッションストアに登録するメソッド `create_user_session` を実装していきます。

ファイル `repository/users_session.rs` を作成し、以下を記述してください。

```rs
use anyhow::Context; // [!code ++]
use async_session::{Session, SessionStore}; // [!code ++]

use super::Repository; // [!code ++]

impl Repository { // [!code ++]
    pub async fn create_user_session(&self, user_id: String) -> anyhow::Result<()> { // [!code ++]
        let mut session = Session::new(); // [!code ++]

        session // [!code ++]
            .insert("user_id", user_id) // [!code ++]
            .with_context(|| "Failed to insert user_id")?; // [!code ++]

        let result = self // [!code ++]
            .session_store // [!code ++]
            .store_session(session) // [!code ++]
            .await // [!code ++]
            .with_context(|| "Failed to store session") // [!code ++]
            .with_context(|| "Failed to store session")?; // [!code ++]

        match result { // [!code ++]
            Some(_) => Ok(()), // [!code ++]
            None => Err(anyhow::anyhow!("Failed to store session")), // [!code ++]
        } // [!code ++]
    } // [!code ++]
} // [!code ++]
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

## userAuthMiddleware の実装

続いて、`userAuthMiddleware` を実装します。
まず、これは Handler ではなく Middleware と呼ばれます。

送られてくるリクエストは、Middleware を経由して、 Handler に流れていきます。

Middleware から次の Middleware/Handler を呼び出す際は `next(c)` と記述します。 Middleware の実装は難しいので、なんとなく理解できれば十分です。

以下を`handler.go`に追加しましょう。
```go
func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc { // [!code ++]
	return func(c echo.Context) error { // [!code ++]
		sess, err := session.Get("sessions", c) // [!code ++]
		if err != nil { // [!code ++]
			log.Println(err) // [!code ++]
			return c.String(http.StatusInternalServerError, "something wrong in getting session") // [!code ++]
		} // [!code ++]
		if sess.Values["userName"] == nil { // [!code ++]
			return c.String(http.StatusUnauthorized, "please login") // [!code ++]
		} // [!code ++]
		c.Set("userName", sess.Values["userName"].(string)) // [!code ++]
		return next(c) // [!code ++]
	} // [!code ++]
} // [!code ++]
```

関数が関数を呼び出していて混乱しそうですが、 2 行目から 13 行目が本質で、外側はおまじないと考えて良いです。

この Middleware はリクエストを送ったユーザーがログインしているのかをチェックし、
ログインしているなら Context (`c`) にそのユーザーの UserName を設定します。

セッションを取得し、ログイン時に設定した `userName` の値を確認しに行きます。

ここで名前が入っていればリクエストの送信者はログイン済みで、そうでなければログインをしていないことが分かります。

これを利用して、ログインしていない場合には処理をここで止めて 401 (Unauthorized) を返却し、していれば次の処理 (`next(c)`)
に進みます。

最後に、Middleware を設定しましょう。
グループ機能を利用して、 `withAuth` に設定されてるエンドポイントは `userAuthMiddleware` を処理してから処理する、という設定をします。

```go
func main() {
	(省略)
	e.POST("/login", h.LoginHandler)

	e.GET("/cities/:cityName", h.GetCityInfoHandler) // [!code --]
	e.POST("/cities", h.PostCityHandler) // [!code --]
	withAuth := e.Group("") // [!code ++]
	withAuth.Use(handler.UserAuthMiddleware) // [!code ++]
	withAuth.GET("/cities/:cityName", h.GetCityInfoHandler) // [!code ++]
	withAuth.POST("/cities", h.PostCityHandler) // [!code ++]

	err = e.Start(":8080")
	(省略)
}
```

これで、この章の目標である「ログインしないと利用できないようにする」が達成されました。

## GetMeHandler の実装

最後に、 `GetMeHandler` を実装します。叩いたときに自分の情報が返ってくるエンドポイントです。

以下を `handler.go` に追加しましょう。
```go
type Me struct { // [!code ++]
	Username string `json:"username,omitempty"  db:"username"` // [!code ++]
} // [!code ++]
```
```go
func GetMeHandler(c echo.Context) error { // [!code ++]
	return c.JSON(http.StatusOK, Me{ // [!code ++]
		Username: c.Get("userName").(string), // [!code ++]
	}) // [!code ++]
} // [!code ++]
```

アクセスしているユーザーの`userName`をセッションから取得して返しています。
`userAuthMiddleware` を実行したあとなので、`c.Get("userName").(string)` によって userName を取得できます。

`main.go`に`withAuth.GET("/me", handler.GetMeHandler)`を追加しましょう。
```go
func main() {
	(省略)
	withAuth := e.Group("")
	withAuth.Use(handler.UserAuthMiddleware)
	withAuth.GET("/me", handler.GetMeHandler) // [!code ++]
	withAuth.GET("/cities/:cityName", h.GetCityInfoHandler)
	withAuth.POST("/cities", h.PostCityHandler)

	err = e.Start(":8080")
	(省略)
}
```
