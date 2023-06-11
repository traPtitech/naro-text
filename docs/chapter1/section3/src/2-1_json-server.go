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
	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/json」というパスのエンドポイントを定義
	e.GET("/json", jsonHandler)

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
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
