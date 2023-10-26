// パッケージの定義
package repository

import "your_module_name/domain/model"

// DBアクセスメソッドのインタフェース定義
type JankenRepository interface {
	//じゃんけんの結果を保存するSave関数
	Save(result *model.Result) error
	//じゃんけんの結果一覧を取得するGetAll関数
	GetAll() ([]*model.Result, error)
}
