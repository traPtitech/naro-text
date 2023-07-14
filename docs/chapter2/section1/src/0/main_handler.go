package main

import (
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)
// #region var
var (
	db *sqlx.DB
	salt = os.Getenv("HASH_SALT")
)
// #endregion var

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

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	db = _db

	if err != nil {
		log.Fatal(err)
	}

	// #region handler
	e := echo.New()
	e.POST("/signup", signUpHandler)

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityHandler)

	e.Start(":8080")
	// #endregion handler
}
