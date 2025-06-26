package master

import (
	"database/sql"
	"fmt"
	"sync"

	dbConfig "game-information/config/database"
	log "game-information/lib/domain/log"

	_ "github.com/lib/pq"
)

// Database Master用 インスタンス
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
	config := dbConfig.GetMaster()

	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Error(err.Error(), nil)
	}

	err = db.Ping()
	if err != nil {
		log.Error(err.Error(), nil)
	}

	return &Database{DB: db}
}