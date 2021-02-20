package maker

import (
	"database/sql"
	"time"

	dbInstance "local.packages/game-information/lib/db/master"
	log "local.packages/game-information/lib/domain/log"
)

// Schema table schema [ maker ]
type Schema struct {
	ID        int64
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Maker インスタンス
type Maker struct {
	DBInstance *sql.DB
}

// GetInstance インスタンス生成
func GetInstance() *Maker {
	instance := dbInstance.GetInstance()
	return &Maker{DBInstance: instance.DB}
}

// Get query [ select * from maker where code = ? ]
func (db Maker) Get(code string) *Schema {

	maker := &Schema{}

	rows, err := db.DBInstance.Query("select * from maker where code = ?", code)
	if err != nil {
		log.Error(err.Error())
		return &Schema{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			log.Error(err.Error())
			return &Schema{}
		}
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error())
		return &Schema{}
	}

	return maker
}

// GetList query [ select * from maker order by name ]
func (db Maker) GetList() *[]Schema {

	rows, err := db.DBInstance.Query("select * from maker order by name")
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}
	defer rows.Close()

	var makers []Schema
	var maker Schema
	for rows.Next() {
		maker = Schema{}
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			log.Error(err.Error())
			return &[]Schema{}
		}
		makers = append(makers, maker)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}

	return &makers
}
