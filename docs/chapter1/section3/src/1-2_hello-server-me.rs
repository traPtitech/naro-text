use axum::{routing::get, Router};

#[tokio::main]
async fn main() {
    // 「/hello」「/kenken」の2つのエンドポイントを持つアプリケーションを作成
    let app = Router::new()
        .route("/hello", get(hello_handler))
        .route("/kenken", get(me_handler));

    // ポート8080でリスナーを作成する
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    // サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// 文字列「Hello, World.」をクライアントに返す
async fn hello_handler() -> String {
    String::from("Hello, World.")
}

// 自己紹介をクライアントに返す
async fn me_handler() -> String {
    String::from(
        "始めまして、@kenkenです。\nきらら作品(特に恋する小惑星、スロウスタート)が好きです。",
    )
}
