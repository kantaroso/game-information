package maker

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	dbMaker "local.packages/game-information/lib/db/master/maker"
	dbMakerdetail "local.packages/game-information/lib/db/master/makerdetail"
	dbMakervideo "local.packages/game-information/lib/db/master/makervideo"
	domainYoutube "local.packages/game-information/lib/domain/youtube"
)

type DomainYoutubeTest struct {
	domainYoutube.Youtube
}

func (domain *DomainYoutubeTest) GetVideos(channelID string, query string, publishedAfter string, token string) *[]domainYoutube.Video {
	return getTestYoutubeData(channelID)
}

func TestGetMaker(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var result *dbMaker.Schema
	var rows *sqlmock.Rows
	queryStr := "select * from maker where code = ?"

	// カラムの要素
	var id int64 = 1
	name := "name_test"
	code := "code_test"
	tmpTime := time.Now()

	// 1件以上取れる場合
	rows = sqlmock.NewRows(getMakerColumns()).AddRow(id, name, code, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(code).WillReturnRows(rows)
	result = domainMakerInstnce.GetMaker(code)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, code, result.Code)
	assert.Equal(t, tmpTime, result.CreatedAt)
	assert.Equal(t, tmpTime, result.UpdatedAt)

	// データなし
	rows = sqlmock.NewRows(getMakerColumns())
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(code).WillReturnRows(rows)
	result = domainMakerInstnce.GetMaker(code)
	assert.Empty(t, result.ID)

	// エラー
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(code).WillReturnError(fmt.Errorf("some error case"))
	result = domainMakerInstnce.GetMaker(code)
	assert.Empty(t, result.ID)

}

func TestGetMakerList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var result *[]dbMaker.Schema
	var rows *sqlmock.Rows
	queryStr := "select * from maker order by name"
	tmpTime := time.Now()

	// カラムの要素
	var id1 int64 = 1
	name1 := "name_test1"
	code1 := "code_test"

	var id2 int64 = 100
	name2 := "name_test100"
	code2 := "code_test100"

	// 1件以上取れる場合
	rows = sqlmock.NewRows(getMakerColumns()).AddRow(id1, name1, code1, tmpTime, tmpTime).AddRow(id2, name2, code2, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WillReturnRows(rows)
	result = domainMakerInstnce.GetMakerList()
	assert.Equal(t, 2, len(*result))
	// 1個目
	assert.Equal(t, id1, (*result)[0].ID)
	assert.Equal(t, name1, (*result)[0].Name)
	assert.Equal(t, code1, (*result)[0].Code)
	assert.Equal(t, tmpTime, (*result)[0].CreatedAt)
	assert.Equal(t, tmpTime, (*result)[0].UpdatedAt)
	// 2個目
	assert.Equal(t, id2, (*result)[1].ID)
	assert.Equal(t, name2, (*result)[1].Name)
	assert.Equal(t, code2, (*result)[1].Code)
	assert.Equal(t, tmpTime, (*result)[1].CreatedAt)
	assert.Equal(t, tmpTime, (*result)[1].UpdatedAt)

	// データなし
	rows = sqlmock.NewRows(getMakerColumns())
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WillReturnRows(rows)
	result = domainMakerInstnce.GetMakerList()
	assert.Equal(t, 0, len(*result))

	// エラー
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WillReturnError(fmt.Errorf("some error case"))
	result = domainMakerInstnce.GetMakerList()
	assert.Equal(t, 0, len(*result))
}

