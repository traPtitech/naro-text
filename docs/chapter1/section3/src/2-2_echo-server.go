package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Jsonで受け取り、構造体に変換して返すための構造体を定義
type jsonData struct {
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}

func main() {
	e := echo.New()

	// `e.GET`と同じように、`e.POST`と書くことで POST を受け取ることができます。
	e.POST("/post", postHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func postHandler(c echo.Context) error {
	// 受け取りたい JSON を示す空の変数を先に用意する。
	data := &jsonData{}
	// 受け取った JSON を data に代入する
	err := c.Bind(data)

	if err != nil { // エラーが発生した時、以下を実行
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}
	// エラーが起きなかったとき、正常なのでステータスコード 200 OK を返し、リクエストデータをそのまま返す
	return c.JSON(http.StatusOK, data)
}
