// result_list_handler_test.go
package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// SQLiteにじゃんけんした結果を1レコード挿入する関数
func insertResult(hand string, result string) error {
	//SQLiteを用いてDBを作成する
	db, err := sql.Open("sqlite3", "./janken.db")
	//DBを閉じる
	defer db.Close()
	if err != nil {
		return err
	}
	//DBにじゃんけんの結果を保存する
	_, err = db.Exec("INSERT INTO results (hand, result) VALUES (?, ?)", hand, result)

	if err != nil {
		return err
	}
	return nil
}

// 後処理として、SQLiteのデータの中身を全て削除する関数
func deleteAllResults() error {
	//SQLiteを用いてDBを作成する
	db, err := sql.Open("sqlite3", "./janken.db")
	//DBを閉じる
	defer db.Close()
	if err != nil {
		return err
	}
	//DBの結果を全て削除する
	_, err = db.Exec("DELETE FROM results")
	if err != nil {
		return err
	}
	return nil
}

func TestResultListHandler(t *testing.T) {
	// テストケース1: じゃんけん結果一覧の取得
	t.Run("Get Results with limit and offset", func(t *testing.T) {

		// テストデータを3つ挿入
		insertResult("rock", "win")
		insertResult("scissors", "lose")
		insertResult("paper", "tie")

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/results?limit=2&offset=0", nil)
		recorder := httptest.NewRecorder()

		ResultListHandler(recorder, request)

		var response []map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		// 期待される結果: 3つのうち、2つの結果が取得できる(中身もあっているのか確認)
		if len(response) != 2 {
			t.Errorf("Expected 2 results, but got: %d", len(response))
		}
		if response[0]["hand"] != "rock" || response[0]["result"] != "win" {
			t.Errorf("Expected hand: rock, result: win, but got: hand: %s, result: %s", response[0]["hand"], response[0]["result"])
		}
		if response[1]["hand"] != "scissors" || response[1]["result"] != "lose" {
			t.Errorf("Expected hand: scissors, result: lose, but got: hand: %s, result: %s", response[1]["hand"], response[1]["result"])
		}

		// 後処理として、テストデータを全て削除
		deleteAllResults()
	})

	// テストケース2: 別のoffsetでのじゃんけん結果一覧の取得
	t.Run("Get Results with different limit and offset", func(t *testing.T) {
		// テストデータを5つ挿入
		insertResult("rock", "win")
		insertResult("scissors", "lose")
		insertResult("paper", "tie")
		insertResult("rock", "win")
		insertResult("scissors", "lose")

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/results?limit=2&offset=1", nil)
		recorder := httptest.NewRecorder()

		ResultListHandler(recorder, request)

		var response []map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		// 期待される結果: 5つのうち、2つの結果が取得できる(中身もあっているのか確認)
		if len(response) != 2 {
			t.Errorf("Expected 2 results, but got: %d", len(response))
		}
		if response[0]["hand"] != "paper" || response[0]["result"] != "tie" {
			t.Errorf("Expected hand: paper, result: tie, but got: hand: %s, result: %s", response[0]["hand"], response[0]["result"])
		}
		if response[1]["hand"] != "rock" || response[1]["result"] != "win" {
			t.Errorf("Expected hand: rock, result: win, but got: hand: %s, result: %s", response[1]["hand"], response[1]["result"])
		}
	})
}
