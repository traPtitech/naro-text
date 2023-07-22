package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("failed to get env PORT")
	}

	e := echo.New()

	e.GET("/greeting", greetingHandler)

	e.Logger.Fatal(e.Start(":" + port))
}

func greetingHandler(c echo.Context) error {
	greeting, ok := os.LookupEnv("GREETING_MESSAGE")
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get env GREETING_MESSAGE")
	}

	return c.String(http.StatusOK, greeting)
}
