use crate::repository::{country::City, Repository};
use axum::{
    extract::rejection::JsonRejection,
    extract::{Path, State},
    http::StatusCode,
    Json,
};

pub async fn get_city_handler(
    State(state): State<Repository>,
    Path(city_name): Path<String>,
) -> Result<Json<City>, StatusCode> {
    let city = Repository::get_city_by_name(&state, city_name).await;
    match city {
        Ok(city) => Ok(Json(city)),
        Err(sqlx::Error::RowNotFound) => Err(StatusCode::NOT_FOUND),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR),
    }
}

pub async fn post_city_handler(
    State(state): State<Repository>,
    query: Result<Json<City>, JsonRejection>,
) -> Result<Json<City>, StatusCode> {
    match query {
        Ok(Json(city)) => {
            let result = Repository::create_city(&state, city).await;
            match result {
                Ok(city) => Ok(Json(city)),
                Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR),
            }
        }
        Err(_) => Err(StatusCode::BAD_REQUEST),
    }
}

use std::collections::HashMap;

//与えられた City のリストから国ごとの人口の和を計算する
#[allow(dead_code)]
pub fn sum_population_by_country(cities: Vec<City>) -> HashMap<String, i32> {
    let mut map = HashMap::new();
    for city in cities {
        if city.country_code.is_empty() {
            continue;
        }
        let entry = map.entry(city.country_code).or_insert(0);
        *entry += city.population;
    }
    map
}

// #[cfg(test)] 属性を追加したモジュールはテストモジュールとして扱われる
#[cfg(test)]
mod tests {
    use super::{sum_population_by_country, City};
    use std::collections::HashMap;

    #[test]
    fn test_sum_population_by_country_empty() {
        // ここにテストを追加する
        let cities = vec![];
        let result = sum_population_by_country(cities);
        assert!(result.is_empty());
    }

    #[test]
    fn test_sum_population_by_country_single() {
        let cities = vec![
            City {
                id: Some(1),
                name: "Tokyo".to_string(),
                country_code: "JPN".to_string(),
                district: "Tokyo".to_string(),
                population: 100,
            },
            City {
                id: Some(2),
                name: "Osaka".to_string(),
                country_code: "JPN".to_string(),
                district: "Osaka".to_string(),
                population: 200,
            },
        ];

        let mut expected = HashMap::new();
        expected.insert("JPN".to_string(), 300);

        let result = sum_population_by_country(cities);

        assert_eq!(result, expected);
    }

    #[test]
    fn test_sum_population_by_country_multiple() {
        let cities = vec![
            City {
                id: Some(1),
                name: "Tokyo".to_string(),
                country_code: "JPN".to_string(),
                district: "Tokyo".to_string(),
                population: 100,
            },
            City {
                id: Some(2),
                name: "Osaka".to_string(),
                country_code: "JPN".to_string(),
                district: "Osaka".to_string(),
                population: 200,
            },
            City {
                id: Some(3),
                name: "New York".to_string(),
                country_code: "USA".to_string(),
                district: "New York".to_string(),
                population: 300,
            },
            City {
                id: Some(4),
                name: "Los Angeles".to_string(),
                country_code: "USA".to_string(),
                district: "California".to_string(),
                population: 400,
            },
        ];

        let mut expected = HashMap::new();
        expected.insert("JPN".to_string(), 300);
        expected.insert("USA".to_string(), 700);

        let result = sum_population_by_country(cities);

        assert_eq!(result, expected);
    }

    #[test]
    fn test_sum_population_by_country_empty_country_code() {
        let cities = vec![
            City {
                id: Some(1),
                name: "Tokyo".to_string(),
                country_code: "JPN".to_string(),
                district: "Tokyo".to_string(),
                population: 100,
            },
            City {
                id: Some(2),
                name: "Osaka".to_string(),
                country_code: "".to_string(),
                district: "Osaka".to_string(),
                population: 200,
            },
        ];

        let mut expected = HashMap::new();
        expected.insert("JPN".to_string(), 100);

        let result = sum_population_by_country(cities);

        assert_eq!(result, expected);
    }
}
