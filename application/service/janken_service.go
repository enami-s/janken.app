// パッケージのインポート
package service

import (
	"github.com/enami-s/janken_app/domain/model"
	"github.com/enami-s/janken_app/domain/service"
	"math/rand"
	"time"
)

func getRandomHand() model.Hand {
	hands := []model.Hand{model.Rock, model.Paper, model.Scissors}
	rand.Seed(time.Now().UnixNano())
	return hands[rand.Intn(len(hands))]
}

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

func getJankenResult(userHand, computerHand model.Hand) model.Result {
	// 引き分けの場合
	if userHand == computerHand {
		return model.Draw
	}

	// ユーザーの勝ちの場合
	if (userHand == model.Rock && computerHand == model.Scissors) ||
		(userHand == model.Paper && computerHand == model.Rock) ||
		(userHand == model.Scissors && computerHand == model.Paper) {
		return model.Win
	}

	// ユーザーの負けの場合
	return model.Lose
}

// PlayJankenの実装
func (s *jankenServiceimpl) PlayJanken(hand model.Hand) (model.JankenResult, error) {
	//コンピュターのじゃんけんの手をランダムで取得する
	computerHand := getRandomHand()
	result := getJankenResult(hand, computerHand)

	//JankenResultの構造体を返す
	return model.JankenResult{
		ComputerHand: computerHand,
		UserHand:     hand,
		Result:       result,
	}, nil
}
