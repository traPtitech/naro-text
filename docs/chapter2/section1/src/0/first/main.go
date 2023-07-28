package main

import (
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/naro-template-backend/handler"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// データーベースの設定
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

	// データベースに接続
	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(db)
	e := echo.New()

	e.GET("/cities/:cityName", h.GetCityInfoHandler)
	e.POST("/cities", h.PostCityHandler)

	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
