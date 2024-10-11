use axum::{
    extract::{Path, Query},
    routing::get,
    Router,
};

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

// クエリパラメータを受け取るための構造体を定義
#[derive(serde::Deserialize)]
pub struct HelloQueryParam {
    lang: Option<String>,
    page: Option<String>,
}

async fn hello_handler(
    Path(username): Path<String>,
    // クエリパラメータを受け取る
    Query(query): Query<HelloQueryParam>,
) -> String {
    format!(
        "Hello, {}.\nlang: {}\npage: {}",
        username,
        query.lang.unwrap_or("".to_string()),
        query.page.unwrap_or("".to_string())
    )
}
