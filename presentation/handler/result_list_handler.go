package handler

import (
	"encoding/json"
	"github.com/enami-s/janken_app/infrastructure/repository"
	"net/http"
	"strconv"
)

// ResultListHandler関数の実装
func ResultListHandler(w http.ResponseWriter, r *http.Request) {
	//GETリクエスト以外の場合は、エラーメッセージを出力する
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	//HttpRequestのリクエストをクライアントに表示する
	json.NewEncoder(w).Encode(r.Method)
	//repositoryのNewJankenRepository関数を実行
	repo := repository.NewJankenRepository()
	// クエリパラメータからlimitとoffsetの値を取得する
	queryParams := r.URL.Query()
	limitStr := queryParams.Get("limit")
	offsetStr := queryParams.Get("offset")

	// limitとoffsetの値をint型に変換する
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//GetAllの実行
	resultList, err := repo.GetAll(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//結果をクライアントに返す
	json.NewEncoder(w).Encode(resultList)
}
