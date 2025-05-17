package maker

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	dbInstance "game-information/lib/db/master"
	log "game-information/lib/domain/log"
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
func (db *Maker) Get(code string) *Schema {

	maker := &Schema{}

	rows, err := db.DBInstance.Query("select * from maker where code = ?", code)
	if err != nil {
		log.Error(err.Error(), {})
		return &Schema{}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			log.Error(err.Error(), {})
			return &Schema{}
		}
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error(), {})
		return &Schema{}
	}

	return maker
}

// GetList query [ select * from maker order by id ]
func (db *Maker) GetList() *[]Schema {

	rows, err := db.DBInstance.Query("select * from maker order by id")
	if err != nil {
		log.Error(err.Error(), {})
		return &[]Schema{}
	}
	defer rows.Close()

	var makers []Schema
	var maker Schema
	for rows.Next() {
		maker = Schema{}
		err := rows.Scan(&maker.ID, &maker.Name, &maker.Code, &maker.CreatedAt, &maker.UpdatedAt)
		if err != nil {
			log.Error(err.Error(), {})
			return &[]Schema{}
		}
		makers = append(makers, maker)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error(), {})
		return &[]Schema{}
	}

	return &makers
}

// Insert [insert into maker(id, name, code, created_at, updated_at) values (?,?,?,?,?)]
func (db *Maker) Insert(schemas *[]Schema) bool {
	_, err := db.DBInstance.Exec(db.CreateBulkInsertQuery(schemas))
	if err != nil {
		log.Error(err.Error(), {})
		return false
	}
	return true
}

// Update [update maker set name=?, code=?, updated_at=? where id=?]
func (db *Maker) Update(schema *Schema) bool {
	_, err := db.DBInstance.Exec("update maker set name=?, code=?, updated_at=NOW() where id=?", schema.Name, schema.Code, schema.ID)
	if err != nil {
		log.Error(err.Error(), {})
		return false
	}
	return true
}

func (db *Maker) CreateBulkInsertQuery(schemas *[]Schema) string {
	baseSQLStr := "insert into maker(id, name, code, created_at, updated_at) values %s"
	valueSQLStr := "(%d, '%s', '%s', NOW(), NOW())"
	var valueSQLArray []string
	for _, item := range *schemas {
		valueSQLArray = append(valueSQLArray, fmt.Sprintf(valueSQLStr, item.ID, item.Name, item.Code))
	}
	return fmt.Sprintf(baseSQLStr, strings.Join(valueSQLArray, ","))
}
