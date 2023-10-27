package main

import (
	"github.com/enami-s/janken_app/presentation/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/rps", handler.PlayJankenHandler)
	http.HandleFunc("/results", handler.ResultListHandler)

	//http.ListenAndServeを用いてサーバーを起動する
	http.ListenAndServe(":8080", nil)
}
