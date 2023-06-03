package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"ID,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}
//#region main
func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conntected")

	var cities []City
	err = db.Select(&cities, "SELECT * FROM city WHERE CountryCode = 'JPN'") //?を使わない場合、第3引数以降は不要
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("日本の都市一覧")
	for _, city := range cities {
		fmt.Printf("都市名: %s, 人口: %d\n", city.Name, city.Population)
	}
}
//#endregion main