func TestGetDetail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var result *dbMakerdetail.Schema
	var rows *sqlmock.Rows
	queryStr := "select * from maker_detail where maker_id IN (?)"
	tmpTime := time.Now()

	// カラムの要素
	var id int64 = 1
	ohpUrl := "ohp_url_test"
	twitterName := "twitter_name_test"
	youtubeChannel_id := "youtube_channel_id_test"
	youtubeKeywords := "youtube_keywords_test"

	// 1件取れる場合
	rows = sqlmock.NewRows(getMakerDetailColumns())
	rows = rows.AddRow(id, ohpUrl, twitterName, youtubeChannel_id, youtubeKeywords, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1").WillReturnRows(rows)
	result = domainMakerInstnce.GetDetail(id)
	assert.Equal(t, id, result.MakerID)
	assert.Equal(t, ohpUrl, result.OHP)
	assert.Equal(t, twitterName, result.TwitterName)
	assert.Equal(t, youtubeChannel_id, result.YoutubeChannelID)
	assert.Equal(t, youtubeKeywords, result.YoutubeKeywords)
	assert.Equal(t, tmpTime, result.CreatedAt)
	assert.Equal(t, tmpTime, result.UpdatedAt)

	// データなし
	rows = sqlmock.NewRows(getMakerDetailColumns())
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1").WillReturnRows(rows)
	result = domainMakerInstnce.GetDetail(id)
	assert.Empty(t, result.MakerID)

	// エラー
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1").WillReturnError(fmt.Errorf("some error case"))
	result = domainMakerInstnce.GetDetail(id)
	assert.Empty(t, result.MakerID)

}

