#region calculate
use std::collections::HashMap;

//与えられた City のリストから国ごとの人口の和を計算する
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
#endregion calculate

// #[cfg(test)] 属性を追加したモジュールはテストモジュールとして扱われる
#[cfg(test)]
mod tests {
    use super::{sum_population_by_country, City};
    use std::collections::HashMap;

	#region empty
    #[test]
    fn test_sum_population_by_country_empty() {
        // ここにテストを追加する
        let cities = vec![];
        let result = sum_population_by_country(cities);
        assert!(result.is_empty());
    }
	#endregion empty

	#region single
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
	#endregion single

	#region multiple
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
	#endregion multiple

	#region empty_country_code
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
	#endregion empty_country_code
}
