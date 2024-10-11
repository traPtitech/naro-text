use axum::{extract::Path, routing::get, Router};

#[tokio::main]
async fn main() {
    // GETリクエストの"/hello/:username"というパターンに対応するルートを設定し、
    // URLのパラメータ(:username)を使用してhello_handler関数を呼び出す
    let app = Router::new().route("/hello/:username", get(hello_handler));

    // ポート8080でリスナーを作成する
    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    // サーバーを起動する
    axum::serve(listener, app).await.unwrap();
}

// ":username"というパスパラメーターを取得する
async fn hello_handler(Path(username): Path<String>) -> String {
    format!("Hello, {}.\n", username)
}
