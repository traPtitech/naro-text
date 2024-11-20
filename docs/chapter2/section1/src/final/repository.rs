use async_sqlx_session::MySqlSessionStore;
use sqlx::mysql::MySqlConnectOptions;
use sqlx::mysql::MySqlPool;
use std::env;

pub mod country;
pub mod users;
pub mod users_session;

#[derive(Clone)]
pub struct Repository {
    pool: MySqlPool,
    session_store: MySqlSessionStore,
}

impl Repository {
    pub async fn connect() -> anyhow::Result<Self> {
        let options = get_options()?;
        let pool = sqlx::MySqlPool::connect_with(options).await?;

        let session_store =
            MySqlSessionStore::from_client(pool.clone()).with_table_name("user_sessions");

        Ok(Self {
            pool,
            session_store,
        })
    }

    pub async fn migrate(&self) -> anyhow::Result<()> {
        sqlx::migrate!("./migrations").run(&self.pool).await?;
        Ok(())
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
