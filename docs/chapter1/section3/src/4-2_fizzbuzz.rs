use axum::{extract::Query, http::StatusCode, routing::get, Router};
#[tokio::main]
async fn main() {
    // 「/ping」というエンドポイントを設定する
    let app = Router::new().route("/fizzbuzz", get(fizzbuzz_handler));

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
struct FizzBuzzQuery {
    count: Option<String>,
}

async fn fizzbuzz_handler(Query(query): Query<FizzBuzzQuery>) -> (StatusCode, String) {
    // クエリパラメータが指定されていない場合はデフォルト値を使用する
    let mut n: i32 = 30;
    // クエリパラメータが指定されている場合はその値を調べる
    if let Some(count) = query.count {
        let count = count.parse();
        match count {
            // 数値に変換できた場合はその値を使用する
            Ok(count) => n = count,
            // ステータスコード 400 Bad Request を返す
            Err(_) => return (StatusCode::BAD_REQUEST, String::from("Bad Request\n")),
        }
    }

    // FizzBuzzの処理をする
    let fizzbuzz_str = fizzbuzz(n);

    // ステータスコード 200 Ok とfizzBuzzの結果を返す
    (StatusCode::OK, fizzbuzz_str + "\n")
}

// fizzBuzzの処理
fn fizzbuzz(n: i32) -> String {
    (1..=n)
        .map(|i| match (i % 3, i % 5) {
            (0, 0) => "FizzBuzz".to_string(),
            (0, _) => "Fizz".to_string(),
            (_, 0) => "Buzz".to_string(),
            (_, _) => i.to_string(),
        })
        .collect::<Vec<String>>()
        .join("\n")
}
