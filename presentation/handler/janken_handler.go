package handler

import (
	"encoding/json"
	"github.com/enami-s/janken_app/application/service"
	"github.com/enami-s/janken_app/domain/model"
	"net/http"
)

// PlayJanlenHandler関数の実装
func PlayJankenHandler(w http.ResponseWriter, r *http.Request) {
	//POSTリクエスト以外の場合は、httpErrorを返すようにする
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	//POSTリクエストからフォームの値を取得する
	handValue := r.PostFormValue("hand")

	//PlayJankenを実行して、結果を取得する
	jankenService := service.NewJankenService()
	result, err := jankenService.PlayJanken(model.Hand(handValue))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//結果をクライアントに返す
	json.NewEncoder(w).Encode(result)
}
