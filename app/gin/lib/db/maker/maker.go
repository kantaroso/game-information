package maker

import (
	"database/sql"
	"fmt"
	"time"

	dbConfig "local.packages/game-information/config/database"
)

// Schema table schema [ maker ]
type Schema struct {
	ID        int64
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get query [ select * from maker where code = ? ]
func Get(code string) *Schema {

	config := dbConfig.GetMaster()
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from maker where code = ?", code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	maker := &Schema{}
	for rows.Next() {
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return maker
}

// GetList query [ select * from maker order by name ]
func GetList() *[]Schema {

	config := dbConfig.GetMaster()
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from maker order by name")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var makers []Schema
	var maker Schema
	for rows.Next() {
		maker = Schema{}
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			panic(err)
		}
		makers = append(makers, maker)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return &makers
}
