// パッケージのインポート
package janken

import "github.com/enami-s/janken.app/domain/service"

// JankenServiceインタフェースの実体となる構造体を実装
type jankenServiceimpl struct{}

// NewJankenService関数の実装
func NewJankenService() service.JankenService {
	return &jankenServiceimpl{}
}

// PlayJankenの実装
func (s *jankenServiceimpl) PlayJanken(hand string) (string, error) {

}

// SaveResultの実装
func (s *jankenServiceimpl) SaveResult(result string) error {

}
