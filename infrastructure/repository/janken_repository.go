// パッケージ定義
package repository

import (
	"database/sql"
	"github.com/enami-s/janken_app/domain/model"
	// パッケージのインポート
	_ "github.com/mattn/go-sqlite3"
)

// じゃんけんのインタフェースを構造体に実装
type JankenRepository struct {
}

// じゃんけんのインタフェースを実装した構造体を生成するNew関数
func NewJankenRepository() *JankenRepository {
	return &JankenRepository{}
}

// じゃんけんの結果を保存するSave関数
func (jr *JankenRepository) Save(result model.JankenResponse) error {
	db, err := LoadDB()
	if err != nil {
		return err
	}
	//DBを閉じる
	defer db.Close()
	_, err = db.Exec(
		"INSERT INTO janken_results(user_hand, computer_hand, result) VALUES(?, ?, ?)",
		result.UserHand,
		result.ComputerHand,
		result.Result,
	)
	if err != nil {
		return err
	}
	return nil
}

func LoadDB() (*sql.DB, error) {
	//SQLiteを用いてDBを作成する
	db, err := sql.Open("sqlite3", "./janken.db")

	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS janken_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_hand TEXT CHECK(user_hand IN ('rock', 'paper', 'scissors')) NOT NULL,
		computer_hand TEXT CHECK(user_hand IN ('rock', 'paper', 'scissors')) NOT NULL,
		result TEXT CHECK(result IN ('win', 'lose', 'draw')) NOT NULL
	);
	`
	if _, err = db.Exec(createTableSQL); err != nil {
		return nil, err
	}
	return db, err

}

// じゃんけんの結果一覧を取得するGetAll関数
func (jr *JankenRepository) GetAll() ([]string, error) {

	return nil, nil
}
