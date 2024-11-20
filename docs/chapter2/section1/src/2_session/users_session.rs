use anyhow::Context;
use async_session::{Session, SessionStore};

use super::Repository;

impl Repository {
    pub async fn create_user_session(&self, user_id: String) -> anyhow::Result<String> {
        let mut session = Session::new();

        session
            .insert("user_id", user_id)
            .with_context(|| "Failed to insert user_id")?;

        let session_id = self
            .session_store
            .store_session(session)
            .await
            .with_context(|| "Failed to store session")?
            .with_context(|| "Failed to create session")?;

        Ok(session_id)
    }
}
