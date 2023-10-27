// パッケージのインポート
package service

import "github.com/enami-s/janken_app/domain/service"

// JankenServiceインタフェースの実体となる構造体を実装
type jankenServiceimpl struct{}

// NewJankenService関数の実装
func NewJankenService() service.JankenService {
	return &jankenServiceimpl{}
}

// SaveResultの実装
func (s *jankenServiceimpl) SaveResult(result string) error {
	//正常な返り値を返す
	return nil
}

// PlayJankenの実装
func (s *jankenServiceimpl) PlayJanken(hand string) (string, error) {
	//正常な返り値を返す
	return "", nil
}
