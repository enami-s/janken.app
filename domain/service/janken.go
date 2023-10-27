// PlayJankenとSaveResultのインタフェースの定義など
package service

// JankenServiceインタフェースの定義
type JankenService interface {
	// PlayJanken関数の定義
	PlayJanken(hand string) (string, error)
	// SaveResult関数の定義
	SaveResult(result string) error
}
