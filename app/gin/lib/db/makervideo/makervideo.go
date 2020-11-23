package makervideo

import (
	"database/sql"
	"fmt"
	"time"

	dbConfig "local.packages/game-information/config/database"
)

// Schema table schema [ maker_video ]
type Schema struct {
	ID        int64
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get query [ select * from maker_video where code = ? ]
func Get(code string) *Schema {

	dbConfig := dbConfig.GetMaster()
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

	video := &Schema{}
	for rows.Next() {
		err := rows.Scan(&video.ID, &video.Name, &video.Code, &video.CreatedAt, &video.UpdatedAt)
		if err != nil {
			panic(err)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return video
}
