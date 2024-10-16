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

// #region main
#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let options = get_option()?;
    let pool = sqlx::MySqlPool::connect_with(options).await?;

    println!("Connected");

    let cities = sqlx::query_as::<_, City>("SELECT * FROM city WHERE CountryCode = ?")
        .bind("JPN")
        .fetch_all(&pool)
        .await?;

    println!("日本の都市一覧");
    for city in cities {
        println!("都市名: {}, 人口: {}", city.name, city.population);
    }
    Ok(())
}
// #endregion main
