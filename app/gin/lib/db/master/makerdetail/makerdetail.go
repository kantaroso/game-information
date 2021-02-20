package makerdetail

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	dbInstance "local.packages/game-information/lib/db/master"
	log "local.packages/game-information/lib/domain/log"
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

// GetList query [ select * from maker_detail where maker_id = ? ]
func (db MakerDetail) GetList(makerIDs []int64) *[]Schema {

	var strMakerIDs []string
	for _, v := range makerIDs {
		strMakerIDs = append(strMakerIDs, strconv.FormatInt(v, 10))
	}

	rows, err := db.DBInstance.Query("select * from maker_detail where maker_id IN (?)", strings.Join(strMakerIDs, ","))
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}
	defer rows.Close()

	var details []Schema
	var detail Schema
	for rows.Next() {
		detail = Schema{}
		err := rows.Scan(&detail.MakerID, &detail.OHP, &detail.TwitterName, &detail.YoutubeChannelID, &detail.YoutubeKeywords, &detail.CreatedAt, &detail.UpdatedAt)
		if err != nil {
			log.Error(err.Error())
			return &[]Schema{}
		}
		details = append(details, detail)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}

	return &details
}
