use sqlx::mysql::MySqlConnectOptions;
use sqlx::mysql::MySqlPool;
use std::env;

pub mod country;

#[derive(Clone)]
pub struct Repository {
    pool: MySqlPool,
}

impl Repository {
    pub async fn connect() -> anyhow::Result<Self> {
        let options = get_options()?;
        let pool = sqlx::MySqlPool::connect_with(options).await?;
        Ok(Self {
            pool,
        })
    }
}

fn get_options() -> anyhow::Result<MySqlConnectOptions> {
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
