package makervideo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	dbInstance "local.packages/game-information/lib/db/master"
	log "local.packages/game-information/lib/domain/log"
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

// MakerVideo インスタンス
type MakerVideo struct {
	DBInstance *sql.DB
}

// GetInstance インスタンス生成
func GetInstance() *MakerVideo {
	instance := dbInstance.GetInstance()
	return &MakerVideo{DBInstance: instance.DB}
}

// GetList query [ select * from maker_video where maker_id = ? order by id DESC ]
func (db *MakerVideo) GetList(makerID int64) *[]Schema {

	rows, err := db.DBInstance.Query("select * from maker_video where maker_id = ? order by id DESC", makerID)
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}
	defer rows.Close()

	var videos []Schema
	var video Schema
	for rows.Next() {
		video = Schema{}
		err := rows.Scan(&video.ID, &video.MakerID, &video.VideoID, &video.Title, &video.PublishedAt, &video.CreatedAt)
		if err != nil {
			log.Error(err.Error())
			return &[]Schema{}
		}
		videos = append(videos, video)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err.Error())
		return &[]Schema{}
	}

	return &videos
}

// BulkInsert [ insert int maker_video(maker_id, video_id, title, published_at) value ....]
func (db *MakerVideo) BulkInsert(makerID int64, videos *[]domainYoutube.Video) bool {
	_, err := db.DBInstance.Query(db.CreateBulkInsertQuery(makerID, videos))
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

func (db *MakerVideo) CreateBulkInsertQuery(makerID int64, videos *[]domainYoutube.Video) string {
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
