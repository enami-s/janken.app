package handler

import (
	"encoding/json"
	"net/http"
)

// ResultListHandler関数の実装
func ResultListHandler(w http.ResponseWriter, r *http.Request) {
	//GETリクエスト以外の場合は、処理を中断する
	if r.Method != http.MethodGet {
		return
	}
	//HttpRequestのリクエストをクライアントに表示する
	json.NewEncoder(w).Encode(r.Method)
	//GetAllの実行

}
