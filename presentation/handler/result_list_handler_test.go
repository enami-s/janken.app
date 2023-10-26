// result_list_handler_test.go
package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResultListHandler(t *testing.T) {
	// テストケース4: じゃんけん結果一覧の取得
	t.Run("Get Results with limit and offset", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/results?limit=2&offset=0", nil)
		recorder := httptest.NewRecorder()

		ResultListHandler(recorder, request)

		var response []map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		if len(response) != 2 {
			t.Errorf("Expected 2 results, but got: %d", len(response))
		}
	})

	// テストケース5: 別のoffsetでのじゃんけん結果一覧の取得
	t.Run("Get Results with different limit and offset", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/results?limit=2&offset=1", nil)
		recorder := httptest.NewRecorder()

		ResultListHandler(recorder, request)

		var response []map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		if len(response) != 2 {
			t.Errorf("Expected 2 results, but got: %d", len(response))
		}

	})
}
