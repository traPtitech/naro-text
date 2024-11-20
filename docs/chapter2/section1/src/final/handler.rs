use axum::{
    middleware::from_fn_with_state,
    routing::{get, post},
    Router,
};

use crate::repository::Repository;

mod auth;
mod country;

pub fn make_router(app_state: Repository) -> Router {
    let with_auth_router = Router::new()
        .route("/cities/:city_name", get(country::get_city_handler))
        .route("/cities", post(country::post_city_handler))
        .route("/me", get(auth::me))
        .route_layer(from_fn_with_state(app_state.clone(), auth::auth_middleware));

    let auth_router = Router::new()
        .route("/signup", post(auth::sign_up))
        .route("/login", post(auth::login))
        .route("/logout", post(auth::logout));

    let ping_router = Router::new().route("/ping", get(|| async { "pong" }));

    Router::new()
        .nest("/", with_auth_router)
        .nest("/", auth_router)
        .nest("/", ping_router)
        .with_state(app_state)
}
