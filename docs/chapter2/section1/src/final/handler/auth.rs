use axum::{
    extract::{Request, State},
    http::{header, StatusCode},
    middleware::Next,
    response::IntoResponse,
    Json,
};
use axum_extra::{headers::Cookie, TypedHeader};
use serde::{Deserialize, Serialize};

use crate::repository::Repository;

pub async fn auth_middleware(
    State(state): State<Repository>,
    TypedHeader(cookie): TypedHeader<Cookie>,
    mut req: Request,
    next: Next,
) -> Result<impl IntoResponse, StatusCode> {
    // セッションIDを取得する
    let session_id = cookie
        .get("session_id")
        .ok_or(StatusCode::UNAUTHORIZED)?
        .to_string();

    // セッションストアからユーザーIDを取得する
    let user_id = state
        .get_user_id_by_session_id(&session_id)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?
        .ok_or(StatusCode::UNAUTHORIZED)?;

    // リクエストにユーザーIDを追加する
    req.extensions_mut().insert(user_id);

    // 次のミドルウェアを呼び出す
    Ok(next.run(req).await)
}

#[derive(Deserialize)]
pub struct SignUp {
    pub username: String,
    pub password: String,
}

pub async fn sign_up(
    State(state): State<Repository>,
    Json(body): Json<SignUp>,
) -> Result<StatusCode, StatusCode> {
    // バリデーションする(PasswordかUsernameが空文字列の場合は400 BadRequestを返す)
    if body.username.is_empty() || body.password.is_empty() {
        return Err(StatusCode::BAD_REQUEST);
    }

    // 登録しようとしているユーザーが既にデータベース内に存在したら409 Conflictを返す
    if let Ok(true) = state.is_exist_username(body.username.clone()).await {
        return Err(StatusCode::CONFLICT);
    }

    // ユーザーを作成する
    let id = state
        .create_user(body.username.clone())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    // パスワードを保存する
    state
        .save_user_password(id as i32, body.password.clone())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(StatusCode::CREATED)
}

#[derive(Deserialize)]
pub struct Login {
    pub username: String,
    pub password: String,
}

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

    // パスワードが一致しているかを確かめる
    if !state
        .verify_user_password(id, body.password.clone())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?
    {
        return Err(StatusCode::UNAUTHORIZED);
    }

    // セッションストアに登録する
    let session_id = state
        .create_user_session(id.to_string())
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    // クッキーをセットする
    let mut headers = header::HeaderMap::new();

    headers.insert(
        header::SET_COOKIE,
        format!("session_id={}; HttpOnly; SameSite=Strict", session_id)
            .parse()
            .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?,
    );

    Ok((StatusCode::OK, headers))
}

pub async fn logout(
    State(state): State<Repository>,
    TypedHeader(cookie): TypedHeader<Cookie>,
) -> Result<impl IntoResponse, StatusCode> {
    // セッションIDを取得する
    let session_id = cookie
        .get("session_id")
        .ok_or(StatusCode::UNAUTHORIZED)?
        .to_string();

    // セッションストアから削除する
    state
        .delete_user_session(session_id)
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    // クッキーを削除する
    let mut headers = header::HeaderMap::new();
    headers.insert(
        header::SET_COOKIE,
        "session_id=; HttpOnly; SameSite=Strict; Max-Age=0"
            .parse()
            .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?,
    );

    Ok((StatusCode::OK, headers))
}

#[derive(Serialize)]
pub struct Me {
    pub username: String,
}

pub async fn me(State(state): State<Repository>, req: Request) -> Result<Json<Me>, StatusCode> {
    // リクエストからユーザーIDを取得する
    let user_id = req
        .extensions()
        .get::<String>()
        .ok_or(StatusCode::UNAUTHORIZED)?
        .to_string();

    // データベースからユーザー名を取得する
    let username = state
        .get_user_name_by_id(
            user_id
                .parse()
                .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?,
        )
        .await
        .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;

    Ok(Json(Me { username }))
}
