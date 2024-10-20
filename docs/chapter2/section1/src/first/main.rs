use tower_http::trace::TraceLayer;
use tracing_subscriber::EnvFilter;

mod handler;
mod repository;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::try_from_default_env().unwrap_or("info".into()))
        .init();

    let app_state = repository::Repository::connect().await?;
    let app = handler::make_router(app_state).layer(TraceLayer::new_for_http());
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await?;

    tracing::info!("listening on {}", listener.local_addr()?);
    axum::serve(listener, app).await.unwrap();
    Ok(())
}
