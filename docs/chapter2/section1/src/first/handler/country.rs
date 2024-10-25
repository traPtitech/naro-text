use crate::repository::{country::City, Repository};
use axum::{
    extract::rejection::JsonRejection,
    extract::{Path, State},
    http::StatusCode,
    Json,
};

pub async fn get_city_handler(
    State(state): State<Repository>,
    Path(city_name): Path<String>,
) -> Result<Json<City>, StatusCode> {
    let city = Repository::get_city_by_name(&state, city_name).await;
    match city {
        Ok(city) => Ok(Json(city)),
        Err(sqlx::Error::RowNotFound) => Err(StatusCode::NOT_FOUND),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR),
    }
}

pub async fn post_city_handler(
    State(state): State<Repository>,
    query: Result<Json<City>, JsonRejection>,
) -> Result<Json<City>, StatusCode> {
    match query {
        Ok(Json(city)) => {
            let result = Repository::create_city(&state, city).await;
            match result {
                Ok(city) => Ok(Json(city)),
                Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR),
            }
        }
        Err(_) => Err(StatusCode::BAD_REQUEST),
    }
}
