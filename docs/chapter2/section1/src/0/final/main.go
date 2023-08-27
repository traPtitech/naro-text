package main

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
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

	// usersテーブルが存在しなかったら、usersテーブルを作成する
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (Username VARCHAR(255) PRIMARY KEY, HashedPass VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
	}

	// セッションの情報を記憶するための場所をデータベース上に設定
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(db)
	e := echo.New()
	e.Use(middleware.Logger())       // ログを取るミドルウェアを追加
	e.Use(session.Middleware(store)) // セッション管理のためのミドルウェアを追加

	e.POST("/signup", h.SignUpHandler)
	e.POST("/login", h.LoginHandler)

	withAuth := e.Group("")
	withAuth.Use(handler.UserAuthMiddleware)
	withAuth.GET("/me", handler.GetMeHandler)
	withAuth.GET("/cities/:cityName", h.GetCityInfoHandler)
	withAuth.POST("/cities", h.PostCityHandler)

	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
