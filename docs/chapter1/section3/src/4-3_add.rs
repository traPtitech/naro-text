use axum::extract::rejection::JsonRejection;
use axum::{http::StatusCode, routing::post, Json, Router};

#[tokio::main]
async fn main() {
    // 「/ping」というエンドポイントを設定する
    let app = Router::new().route("/add", post(add_handler));

    // ポート8080でリスナーを作成する
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    // サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// クエリパラメータを受け取るための構造体を定義
#[derive(serde::Deserialize)]
struct AddQuery {
    left: f64,
    right: f64,
}

// レスポンスとして返す構造体を定義
#[derive(serde::Serialize)]
struct AddResponse {
    result: i64,
}

// エラーレスポンスとして返す構造体を定義
#[derive(serde::Serialize)]
struct AddError {
    error: String,
}

async fn add_handler(
    query: Result<Json<AddQuery>, JsonRejection>,
) -> Result<Json<AddResponse>, (StatusCode, Json<AddError>)> {
    match query {
        // クエリが正しく受け取れた場合、クライアントに結果を返す
        Ok(query) => Ok(Json(AddResponse {
            result: (query.left + query.right) as i64,
        })),
        // クエリが正しく受け取れなかった場合、エラーを返す
        Err(_) => Err((
            StatusCode::BAD_REQUEST,
            Json(AddError {
                error: String::from("Bad Request"),
            }),
        )),
    }
}