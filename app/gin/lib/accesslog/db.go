package accesslog

import (
	"database/sql"
	"fmt"

	"github.com/kantaroso/game-information/config"
)

// CountAll アクセスログ数取得
func countAll() int {

	config := config.GetDatabase()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	var count int
	rows, err := db.Query("select count(id) from access_log")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return count
}

func insert(method string, endpoint string, queryString string, userAgent string) {

	config := config.GetDatabase()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("insert into access_log(method,endpoint,query_string,user_agent) values(?,?,?,?)", method, endpoint, queryString, userAgent)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
