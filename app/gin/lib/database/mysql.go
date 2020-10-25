package lib/database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kantaroso/game-information/config"
)

func query(sql string, args interface{}, result *[]interface{}) bool {

	config := config.GetDatabase()

	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int
	var name string
	rows, err := db.Query(sql, 6)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
