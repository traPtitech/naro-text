package main

import (
	"database/sql"
	"maps"
	"testing"
)

func Test_sumPopulationByCountryCode(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{}
	got := sumPopulationByCountryCode(cities)
	want := map[string]int{}
	// 長さが0になっているかどうかを確認する
	if len(got) != 0 {
		t.Errorf("sumPopulationByCountryCode(%v) = %v, want %v", cities, got, want)
	}
}

func Sample1(t *testing.T) {
	cases := []struct {
		name   string
		cities []City
		want   map[string]int64
	}{
		{
			name:   "empty input",
			cities: []City{},
			want:   map[string]int64{},
		},
		// #region single
		{
			name: "single country",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100},
		},
		// #endregion single
		// #region multiple
		{
			name: "multiple countries",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
				{
					CountryCode: sql.NullString{
						String: "USA",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 200,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100, "USA": 200},
		},
		// #endregion multiple
		// #region null
		{
			name: "empty country code",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "",
						Valid:  false,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{},
		},
		// #endregion null
	}
	for _, tt := range cases {
		// サブテストの実行
		t.Run(tt.name, func(t *testing.T) {
			got := sumPopulationByCountryCode(tt.cities)
			if !maps.Equal(got, tt.want) {
				t.Errorf("sumPopulationByCountryCode(%v) = %v, want %v", tt.cities, got, tt.want)
			}
		})
	}
}

func Sample2(t *testing.T) {
	// #region test_cases
	cases := []struct {
		name   string
		cities []City
		want   map[string]int64
	}{
		{
			name:   "empty input",
			cities: []City{},
			want:   map[string]int64{},
		},
		{
			name: "single country",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100},
		},
		{
			name: "multiple countries",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
				{
					CountryCode: sql.NullString{
						String: "USA",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 200,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100, "USA": 200},
		},
		{
			name: "empty country code",
			cities: []City{
				{
					CountryCode: sql.NullString{
						String: "",
						Valid:  false,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{},
		},
	}
	// #endregion test_cases
	for _, tt := range cases {
		// サブテストの実行
		t.Run(tt.name, func(t *testing.T) {
			got := sumPopulationByCountryCode(tt.cities)
			if !maps.Equal(got, tt.want) {
				t.Errorf("sumPopulationByCountryCode(%v) = %v, want %v", tt.cities, got, tt.want)
			}
		})
	}
}
