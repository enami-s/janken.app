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

// PlayJankenの実装
func (s *jankenServiceimpl) PlayJanken(hand model.Hand) (model.JankenResult, error) {
	//コンピュターのじゃんけんの手をランダムで取得する
	computerHand := getRandomHand()

	//勝敗をIf分で判定する
	// 引き分けの場合
	if hand == computerHand {
		return model.JankenResult{
			ComputerHand: computerHand,
			UserHand:     model.Hand(hand),
			Result:       model.Draw,
		}, nil
	}

	// ユーザーの勝ちの場合
	if (hand == model.Rock && computerHand == model.Scissors) ||
		(hand == model.Paper && computerHand == model.Rock) ||
		(hand == model.Scissors && computerHand == model.Paper) {
		return model.JankenResult{
			ComputerHand: computerHand,
			UserHand:     model.Hand(hand),
			Result:       model.Win,
		}, nil
	}

	// ユーザーの負けの場合
	return model.JankenResult{
		ComputerHand: computerHand,
		UserHand:     model.Hand(hand),
		Result:       model.Lose,
	}, nil
}
