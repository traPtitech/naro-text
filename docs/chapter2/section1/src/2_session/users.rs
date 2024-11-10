use super::Repository;

impl Repository {
    pub async fn is_exist_username(&self, username: String) -> sqlx::Result<bool> {
        let result = sqlx::query("SELECT * FROM users WHERE username = ?")
            .bind(&username)
            .fetch_optional(&self.pool)
            .await?;
        Ok(result.is_some())
    }

    pub async fn create_user(&self, username: String) -> sqlx::Result<u64> {
        let result = sqlx::query("INSERT INTO users (username) VALUES (?)")
            .bind(&username)
            .execute(&self.pool)
            .await?;
        Ok(result.last_insert_id())
    }

    pub async fn get_user_id_by_name(&self, username: String) -> sqlx::Result<u64> {
        let result = sqlx::query_scalar("SELECT id FROM users WHERE username = ?")
            .bind(&username)
            .fetch_one(&self.pool)
            .await?;
        Ok(result)
    }

    pub async fn save_user_password(&self, id: i32, password: String) -> anyhow::Result<()> {
        let hash = bcrypt::hash(password, bcrypt::DEFAULT_COST)?;

        sqlx::query("INSERT INTO user_passwords (id, hashed_pass) VALUES (?, ?)")
            .bind(id)
            .bind(hash)
            .execute(&self.pool)
            .await?;

        Ok(())
    }

    pub async fn verify_user_password(&self, id: u64, password: String) -> anyhow::Result<bool> {
        let hash =
            sqlx::query_scalar::<_, String>("SELECT hashed_pass FROM user_passwords WHERE id = ?")
                .bind(id)
                .fetch_one(&self.pool)
                .await?;

        Ok(bcrypt::verify(password, &hash)?)
    }
}
