package master

import (
	"database/sql"
	"fmt"
	"sync"

	dbConfig "game-information/config/database"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
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
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	return &Database{DB: db}
}
