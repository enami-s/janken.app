package main

import (
	"fmt"
	"net/http"
)

// じゃんけんのサーバーを定義する関数
func janken(w http.ResponseWriter, r *http.Request) {
	//じゃんけんアプリっていう出力を出す
	fmt.Fprintf(w, "じゃんけんアプリ")
}

func main() {
	//http.HandleFuncを用いて、janken関数をルートパスに登録する
	http.HandleFunc("/", janken)
	//http.ListenAndServeを用いてサーバーを起動する
	http.ListenAndServe(":8080", nil)

}
