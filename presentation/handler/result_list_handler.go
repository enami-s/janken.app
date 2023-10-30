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

	//ErrがNilでなかったら、エラーメッセージを出力する
	if err != nil {
		http.Error(w, "limit must be between 0 and 100", http.StatusBadRequest)
		return
	}
	//0＜Limit＜100の範囲外の場合は、エラーメッセージを出力する
	if limit < 0 || limit > 100 {
		http.Error(w, "limit must be between 0 and 100", http.StatusBadRequest)
		return
	}

	//ErrがNilでなかったら、エラーメッセージを出力する
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//0＜Offset＜2147483647の範囲外の場合はエラーメッセージを出力する
	if offset < 0 || offset > 2147483647 {
		http.Error(w, "offset must be between 0 and 2147483647", http.StatusBadRequest)
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
