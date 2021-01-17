package makervideo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	dbConfig "local.packages/game-information/config/database"
	domainYoutube "local.packages/game-information/lib/domain/youtube"
)

// Schema table schema [ maker_video ]
type Schema struct {
	ID          int64
	MakerID     int64
	VideoID     string
	Title       string
	PublishedAt time.Time
	CreatedAt   time.Time
}

// GetList query [ select * from maker_video where maker_id = ? order by id DESC ]
func GetList(makerID int64) *[]Schema {

	config := dbConfig.GetMaster()
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from maker_video where maker_id = ? order by id DESC", makerID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var videos []Schema
	var video Schema
	for rows.Next() {
		video = Schema{}
		err := rows.Scan(&video.ID, &video.MakerID, &video.VideoID, &video.Title, &video.PublishedAt, &video.CreatedAt)
		if err != nil {
			panic(err)
		}
		videos = append(videos, video)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return &videos
}

// BulkInsert [ insert int maker_video(maker_id, video_id, title, published_at) value ....]
func BulkInsert(makerID int64, videos *[]domainYoutube.Video) bool {

	config := dbConfig.GetMaster()
	connection := fmt.Sprintf(dbConfig.ConnectionOption, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	_, err = db.Query(createBulkInsertQuery(makerID, videos))
	if err != nil {
		panic(err)
	}
	return true
}

func createBulkInsertQuery(makerID int64, videos *[]domainYoutube.Video) string {
	baseSQLStr := "insert into maker_video (maker_id, video_id, title, published_at, created_at) values %s"
	valueSQLStr := "(%d, '%s', '%s', '%s', '%s')"
	var valueSQLArray []string
	jst, _ := time.LoadLocation("Asia/Tokyo")
	createdAt := time.Now().In(jst).Format("2006-01-02 15:04:05")
	for _, item := range *videos {
		valueSQLArray = append(valueSQLArray, fmt.Sprintf(valueSQLStr, makerID, item.ID, item.Title, item.PublishedAt.In(jst).Format("2006-01-02 15:04:05"), createdAt))
	}
	return fmt.Sprintf(baseSQLStr, strings.Join(valueSQLArray, ","))
}
