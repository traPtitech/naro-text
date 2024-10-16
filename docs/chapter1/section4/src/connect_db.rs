use anyhow::Ok;
use sqlx::mysql::MySqlConnectOptions;
use std::env;

// #region city
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
// #endregion city  

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

    println!("Connected");
    // #region get
    let city = sqlx::query_as::<_, City>("SELECT * FROM city WHERE Name = ?")
        .bind("Tokyo")
        .fetch_one(&pool)
        .await
        .map_err(|e| match e {
            sqlx::Error::RowNotFound => anyhow::anyhow!("no such city Name = {}\n", "Tokyo"),
            _ => anyhow::anyhow!("DB error: {}", e),
        })?;
    // #endregion get
    println!("Tokyoの人口は{}人です", &city.population);
    Ok(())
}
