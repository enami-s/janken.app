// パッケージの定義
package repository

import "github.com/enami-s/janken_app/domain/model"

// DBアクセスメソッドのインタフェース定義
type JankenRepository interface {
	//じゃんけんの結果を保存するSave関数
	Save(model.JankenResponse) error
	//じゃんけんの結果一覧を取得するGetAll関数
	GetAll(int, int) ([]string, error)
}
