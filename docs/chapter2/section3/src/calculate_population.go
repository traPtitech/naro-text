package main

import (
	"database/sql"
)

type City struct {
	ID          int            `json:"id,omitempty"  db:"ID"`
	Name        sql.NullString `json:"name,omitempty"  db:"Name"`
	CountryCode sql.NullString `json:"countryCode,omitempty"  db:"CountryCode"`
	District    sql.NullString `json:"district,omitempty"  db:"District"`
	Population  sql.NullInt64  `json:"population,omitempty"  db:"Population"`
}

// #region calculate
func sumPopulationByCountryCode(cities []City) map[string]int64 {
	result := make(map[string]int64)
	for _, city := range cities {
		if city.CountryCode.Valid {
			// まだmapに存在しなかった場合、初期化する
			if _, ok := result[city.CountryCode.String]; !ok {
				result[city.CountryCode.String] = 0
			}
			result[city.CountryCode.String] += city.Population.Int64
		}
	}
	return result
}

// #endregion calculate
