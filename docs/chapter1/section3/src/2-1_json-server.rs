use axum::{routing::get, Json, Router};

#[tokio::main]
async fn main() {
	// 「/json」というパスのエンドポイントを定義
    let app = Router::new()
        .route("/json", get(json_handler));

	// Webサーバーをポート番号8080にバインドする
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

	// サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// JSONで返すための構造体を定義
// 構造体を JSON に変換するためにserde::Serializeを導出する
#[derive(serde::Serialize)]
struct JsonData {
    number: i32,
    string: String,
    bool: bool,
}

async fn json_handler() -> Json<JsonData> {
	// レスポンスとして返す値を構造体として定義
    let res = JsonData {
        number: 10,
        string: String::from("hoge"),
        bool: false,
    };

	// 構造体をJSONに変換してクライアントに返す
    Json(res)
}