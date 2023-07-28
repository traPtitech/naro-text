package main

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	db *sqlx.DB
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".envファイルが読み込めませんでした。: %v", err)
	}

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

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	// #region setup_table
	db = _db

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))")

	if err != nil {
		log.Fatal(err)
	} // [!code ++]
	// #endregion setup_table

	// #region handler
	e := echo.New()
	e.POST("/signup", signUpHandler)

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler)

	e.Start(":8080")
	// #endregion handler
}
