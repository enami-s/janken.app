// パッケージ定義
package repository

import (
	"database/sql"
	"github.com/enami-s/janken_app/domain/model"
)

// じゃんけんのインタフェースを構造体に実装
type JankenRepository struct {
}

// じゃんけんのインタフェースを実装した構造体を生成するNew関数
func NewJankenRepository() *JankenRepository {
	return &JankenRepository{}
}

// じゃんけんの結果を保存するSave関数
func (jr *JankenRepository) Save(result *model.Result) error {
	//SQLiteを用いてDBを作成する
	db, err := sql.Open("sqlite3", "./janken.db")
	if err != nil {
		return err
	}

	//DBにじゃんけんの結果を保存する

	//DBを閉じる
	defer db.Close()

	return nil
}

// じゃんけんの結果一覧を取得するGetAll関数
func (jr *JankenRepository) GetAll() ([]*model.Result, error) {
	return nil, nil
}
