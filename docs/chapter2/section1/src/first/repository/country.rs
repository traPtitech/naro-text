use super::Repository;

#[derive(sqlx::FromRow, serde::Serialize, serde::Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct City {
    #[sqlx(rename = "ID")]
    pub id: Option<i32>,
    #[sqlx(rename = "Name")]
    pub name: String,
    #[sqlx(rename = "CountryCode")]
    pub country_code: String,
    #[sqlx(rename = "District")]
    pub district: String,
    #[sqlx(rename = "Population")]
    pub population: i32,
}

impl Repository {
    pub async fn get_city_by_name(&self, city_name: String) -> sqlx::Result<City> {
        sqlx::query_as::<_, City>("SELECT * FROM city WHERE Name = ?")
            .bind(&city_name)
            .fetch_one(&self.pool)
            .await
    }

    pub async fn create_city(&self, city: City) -> sqlx::Result<City> {
        let result = sqlx::query(
            "INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)",
        )
        .bind(&city.name)
        .bind(&city.country_code)
        .bind(&city.district)
        .bind(city.population)
        .execute(&self.pool)
        .await?;

        let id = result.last_insert_id() as i32;
        Ok(City {
            id: Some(id),
            ..city
        })
    }
}
