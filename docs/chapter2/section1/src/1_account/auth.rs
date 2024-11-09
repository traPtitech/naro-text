use axum::{extract::State, http::StatusCode, Json};
use serde::Deserialize;

use crate::repository::Repository;

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
