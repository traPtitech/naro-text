package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// GETリクエストの"/hello/:username"というパターンに対応するルートを設定し、
	// URLのパラメータ(:username)を使用してhelloHandler関数を呼び出す
	e.GET("/hello/:username", helloHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloHandler(c echo.Context) error {
	// ":username"というパスパラメーターを取得し、userID変数に格納する
	userID := c.Param("username")
	return c.String(http.StatusOK, "Hello, "+userID+".\n")
}
