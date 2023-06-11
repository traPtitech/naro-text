package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Jsonで受け取って構造体に変換するための構造体を定義
type requestData struct {
	Left  *int `json:"left"`
	Right *int `json:"right,omitempty"`
}

// 計算結果をJsonで返すための構造体を定義
type responseData struct {
	Answer int `json:"answer"`
}

// エラーをJsonで返すための構造体を定義
type errorData struct {
	Error string `json:"error,omitempty"`
}

func main() {
	e := echo.New()

	// 「/add」というエンドポイントを設定する
	e.POST("/add", addHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func addHandler(c echo.Context) error {
	// 受け取りたい JSON を示す空の変数を先に用意する。
	numbers := &requestData{}
	// 受け取った JSON を data に代入する
	err := c.Bind(numbers)

	// JSON を data に代入する時にエラーが発生したり、値が空だった時、つまりリクエストが適切ではなかった時、
	if err != nil || numbers.Left == nil || numbers.Right == nil {
		// ステータスコード 400 Bad Request で"Bad Request"が入ったJSONを返す
		return c.JSON(http.StatusBadRequest, &errorData{Error: "Bad Request"})
	}

	// 計算結果
	result := *numbers.Left + *numbers.Right
	// 計算結果をステータスコード 200 Ok で返す
	return c.JSON(http.StatusOK, &responseData{Answer: result})
}
