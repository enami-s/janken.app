package handler

import (
	"encoding/json"
	"github.com/enami-s/janken_app/application/service"
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

	//HandValueの値を出力
	json.NewEncoder(w).Encode(handValue)

	//PlayJankenの実行
	service.NewJankenService().PlayJanken(handValue)

	//SaveResultの実行
}
