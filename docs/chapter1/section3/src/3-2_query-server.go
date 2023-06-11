package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello/:username", helloHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloHandler(c echo.Context) error {
	userID := c.Param("username")

	// クエリパラメータ "lang" の値を取得し、language 変数に格納する
	// 例: /hello/pikachu?lang=ja であれば、language は "ja" になる
	language := c.QueryParam("lang")

	// 同様にクエリパラメータ "page" の値を取得し、pageNum 変数に格納する
	pageNum := c.QueryParam("page")

	return c.String(http.StatusOK, "Hello, "+userID+".\nlanguage: "+language+"\npage: "+pageNum)
}
