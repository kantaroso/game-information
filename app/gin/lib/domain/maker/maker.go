package maker

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	dbMaker "game-information/lib/db/master/maker"
	dbMakerdetail "game-information/lib/db/master/makerdetail"
	dbMakervideo "game-information/lib/db/master/makervideo"
	log "game-information/lib/domain/log"
	domainSpreadsheet "game-information/lib/domain/spreadsheet"
	domainYoutube "game-information/lib/domain/youtube"
)

// Maker インスタンス
type Maker struct {
	DBMaker           *dbMaker.Maker
	DBMakerdetail     *dbMakerdetail.MakerDetail
	DBMakervideo      *dbMakervideo.MakerVideo
	DomainYoutube     domainYoutube.InterfaceYoutube
	DomainSpreadsheet domainSpreadsheet.InterfaceSpreadsheet
}

const spreadsheetName = "maker"
const sheetColumnID = 0
const sheetColumnName = 1
const sheetColumnCode = 2
const sheetColumnOHP = 3
const sheetColumnTwitterName = 4
const sheetColumnYoutubeChannelID = 5
const sheetColumnYoutubeKeywords = 6

// spreadsheetからデータを取得するためのrange文字列の生成
func getSpreadsheetRange() string {
	endAlphabet := domainSpreadsheet.ConvertColNumberToAlphabet(sheetColumnYoutubeKeywords + 1)
	return fmt.Sprintf("%s!A2:%s1000", spreadsheetName, endAlphabet)
}

// spreadsheetのデータからMakerの構造体を作成する
func createMakerSchemaFromSpreadsheet(makerID int64, row []interface{}) dbMaker.Schema {
	return dbMaker.Schema{ID: makerID, Name: row[sheetColumnName].(string), Code: row[sheetColumnCode].(string)}
}

// spreadsheetのデータからMakerDetailの構造体を作成する
func createDetailSchemaFromSpreadsheet(makerID int64, row []interface{}) dbMakerdetail.Schema {
	return dbMakerdetail.Schema{MakerID: makerID, OHP: row[sheetColumnOHP].(string), TwitterName: row[sheetColumnTwitterName].(string), YoutubeChannelID: row[sheetColumnYoutubeChannelID].(string), YoutubeKeywords: row[sheetColumnYoutubeKeywords].(string)}
}

// Makerのspreadsheetのデータと差分があるか
func isDiffMaker(schema *dbMaker.Schema, row []interface{}) bool {
	if schema.Name != row[sheetColumnName].(string) {
		return true
	}
	if schema.Code != row[sheetColumnCode].(string) {
		return true
	}
	return false
}

// MakerDetailのspreadsheetのデータと差分があるか
func isDiffDetail(schema *dbMakerdetail.Schema, row []interface{}) bool {
	if schema.OHP != row[sheetColumnOHP].(string) {
		return true
	}
	if schema.TwitterName != row[sheetColumnTwitterName].(string) {
		return true
	}
	if schema.YoutubeChannelID != row[sheetColumnYoutubeChannelID].(string) {
		return true
	}
	if schema.YoutubeKeywords != row[sheetColumnYoutubeKeywords].(string) {
		return true
	}
	return false
}

// GetInstance インスタンス生成
func GetInstance() *Maker {
	return &Maker{DBMaker: dbMaker.GetInstance(), DBMakerdetail: dbMakerdetail.GetInstance(), DBMakervideo: dbMakervideo.GetInstance(), DomainYoutube: domainYoutube.GetInstance(), DomainSpreadsheet: domainSpreadsheet.GetInstance()}
}

// GetMaker メーカー情報取得
func (domain *Maker) GetMaker(code string) *dbMaker.Schema {
	return domain.DBMaker.Get(code)
}

// GetMakerList メーカー情報一覧取得
func (domain *Maker) GetMakerList() *[]dbMaker.Schema {
	return domain.DBMaker.GetList()
}

// GetDetail メーカー情報詳細取得
func (domain *Maker) GetDetail(makerID int64) *dbMakerdetail.Schema {
	mkaerIDs := []int64{makerID}
	details := domain.GetDetailList(mkaerIDs)
	if len(*details) == 0 {
		return &dbMakerdetail.Schema{}
	}
	return &(*details)[0]
}

