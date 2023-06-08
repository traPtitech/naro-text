package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// JSONで返すための構造体を定義
type jsonData struct {
	Number int
	String string
	Bool   bool
}

func main() {
	e := echo.New()

	// 「/hello」というパスのエンドポイントを定義
	e.GET("/hello", func(c echo.Context) error {
		// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// 「/json」というパスのエンドポイントを定義
	e.GET("/json", jsonHandler)

	// 8080という数字は、WebサーバーがListenするポート番号
	e.Logger.Fatal(e.Start(":8080"))
}

func jsonHandler(c echo.Context) error {
	// レスポンスとして返す値を構造体として定義
	response := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}

	// HTTPステータスコードは200番で、構造体をJSONに変換してクライアントに返す
	return c.JSON(http.StatusOK, &response)
}
