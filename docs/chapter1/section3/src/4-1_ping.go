package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// 「/ping」というエンドポイントを設定する
	e.GET("/ping", func(c echo.Context) error {
		// HTTPステータスコードは200番で、文字列「pong」をクライアントに返す
		return c.String(http.StatusOK, "pong\n")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
