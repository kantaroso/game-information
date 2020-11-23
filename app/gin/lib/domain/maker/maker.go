package maker

import (
	dbMaker "local.packages/game-information/lib/db/maker"
	dbMakerdetail "local.packages/game-information/lib/db/makerdetail"
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
