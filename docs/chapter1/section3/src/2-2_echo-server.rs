use core::fmt;

use axum::{
    extract::rejection::JsonRejection,
    http::StatusCode,
    routing::{get, post},
    Json, Router,
};

#[tokio::main]
async fn main() {
    // post() メソッドを使ってPOSTリクエストを処理する
    let app = Router::new()
        .route("/post", post(post_handler));

    // ポート8080でリスナーを作成する
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    // サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// JSONを受け取るための構造体を定義
// 構造体を JSON に変換するためにserde::Serializeを導出する
#[derive(serde::Deserialize, serde::Serialize)]
struct JsonData {
    number: i32,
    string: String,
    bool: bool,
}

async fn post_handler(
    query: Result<Json<JsonData>, JsonRejection>,
) -> Result<Json<JsonData>, (StatusCode, JsonRejection)> {
    match query {
        // 正常なときリクエストデータをそのまま返す
        Ok(data) => Ok(data),
        // 正常でないときステータスコード 400 Bad Requestを返す
        Err(rejection) => Err((StatusCode::BAD_REQUEST, rejection)),
    }
}
