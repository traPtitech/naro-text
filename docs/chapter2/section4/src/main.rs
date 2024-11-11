use axum::{http::StatusCode, routing::get, Router};

#[tokio::main]
async fn main() {
    let port = std::env::var("PORT")
        .expect("failed to get env PORT")
        .parse::<u16>()
        .expect("failed to parse PORT");

    let app = Router::new().route("/greeting", get(greeting_handler));

    let addr = std::net::SocketAddr::from(([0, 0, 0, 0], port));

    let listener = tokio::net::TcpListener::bind(&addr).await.unwrap();

    println!("Listening on {}", addr);

    axum::serve(listener, app).await.unwrap();
}

async fn greeting_handler() -> Result<(StatusCode, String), StatusCode> {
    let greeting =
        std::env::var("GREETING_MESSAGE").map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?;
    Ok((StatusCode::OK, greeting))
}
