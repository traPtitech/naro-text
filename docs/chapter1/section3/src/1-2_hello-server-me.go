package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", func(c echo.Context) error {
		// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// 「/pikachu」というエンドポイントを設定する
	e.GET("/pikachu", func(c echo.Context) error {
		return c.String(http.StatusOK, "始めまして、@pikachuです。\nケモノ(特に四足歩行)や、低頭身デフォルメマスコット(TDM)が大好きです。\n普段はVRChatに生息しています。twitter: @pikachu0310VRC")
	})

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}
