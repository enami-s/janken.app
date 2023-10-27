// パッケージの定義
package repository

// DBアクセスメソッドのインタフェース定義
type JankenRepository interface {
	//じゃんけんの結果を保存するSave関数
	Save(janken_data string) error
	//じゃんけんの結果一覧を取得するGetAll関数
	GetAll() ([]string, error)
}
