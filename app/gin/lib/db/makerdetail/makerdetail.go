package makerdetail

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	dbConfig "local.packages/game-information/config/database"
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

// GetList query [ select * from maker_detail where maker_id = ? ]
func GetList(makerIDs []int64) *[]Schema {

	config := dbConfig.GetMaster()
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	var strMakerIDs []string
	for _, v := range makerIDs {
		strMakerIDs = append(strMakerIDs, strconv.FormatInt(v, 10))
	}

	rows, err := db.Query("select * from maker_detail where maker_id IN (?)", strings.Join(strMakerIDs, ","))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var details []Schema
	var detail Schema
	for rows.Next() {
		detail = Schema{}
		err := rows.Scan(&detail.MakerID, &detail.OHP, &detail.TwitterName, &detail.YoutubeChannelID, &detail.YoutubeKeywords, &detail.CreatedAt, &detail.UpdatedAt)
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return &details
}
