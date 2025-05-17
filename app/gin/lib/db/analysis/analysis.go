package analysis

import (
	"database/sql"
	"fmt"
	"sync"

	dbConfig "game-information/config/database"
	log "game-information/lib/domain/log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database Analysis用 インスタンス
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
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Error(err.Error(), {})
		return nil
	}
	return &Database{DB: db}
}
