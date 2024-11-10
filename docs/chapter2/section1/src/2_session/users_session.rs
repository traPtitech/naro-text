use anyhow::Context;
use async_session::{Session, SessionStore};

use super::Repository;

impl Repository {
    pub async fn create_user_session(&self, user_id: String) -> anyhow::Result<()> {
        let mut session = Session::new();

        session
            .insert("user_id", user_id)
            .with_context(|| "Failed to insert user_id")?;

        let result = self
            .session_store
            .store_session(session)
            .await
            .with_context(|| "Failed to store session")
            .with_context(|| "Failed to store session")?;

        match result {
            Some(_) => Ok(()),
            None => Err(anyhow::anyhow!("Failed to store session")),
        }
    }
}
