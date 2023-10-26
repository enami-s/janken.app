// janken_handler_test.go
package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPlayJankenHandler(t *testing.T) {
	// テストケース1: じゃんけんの手の選択、勝敗の判定と結果表示
	t.Run("Valid Hand Input - Rock", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/rps", strings.NewReader("hand=rock"))
		recorder := httptest.NewRecorder()

		PlayJankenHandler(recorder, request)

		var response map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		// 期待される結果: "win", "lose", or "tie"
		if response["result"] != "win" && response["result"] != "lose" && response["result"] != "tie" {
			t.Errorf("Unexpected result: %s", response["result"])
		}
	})

	// テストケース2: 不正な入力のハンドリング
	t.Run("Invalid Hand Input", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/rps", strings.NewReader("hand=test"))
		recorder := httptest.NewRecorder()

		PlayJankenHandler(recorder, request)

		var response map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		// 期待されるエラーメッセージ
		expectedError := "rock,scissors,paperのどれかを入力し直してください"
		if response["error"] != expectedError {
			t.Errorf("Expected error message: %s, but got: %s", expectedError, response["error"])
		}
	})

}
