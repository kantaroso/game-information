package maker

import (
	"sort"
	"time"

	dbMaker "local.packages/game-information/lib/db/master/maker"
	dbMakerdetail "local.packages/game-information/lib/db/master/makerdetail"
	dbMakervideo "local.packages/game-information/lib/db/master/makervideo"
	domainYoutube "local.packages/game-information/lib/domain/youtube"
)

// Maker インスタンス
type Maker struct {
	DBMaker       *dbMaker.Maker
	DBMakerdetail *dbMakerdetail.MakerDetail
	DBMakervideo  *dbMakervideo.MakerVideo
}

// GetInstance インスタンス生成
func GetInstance() *Maker {
	return &Maker{DBMaker: dbMaker.GetInstance(), DBMakerdetail: dbMakerdetail.GetInstance(), DBMakervideo: dbMakervideo.GetInstance()}
}

// GetMaker メーカー情報取得
func (domain Maker) GetMaker(code string) *dbMaker.Schema {
	return domain.DBMaker.Get(code)
}

// GetMakerList メーカー情報一覧取得
func (domain Maker) GetMakerList() *[]dbMaker.Schema {
	return domain.DBMaker.GetList()
}

// GetDetail メーカー情報詳細取得
func (domain Maker) GetDetail(makerID int64) *dbMakerdetail.Schema {
	mkaerIDs := []int64{makerID}
	details := domain.GetDetailList(mkaerIDs)
	if len(*details) == 0 {
		return &dbMakerdetail.Schema{}
	}
	return &(*details)[0]
}

// GetDetailList メーカー情報詳細一覧取得
func (domain Maker) GetDetailList(mkaerIDs []int64) *[]dbMakerdetail.Schema {
	return domain.DBMakerdetail.GetList(mkaerIDs)
}

// GetVideoList メーカー動画情報取得
func (domain Maker) GetVideoList(mkaerID int64) *[]dbMakervideo.Schema {
	return domain.DBMakervideo.GetList(mkaerID)
}

// UpdateVideoList 動画情報の更新
func (domain Maker) UpdateVideoList(makerID int64) bool {

	detail := domain.GetDetail(makerID)
	if detail.YoutubeChannelID == "" {
		return true
	}

	makervideos := domain.GetVideoList(makerID)
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
	domain.DBMakervideo.BulkInsert(makerID, videos)
	return true
}
