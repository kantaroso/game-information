package maker

import (
	"sort"
	"time"

	dbMaker "local.packages/game-information/lib/db/maker"
	dbMakerdetail "local.packages/game-information/lib/db/makerdetail"
	dbMakervideo "local.packages/game-information/lib/db/makervideo"
	domainYoutube "local.packages/game-information/lib/domain/youtube"
)

// GetMaker メーカー情報取得
func GetMaker(code string) *dbMaker.Schema {
	return dbMaker.Get(code)
}

// GetMakerList メーカー情報一覧取得
func GetMakerList() *[]dbMaker.Schema {
	return dbMaker.GetList()
}

// GetDetail メーカー情報詳細取得
func GetDetail(makerID int64) *dbMakerdetail.Schema {
	mkaerIDs := []int64{makerID}
	details := GetDetailList(mkaerIDs)
	if len(*details) == 0 {
		return &dbMakerdetail.Schema{}
	}
	return &(*details)[0]
}

// GetDetailList メーカー情報詳細一覧取得
func GetDetailList(mkaerIDs []int64) *[]dbMakerdetail.Schema {
	return dbMakerdetail.GetList(mkaerIDs)
}

// GetVideoList メーカー動画情報取得
func GetVideoList(mkaerID int64) *[]dbMakervideo.Schema {
	return dbMakervideo.GetList(mkaerID)
}

// UpdateVideoList 動画情報の更新
func UpdateVideoList(makerID int64) bool {

	detail := GetDetail(makerID)
	if detail.YoutubeChannelID == "" {
		return true
	}

	makervideos := GetVideoList(makerID)
	latestat := ""
	if len(*makervideos) != 0 {
		// 形式：1970-01-01T00:00:00Z
		// 保存されている最新の動画よりも新しいものだけ取得
		// 指定した時刻を含むデータが返却されるので1s足す
		latestat = (*makervideos)[0].PublishedAt.Add(1 * time.Second).Format(time.RFC3339)
	}

	videos := domainYoutube.GetVideos(detail.YoutubeChannelID, detail.YoutubeKeywords, latestat, "")
	if len(*videos) == 0 {
		return true
	}
	// 新しいものがid大きくなるようにするため並べ替え
	sort.Slice(*videos, func(i, j int) bool { return (*videos)[i].PublishedAt.Unix() < (*videos)[j].PublishedAt.Unix() })

	// データinsert
	dbMakervideo.BulkInsert(makerID, videos)
	return true
}