func TestGetDetailList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var result *[]dbMakerdetail.Schema
	var rows *sqlmock.Rows
	var param []int64
	queryStr := "select * from maker_detail where maker_id IN (?)"
	tmpTime := time.Now()

	// カラムの要素
	var id1 int64 = 1
	ohpUrl1 := "ohp_url_test_1"
	twitterName1 := "twitter_name_test_1"
	youtubeChannel_id1 := "youtube_channel_id_test_1"
	youtubeKeywords1 := "youtube_keywords_test_1"

	var id2 int64 = 2
	ohpUrl2 := "ohp_url_test_2"
	twitterName2 := "twitter_name_test_2"
	youtubeChannel_id2 := "youtube_channel_id_test_2"
	youtubeKeywords2 := "youtube_keywords_test_2"

	var id3 int64 = 100
	ohpUrl3 := "ohp_url_test_100"
	twitterName3 := "twitter_name_test_100"
	youtubeChannel_id3 := "youtube_channel_id_test_100"
	youtubeKeywords3 := "youtube_keywords_test_100"

	// 1件取れる場合
	rows = sqlmock.NewRows(getMakerDetailColumns())
	rows = rows.AddRow(id3, ohpUrl3, twitterName3, youtubeChannel_id3, youtubeKeywords3, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("100").WillReturnRows(rows)
	param = []int64{id3}
	result = domainMakerInstnce.GetDetailList(param)
	assert.Equal(t, id3, (*result)[0].MakerID)
	assert.Equal(t, ohpUrl3, (*result)[0].OHP)
	assert.Equal(t, twitterName3, (*result)[0].TwitterName)
	assert.Equal(t, youtubeChannel_id3, (*result)[0].YoutubeChannelID)
	assert.Equal(t, youtubeKeywords3, (*result)[0].YoutubeKeywords)
	assert.Equal(t, tmpTime, (*result)[0].CreatedAt)
	assert.Equal(t, tmpTime, (*result)[0].UpdatedAt)

	// 2件以上取れる場合
	rows = sqlmock.NewRows(getMakerDetailColumns())
	rows = rows.AddRow(id1, ohpUrl1, twitterName1, youtubeChannel_id1, youtubeKeywords1, tmpTime, tmpTime)
	rows = rows.AddRow(id2, ohpUrl2, twitterName2, youtubeChannel_id2, youtubeKeywords2, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1,2").WillReturnRows(rows)
	param = []int64{id1, id2}
	result = domainMakerInstnce.GetDetailList(param)
	assert.Equal(t, 2, len(*result))
	// 1個目
	assert.Equal(t, id1, (*result)[0].MakerID)
	assert.Equal(t, ohpUrl1, (*result)[0].OHP)
	assert.Equal(t, twitterName1, (*result)[0].TwitterName)
	assert.Equal(t, youtubeChannel_id1, (*result)[0].YoutubeChannelID)
	assert.Equal(t, youtubeKeywords1, (*result)[0].YoutubeKeywords)
	assert.Equal(t, tmpTime, (*result)[0].CreatedAt)
	assert.Equal(t, tmpTime, (*result)[0].UpdatedAt)
	// 2個目
	assert.Equal(t, id2, (*result)[1].MakerID)
	assert.Equal(t, ohpUrl2, (*result)[1].OHP)
	assert.Equal(t, twitterName2, (*result)[1].TwitterName)
	assert.Equal(t, youtubeChannel_id2, (*result)[1].YoutubeChannelID)
	assert.Equal(t, youtubeKeywords2, (*result)[1].YoutubeKeywords)
	assert.Equal(t, tmpTime, (*result)[1].CreatedAt)
	assert.Equal(t, tmpTime, (*result)[1].UpdatedAt)

	// データなし
	rows = sqlmock.NewRows(getMakerDetailColumns())
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1,2").WillReturnRows(rows)
	param = []int64{id1, id2}
	result = domainMakerInstnce.GetDetailList(param)
	assert.Equal(t, 0, len(*result))

	// エラー
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs("1,2").WillReturnError(fmt.Errorf("some error case"))
	param = []int64{id1, id2}
	result = domainMakerInstnce.GetDetailList(param)
	assert.Equal(t, 0, len(*result))

}

func TestGetVideoList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var result *[]dbMakervideo.Schema
	var rows *sqlmock.Rows
	var makerID int64 = 1000
	queryStr := "select * from maker_video where maker_id = ? order by id DESC"
	tmpTime := time.Now()

	// カラムの要素
	var id1 int64 = 1
	videoID1 := "video_id_1"
	title1 := "title_1"

	var id2 int64 = 100
	videoID2 := "video_id_100"
	title2 := "title_100"

	var id3 int64 = 3
	videoID3 := "video_id_3"
	title3 := "title_3"

	// 1件取れる場合
	rows = sqlmock.NewRows(getMakerVideoColumns())
	rows = rows.AddRow(id1, makerID, videoID1, title1, tmpTime, tmpTime)
	rows = rows.AddRow(id2, makerID, videoID2, title2, tmpTime, tmpTime)
	rows = rows.AddRow(id3, makerID, videoID3, title3, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(makerID).WillReturnRows(rows)
	result = domainMakerInstnce.GetVideoList(makerID)
	assert.Equal(t, 3, len(*result))
	// id=1
	assert.Equal(t, id1, (*result)[0].ID)
	assert.Equal(t, makerID, (*result)[0].MakerID)
	assert.Equal(t, videoID1, (*result)[0].VideoID)
	assert.Equal(t, title1, (*result)[0].Title)
	assert.Equal(t, tmpTime, (*result)[0].PublishedAt)
	assert.Equal(t, tmpTime, (*result)[0].CreatedAt)
	// id=100
	assert.Equal(t, id2, (*result)[1].ID)
	assert.Equal(t, makerID, (*result)[1].MakerID)
	assert.Equal(t, videoID2, (*result)[1].VideoID)
	assert.Equal(t, title2, (*result)[1].Title)
	assert.Equal(t, tmpTime, (*result)[1].PublishedAt)
	assert.Equal(t, tmpTime, (*result)[1].CreatedAt)
	// id=2
	assert.Equal(t, id3, (*result)[2].ID)
	assert.Equal(t, makerID, (*result)[2].MakerID)
	assert.Equal(t, videoID3, (*result)[2].VideoID)
	assert.Equal(t, title3, (*result)[2].Title)
	assert.Equal(t, tmpTime, (*result)[2].PublishedAt)
	assert.Equal(t, tmpTime, (*result)[2].CreatedAt)

	// データなし
	rows = sqlmock.NewRows(getMakerVideoColumns())
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(makerID).WillReturnRows(rows)
	result = domainMakerInstnce.GetVideoList(makerID)
	assert.Equal(t, 0, len(*result))

	// エラー
	mock.ExpectQuery(regexp.QuoteMeta(queryStr)).WithArgs(makerID).WillReturnError(fmt.Errorf("some error case"))
	result = domainMakerInstnce.GetVideoList(makerID)
	assert.Equal(t, 0, len(*result))
}

func UpdateVideoList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainMakerInstnce := getMockDomain(db)

	var rowsDetail, rowsVideo *sqlmock.Rows
	var makerID int64 = 1000
	var result bool
	tmpTime := time.Now()

	// makerdetail
	ohpUrl := "ohp_url_test"
	twitterName := "twitter_name_test"
	youtubeChannelID := "youtube_channel_id_test"
	youtubeKeywords := "youtube_keywords_test"

	// makervideo
	var makerVideoID1 int64 = 1
	videoID1 := "video_id_1"
	title1 := "title_1"
	var makerVideoID2 int64 = 1
	videoID2 := "video_id_2"
	title2 := "title_2"

	makerDetailQueryStr := "select * from maker_detail where maker_id IN (?)"
	videoListQueryStr := "select * from maker_video where maker_id = ? order by id DESC"
	videoInsertQuery := domainMakerInstnce.DBMakervideo.CreateBulkInsertQuery(makerID, getTestYoutubeData(youtubeChannelID))

	// youtubeAPI mock
	domainMakerInstnce.DomainYoutube = &DomainYoutubeTest{}

	// 成功
	rowsDetail = sqlmock.NewRows(getMakerDetailColumns())
	rowsDetail = rowsDetail.AddRow(makerID, ohpUrl, twitterName, youtubeChannelID, youtubeKeywords, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(makerDetailQueryStr)).WithArgs("1000").WillReturnRows(rowsDetail)
	rowsVideo = sqlmock.NewRows(getMakerVideoColumns())
	rowsVideo = rowsVideo.AddRow(makerVideoID1, makerID, videoID1, title1, tmpTime, tmpTime)
	rowsVideo = rowsVideo.AddRow(makerVideoID2, makerID, videoID2, title2, tmpTime, tmpTime)
	mock.ExpectQuery(regexp.QuoteMeta(videoListQueryStr)).WithArgs("1000").WillReturnRows(rowsVideo)
	mock.ExpectExec(regexp.QuoteMeta(videoInsertQuery)).WillReturnResult(sqlmock.NewResult(1, 1))
	result = domainMakerInstnce.UpdateVideoList(makerID)
	assert.Equal(t, true, result)

	// 失敗
	mock.ExpectExec(regexp.QuoteMeta(videoInsertQuery)).WillReturnError(fmt.Errorf("some error case"))
	result = domainMakerInstnce.UpdateVideoList(makerID)
	assert.Equal(t, false, result)
}

func getMockDomain(db *sql.DB) *Maker {
	domainMakerInstnce := GetInstance()
	domainMakerInstnce.DBMaker.DBInstance = db
	domainMakerInstnce.DBMakerdetail.DBInstance = db
	domainMakerInstnce.DBMakervideo.DBInstance = db
	return domainMakerInstnce
}

func getMakerColumns() []string {
	return []string{"id", "name", "code", "created_at", "updated_at"}
}
func getMakerDetailColumns() []string {
	return []string{"maker_id", "ohp", "twitter_name", "youtube_channel_id", "youtube_keywords", "created_at", "updated_at"}
}
func getMakerVideoColumns() []string {
	return []string{"id", "maker_id", "video_id", "title", "published_at", "created_at"}
}
func getTestYoutubeData(channelID string) *[]domainYoutube.Video {
	publishedAt, _ := time.Parse("2006/01/02", "2020/01/01")
	return &[]domainYoutube.Video{
		domainYoutube.Video{"ID_test", channelID, "Description_test", publishedAt, "Title_test"},
	}
}
