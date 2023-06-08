package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", func(c echo.Context) error {
		// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// 8080という数字は、WebサーバーがListenするポート番号
	e.Logger.Fatal(e.Start(":8080"))
}
