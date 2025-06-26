package makerdetail

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	dbInstance "game-information/lib/db/master"
	log "game-information/lib/domain/log"
)

// Schema table schema [ maker_detail ]
type Schema struct {
	MakerID          int64
	OHP              string
	TwitterName      string
	YoutubeChannelID string
	YoutubeKeywords  string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// MakerDetail インスタンス
type MakerDetail struct {
	DBInstance *sql.DB
}

// GetInstance インスタンス生成
func GetInstance() *MakerDetail {
	instance := dbInstance.GetInstance()
	return &MakerDetail{DBInstance: instance.DB}
}

// GetList query [ select * from master.maker_detail where maker_id = ? ]
func (db *MakerDetail) GetList(makerIDs []int64) *[]Schema {

	var strMakerIDs []string
	for _, v := range makerIDs {
		strMakerIDs = append(strMakerIDs, strconv.FormatInt(v, 10))
	}
	if len(strMakerIDs) == 0 {
		return &[]Schema{}
	}

	rows, err := db.DBInstance.Query(db.CreateSelectQuery(strings.Join(strMakerIDs, ",")))
	if err != nil {
		log.Error(err.Error(), nil)
		return &[]Schema{}
	}
	defer rows.Close()

	var details []Schema
	var detail Schema
	for rows.Next() {
		detail = Schema{}
		err := rows.Scan(&detail.MakerID, &detail.OHP, &detail.TwitterName, &detail.YoutubeChannelID, &detail.YoutubeKeywords, &detail.CreatedAt, &detail.UpdatedAt)
		if err != nil {
			log.Error(err.Error(), nil)
			return &[]Schema{}
		}
		details = append(details, detail)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error(), nil)
		return &[]Schema{}
	}

	return &details
}

// Insert [insert into master.maker_detail(maker_id, ohp_url, twitter_name, youtube_channel_id, youtube_keywords, created_at, updated_at) values (?,?,?,?,?)]
func (db *MakerDetail) Insert(schemas *[]Schema) bool {
	_, err := db.DBInstance.Exec(db.CreateBulkInsertQuery(schemas))
	if err != nil {
		log.Error(err.Error(), nil)
		return false
	}
	return true
}

// Update [update master.maker_detail set ohp_url=$1, twitter_name=$2, youtube_channel_id=$3, youtube_keywords=$4, updated_at=NOW() where maker_id=$5]
func (db *MakerDetail) Update(schema *Schema) bool {
	_, err := db.DBInstance.Exec("update master.maker_detail set ohp_url=$1, twitter_name=$2, youtube_channel_id=$3, youtube_keywords=$4, updated_at=NOW() where maker_id=$5", schema.OHP, schema.TwitterName, schema.YoutubeChannelID, schema.YoutubeKeywords, schema.MakerID)
	if err != nil {
		log.Error(err.Error(), nil)
		return false
	}
	return true
}

func (db *MakerDetail) CreateBulkInsertQuery(schemas *[]Schema) string {
	baseSQLStr := "insert into master.maker_detail(maker_id, ohp_url, twitter_name, youtube_channel_id, youtube_keywords, created_at, updated_at) values %s"
	valueSQLStr := "(%d, '%s', '%s', '%s', '%s', NOW(), NOW())"
	var valueSQLArray []string
	for _, item := range *schemas {
		valueSQLArray = append(valueSQLArray, fmt.Sprintf(valueSQLStr, item.MakerID, item.OHP, item.TwitterName, item.YoutubeChannelID, item.YoutubeKeywords))
	}
	return fmt.Sprintf(baseSQLStr, strings.Join(valueSQLArray, ","))
}

func (db *MakerDetail) CreateSelectQuery(IDs string) string {
	baseSQLStr := "select * from master.maker_detail where maker_id IN (%s)"
	return fmt.Sprintf(baseSQLStr, IDs)
}
