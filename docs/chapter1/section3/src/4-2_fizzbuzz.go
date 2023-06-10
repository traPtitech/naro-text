package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// 「/fizzbuzz」というエンドポイントを設定する
	e.GET("/fizzbuzz", fizzBuzzHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func fizzBuzzHandler(c echo.Context) error {
	countStr := c.QueryParam("count")
	if countStr == "" {
		countStr = "30"
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request\n")
	}

	fizzBuzzStr := fizzBuzz(count)
	return c.String(http.StatusOK, fizzBuzzStr)
}

func fizzBuzz(n int) string {
	result := ""
	for i := 1; i <= n; i++ {
		fizz := i%3 == 0
		buzz := i%5 == 0

		if fizz && buzz {
			result += "FizzBuzz\n"
		} else if fizz {
			result += "Fizz\n"
		} else if buzz {
			result += "Buzz\n"
		} else {
			result += fmt.Sprintf("%d\n", i)
		}
	}

	return result
}
