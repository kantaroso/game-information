package analysis

import (
	"database/sql"
	"fmt"
	"sync"

	dbConfig "game-information/config/database"
	log "game-information/lib/domain/log"

	// MySQLドライバーの代わりにPostgreSQLドライバーをインポート
	_ "github.com/lib/pq" // PostgreSQLドライバー
)

// Database analysis用 インスタンス
type Database struct {
	DB *sql.DB
}

var instance *Database
var once sync.Once

// GetInstance returns singleton instance
func GetInstance() *Database {
	once.Do(func() {
		instance = getInstance()
	})
	return instance
}

func getInstance() *Database {
	config := dbConfig.GetAnalysis()

	// PostgreSQL用の接続文字列フォーマットに変更
	// search_path を追加し、デフォルトで参照するスキーマを指定
	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s", // search_path を追加
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
		config.Schema,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Error(err.Error(), nil)
	}

	err = db.Ping()
	if err != nil {
		log.Error(err.Error(), nil)
	}

	// supabaseのmasterスキーマを参照するためにsearch_pathを設定
	// 接続プーラーを利用するのでそのあたりが関係してるかも
	_, err = db.Exec(fmt.Sprintf("SET search_path TO %s", config.Schema))
	if err != nil {
		log.Error(fmt.Sprintf("Failed to set search_path to '%s': %s", config.Schema, err.Error()), nil)
	}

	return &Database{DB: db}
}