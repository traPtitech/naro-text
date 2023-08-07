package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"ID,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	conf := mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_DATABASE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")

	cityName := os.Args[1]

	var city City
	err = db.Get(&city, "SELECT * FROM city WHERE Name = ?", cityName)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = '%s'\n", cityName)
		return
	} else if err != nil {
		log.Fatalf("DB Error: %s\n", err)
	}

	fmt.Printf("%sの人口は%d人です\n", city.Name, city.Population)

	var population int                                                                           //[!code ++]
	err = db.Get(&population, "SELECT Population FROM country WHERE Code = ?", city.CountryCode) //[!code ++]
	if errors.Is(err, sql.ErrNoRows) {                                                           //[!code ++]
		log.Printf("no such country Code = '%s'\n", city.CountryCode) //[!code ++]
		return                                                        //[!code ++]
	} else if err != nil { //[!code ++]
		log.Fatalf("DB Error: %s\n", err) //[!code ++]
	} //[!code ++]
	//[!code ++]
	percent := (float64(city.Population) / float64(population)) * 100 //[!code ++]
	//[!code ++]
	fmt.Printf("これは%sの人口の%f%%です\n", city.CountryCode, percent) //[!code ++]
}
