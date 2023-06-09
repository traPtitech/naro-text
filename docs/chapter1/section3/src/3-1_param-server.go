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
	return c.String(http.StatusOK, "Hello, "+userID+".\n")
}
