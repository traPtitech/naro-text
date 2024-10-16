use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    routing::{get, post},
    Json, Router,
};
use sqlx::{mysql::MySqlConnectOptions, Pool};
use std::env;

#[derive(sqlx::FromRow, serde::Serialize, serde::Deserialize)] // [!code warning]
#[sqlx(rename_all = "PascalCase")]
#[serde(rename_all = "camelCase")]
struct City {
    #[sqlx(rename = "ID")]
    id: Option<i32>, // [!code warning]
    name: String,
    country_code: String,
    district: String,
    population: i32,
}

fn get_option() -> anyhow::Result<MySqlConnectOptions> {
    let host = env::var("DB_HOSTNAME")?;
    let port = env::var("DB_PORT")?.parse()?;
    let username = env::var("DB_USERNAME")?;
    let password = env::var("DB_PASSWORD")?;
    let database = env::var("DB_DATABASE")?;
    let timezone = Some(String::from("Asia/Tokyo"));
    let collation = String::from("utf8mb4_unicode_ci");

    Ok(MySqlConnectOptions::new()
        .host(&host)
        .port(port)
        .username(&username)
        .password(&password)
        .database(&database)
        .timezone(timezone)
        .collation(&collation))
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let options = get_option()?;
    let pool = sqlx::MySqlPool::connect_with(options).await?;

    let app = Router::new()
        .route("/cities/:cityName", get(get_city_handler)) 
        .route("/cities", post(post_city_handler)) // [!code ++]
        .with_state(pool);

    let listener = tokio::net::TcpListener::bind("127.0.0.1:8080")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();

    Ok(())
}

async fn get_city_handler(
    State(pool): State<Pool<sqlx::MySql>>,
    Path(city_name): Path<String>,
) -> Result<Json<City>, StatusCode> {
    let city = sqlx::query_as::<_, City>("SELECT * FROM city WHERE Name = ?")
        .bind(&city_name)
        .fetch_one(&pool)
        .await;

    match city {
        Ok(city) => Ok(Json(city)),
        Err(sqlx::Error::RowNotFound) => Err(StatusCode::NOT_FOUND),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR),
    }
}

async fn post_city_handler( // [!code ++]
    State(pool): State<Pool<sqlx::MySql>>, // [!code ++]
    query: Result<Json<City>, JsonRejection>, // [!code ++]
) -> Result<Json<City>, StatusCode> { // [!code ++]
    match query { // [!code ++]
        Ok(Json(mut city)) => {     // [!code ++]
            let result = sqlx::query( // [!code ++]
                "INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)", // [!code ++]
            ) // [!code ++]
            .bind(&city.name) // [!code ++]
            .bind(&city.country_code) // [!code ++]
            .bind(&city.district) // [!code ++]
            .bind(city.population) // [!code ++]
            .execute(&pool) // [!code ++]
            .await // [!code ++]
            .map_err(|_| StatusCode::INTERNAL_SERVER_ERROR)?; // [!code ++]

            city.id = Some(result.last_insert_id() as i32); // [!code ++]
            Ok(Json(city)) // [!code ++]
        } // [!code ++]
        Err(_) => Err(StatusCode::BAD_REQUEST), // [!code ++]
    }   // [!code ++]
}  // [!code ++]
