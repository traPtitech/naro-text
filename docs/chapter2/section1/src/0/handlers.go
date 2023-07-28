package main

import (
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

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

	// #region setup_table
	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	db = _db

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))") // [!code ++]
	// [!code ++]
	if err != nil { // [!code ++]
		log.Fatal(err) // [!code ++]
	} // [!code ++]
	// #endregion setup_table

	// #region signup
	e := echo.New()
	e.POST("/signup", signUpHandler) // [!code ++]

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler)
	// #endregion signup

	e.Start(":8080")
}
