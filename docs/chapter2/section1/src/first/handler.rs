use axum::{
    routing::{get, post},
    Router,
};

use crate::repository::Repository;
mod country;

pub fn make_router(app_state: Repository) -> Router {
    let city_router = Router::new()
        .route("/city/:city_name", get(country::get_city_handler))
        .route("/cities", post(country::post_city_handler));

    Router::new().nest("/", city_router).with_state(app_state)
}
