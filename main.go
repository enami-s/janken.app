package main

import (
	"fmt"
	"github.com/enami-s/janken_app/presentation/handler"
	"net/http"
)

// じゃんけんのサーバーを定義する関数
func janken(w http.ResponseWriter, r *http.Request) {
	//じゃんけんアプリっていう出力を出す
	fmt.Fprintf(w, "じゃんけんアプリ\n")

	switch r.URL.Path {
	case "/rps":
		if r.Method == http.MethodPost {
			// PresentationのhandlerのPlayJankenHandler関数を呼び出す
			handler.PlayJankenHandler(w, r)
		}
	case "/results":
		if r.Method == http.MethodGet {
			// PresentationのhandlerのResultListHandler関数を呼び出す
			handler.ResultListHandler(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/rps", janken)
	http.HandleFunc("/results", janken)

	//http.ListenAndServeを用いてサーバーを起動する
	http.ListenAndServe(":8080", nil)

}
