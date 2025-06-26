package accesslog

import (
	"database/sql"

	dbInstance "game-information/lib/db/analysis"
	log "game-information/lib/domain/log"
)

// Accesslog インスタンス
type Accesslog struct {
	DBInstance *sql.DB
}

// GetInstance インスタンス生成
func GetInstance() *Accesslog {
	instance := dbInstance.GetInstance()
	return &Accesslog{DBInstance: instance.DB}
}

// CountAll query [ select count(id) from access_log ]
func (db *Accesslog) CountAll() int {
	var count int
	rows, err := db.DBInstance.Query("select count(id) from access_log")
	if err != nil {
		log.Error(err.Error(), nil)
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Error(err.Error(), nil)
			return 0
		}
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error(), nil)
		return 0
	}

	return count
}

// Insert query [ insert into access_log(method,endpoint,query_string,user_agent) values($1,$2,$3,$4) ]
func (db *Accesslog) Insert(method string, endpoint string, queryString string, userAgent string) bool {
	_, err := db.DBInstance.Exec("insert into access_log(method,endpoint,query_string,user_agent) values($1,$2,$3,$4)", method, endpoint, queryString, userAgent)
	if err != nil {
		log.Error(err.Error(), nil)
		return false
	}
	return true
}
