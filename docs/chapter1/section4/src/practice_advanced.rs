use anyhow::Ok;
use sqlx::mysql::MySqlConnectOptions;
use std::env;

#[derive(sqlx::FromRow)]
#[sqlx(rename_all = "PascalCase")]
#[allow(dead_code)] // 使用していないフィールドへの警告を抑制
struct City {
    #[sqlx(rename = "ID")]
    id: i32,
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

    let city_name = env::args().nth(1).expect("city name is required");
    println!("Connected");
    let city = sqlx::query_as::<_, City>("SELECT * FROM city WHERE Name = ?")
        .bind(&city_name)
        .fetch_one(&pool)
        .await
        .map_err(|e| match e {
            sqlx::Error::RowNotFound => anyhow::anyhow!("no such city Name = {}\n", &city_name),
            _ => anyhow::anyhow!("DB error: {}", e),
        })?;

    println!("{}の人口は{}人です", &city.name, &city.population);

    let population: i64 = sqlx::query_scalar("SELECT Population FROM country WHERE Code = ?") // [!code ++]
        .bind(&city.country_code)  // [!code ++]
        .fetch_one(&pool)  // [!code ++]
        .await  // [!code ++]
        .map_err(|e| match e {  // [!code ++]
            sqlx::Error::RowNotFound => { // [!code ++]
                anyhow::anyhow!("no such country Code = {}\n", &city.country_code) // [!code ++]
            } // [!code ++]
            _ => anyhow::anyhow!("DB error: {}", e), // [!code ++]
        })?; // [!code ++]
    let percent = city.population as f64 / population as f64 * 100.0; // [!code ++]
    println!("これは、{}の人口の{:.2}%です", &city.country_code, percent); // [!code ++]

    Ok(())
}
