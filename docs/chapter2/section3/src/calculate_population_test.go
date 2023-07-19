package main

import (
	"database/sql"
	"testing"
)

func Test_calculatePopulation(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{}
	got := calculatePopulation(cities)
	want := map[string]int64{}
	// 長さが0になっているかどうかを確認する
	if len(got) != 0 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

// #region single
// 1 つの国のみのデータが入っている場合
func Test_calculatePopulation_single(t *testing.T) {
	cities := []City{
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
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 300,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	// JPNの人口が100になっているかどうかを確認する
	if got["JPN"] != want["JPN"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

// #endregion single

// #region multiple
// 複数の国のデータが入っている場合
func Test_calculatePopulation_multiple(t *testing.T) {
	cities := []City{
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
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 300,
				Valid: true,
			},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 300,
		"USA": 300,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	for k, v := range got {
		// 国ごとの人口が一致しているかどうかを確認する
		if v != want[k] {
			t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
		}
	}
}

// #endregion multiple

// #region null
// 空のデータ(`city.CountryCode.Valid = false`)が入っている場合
func Test_calculatePopulation_null(t *testing.T) {
	cities := []City{
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
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 200,
	}
	// 長さが1になっているかどうかを確認する
	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	// JPNの人口が100になっているかどうかを確認する
	if got["JPN"] != want["JPN"] {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}

// #endregion null
