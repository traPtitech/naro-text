use axum::{routing::get, Router};

#[tokio::main]
async fn main() {
    // 「/ping」というエンドポイントを設定する
    let app = Router::new().route("/ping", get(handler));

    // ポート8080でリスナーを作成する
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    // サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// 文字列「pong」をクライアントに返す
async fn handler() -> String {
    String::from("pong")
}
