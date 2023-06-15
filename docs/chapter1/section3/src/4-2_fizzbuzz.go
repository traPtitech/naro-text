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
	// クエリパラメーター "count" の値をcountStrに格納
	countStr := c.QueryParam("count")
	// もしcountStrが空っぽ("")だったら、
	if countStr == "" {
		// countStrを30に設定
		countStr = "30"
	}

	// strconv.Atoi()でstringをintに変換する(分からなかったらググりましょう！)
	count, err := strconv.Atoi(countStr)
	// ↑でエラーが起きたら、"count" が整数として解釈できないということなので、
	if err != nil {
		// ステータスコード 400 Bad Request を返す
		return c.String(http.StatusBadRequest, "Bad Request\n")
	}

	// fizzBuzzの処理を行う
	fizzBuzzStr := fizzBuzz(count)

	// ステータスコード 200 Ok とfizzBuzzの結果を返す
	return c.String(http.StatusOK, fizzBuzzStr)
}

// fizzBuzzの処理。これは競プロ
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