// GetDetailList メーカー情報詳細一覧取得
func (domain *Maker) GetDetailList(makerIDs []int64) *[]dbMakerdetail.Schema {
	return domain.DBMakerdetail.GetList(makerIDs)
}

// GetVideoList メーカー動画情報取得
func (domain *Maker) GetVideoList(makerID int64) *[]dbMakervideo.Schema {
	return domain.DBMakervideo.GetList(makerID)
}

// UpdateMaker Maker,MakerDetailの更新
func (domain *Maker) UpdateMakerList() bool {
	res := domain.DomainSpreadsheet.GetSheetData(getSpreadsheetRange())
	if res == nil {
		return false
	}

	var makerIDList []int64
	makerMapList := map[int64]dbMaker.Schema{}
	makerList := domain.GetMakerList()
	for _, tmpMaker := range *makerList {
		makerIDList = append(makerIDList, tmpMaker.ID)
		makerMapList[tmpMaker.ID] = tmpMaker
	}
	detailMapList := map[int64]dbMakerdetail.Schema{}
	detailList := domain.GetDetailList(makerIDList)
	for _, tmpDetail := range *detailList {
		detailMapList[tmpDetail.MakerID] = tmpDetail
	}

	var insertMakerList []dbMaker.Schema
	var insertDetailList []dbMakerdetail.Schema
	dbResult := true
	for _, row := range res {
		if row[sheetColumnID] == nil {
			continue
		}
		makerID, makerIDError := strconv.ParseInt(row[sheetColumnID].(string), 10, 64)
		if makerIDError != nil {
			continue
		}
		makerName := row[sheetColumnName].(string)
		tmpMaker, existMaker := makerMapList[makerID]
		tmpDetail, _ := detailMapList[makerID]
		if !existMaker {
			insertMakerList = append(insertMakerList, createMakerSchemaFromSpreadsheet(makerID, row))
			insertDetailList = append(insertDetailList, createDetailSchemaFromSpreadsheet(makerID, row))
			log.Info(fmt.Sprintf("insert target maker id:%d name:%s", makerID, makerName), nil)
			continue
		}
		if isDiffMaker(&tmpMaker, row) {
			log.Info(fmt.Sprintf("update target maker id:%d name:%s", makerID, makerName), nil)
			targetMaker := createMakerSchemaFromSpreadsheet(makerID, row)
			if !domain.DBMaker.Update(&targetMaker) {
				dbResult = false
				log.Error(fmt.Sprintf("update master.maker error id:%d name:%s", makerID, makerName), nil)
			}
		}
		if isDiffDetail(&tmpDetail, row) {
			log.Info(fmt.Sprintf("update target maker_detail id:%d name:%s", makerID, makerName), nil)
			targetDetail := createDetailSchemaFromSpreadsheet(makerID, row)
			if !domain.DBMakerdetail.Update(&targetDetail) {
				dbResult = false
				log.Error(fmt.Sprintf("update master.maker_detail error id:%d name:%s", makerID, makerName), nil)
			}
		}
	}
	if len(insertMakerList) > 0 {
		if !domain.DBMaker.Insert(&insertMakerList) {
			dbResult = false
			log.Error("insert maker error", insertMakerList)
		}
		if !domain.DBMakerdetail.Insert(&insertDetailList) {
			dbResult = false
			log.Error("insert maker_detail error", insertDetailList)
		}
	}
	return dbResult
}

// UpdateVideoList 動画情報の更新
func (domain *Maker) UpdateVideoList(makerID int64) bool {

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

	videos := domain.DomainYoutube.GetVideos(detail.YoutubeChannelID, detail.YoutubeKeywords, latestat, "")
	if len(*videos) == 0 {
		return true
	}
	// 新しいものがid大きくなるようにするため並べ替え
	sort.Slice(*videos, func(i, j int) bool { return (*videos)[i].PublishedAt.Unix() < (*videos)[j].PublishedAt.Unix() })

	// データinsert
	return domain.DBMakervideo.BulkInsert(makerID, videos)
}
