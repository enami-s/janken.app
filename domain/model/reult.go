//じゃんけんの結果モデルの定義

package model

// じゃんけんの手をEnumで定義(rock, paper, scissors)
type Hand string

const (
	Rock     = Hand("rock")
	Paper    = Hand("paper")
	Scissors = Hand("scissors")
)

// じゃんけんの結果をEnumで定義(win, lose, draw)
type Result string

const (
	Win  = Result("win")
	Lose = Result("lose")
	Draw = Result("draw")
)

// じゃんけんの結果モデルの定義(Computerの手とUserの手と結果)
type JankenResult struct {
	ComputerHand Hand
	UserHand     Hand
	Result       Result
}
