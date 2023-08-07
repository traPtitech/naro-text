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

// #region city
type City struct {
	ID          int    `json:"ID,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

// #endregion city
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
	// #region get
	var city City
	err = db.Get(&city, "SELECT * FROM city WHERE Name = ?", "Tokyo")
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = '%s'\n", "Tokyo")
	} else if err != nil {
		log.Fatalf("DB Error: %s\n", err)
	}
	// #endregion get
	fmt.Printf("Tokyoの人口は%d人です\n", city.Population)
}
