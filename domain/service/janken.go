// PlayJankenとSaveResultのインタフェースの定義など
package service

// JankenServiceインタフェースの定義
type JankenService interface {
	// PlayJanken関数の定義
	PlayJanken(hand string) (string, error)
}
