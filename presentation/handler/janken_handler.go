package handler

import (
	"encoding/json"
	"github.com/enami-s/janken_app/application/service"
	"net/http"
)

// PlayJanlenHandler関数の実装
func PlayJankenHandler(w http.ResponseWriter, r *http.Request) {
	//POSTリクエスト以外の場合は、処理を中断する
	if r.Method != http.MethodPost {
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
