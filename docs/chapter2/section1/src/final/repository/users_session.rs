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

    pub async fn delete_user_session(&self, session_id: String) -> anyhow::Result<()> {
        let session = self
            .session_store
            .load_session(session_id.clone())
            .await
            .with_context(|| "Failed to load session")?
            .with_context(|| "Failed to find session")?;

        self.session_store
            .destroy_session(session)
            .await
            .with_context(|| "Failed to destroy session")?;

        Ok(())
    }

    pub async fn get_user_id_by_session_id(
        &self,
        session_id: &String,
    ) -> anyhow::Result<Option<String>> {
        let session = self
            .session_store
            .load_session(session_id.clone())
            .await
            .with_context(|| "Failed to load session")?;

        Ok(session.and_then(|s| s.get::<String>("user_id")))
    }
}
