// PlayJankenとSaveResultのインタフェースの定義など
package service

import "github.com/enami-s/janken_app/domain/model"

// JankenServiceインタフェースの定義
type JankenService interface {
	// PlayJanken関数の定義
	PlayJanken(hand model.Hand) (model.JankenResult, error)
}
