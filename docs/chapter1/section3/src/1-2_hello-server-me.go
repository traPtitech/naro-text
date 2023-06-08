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

	e.GET("/pikachu", func(c echo.Context) error {
		return c.String(http.StatusOK, "始めまして、@pikachuです。\nケモノ(特に四足歩行)や、低頭身デフォルメマスコット(TDM)が大好きです。\nVRChatに生息しています。")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
