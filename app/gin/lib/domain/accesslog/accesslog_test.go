package accesslog

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccessCount(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainAccesslogInstance := getMockDomain(db)

	var result int
	var rows *sqlmock.Rows

	// 1件以上取れる場合
	rows = sqlmock.NewRows([]string{"count(id)"}).AddRow(12345)
	mock.ExpectQuery(regexp.QuoteMeta("select count(id) from access_log")).WillReturnRows(rows)
	result = domainAccesslogInstance.GetAccessCount()
	assert.Equal(t, 12345, result)

	// 0件の場合
	rows = sqlmock.NewRows([]string{"count(id)"}).AddRow(0)
	mock.ExpectQuery(regexp.QuoteMeta("select count(id) from access_log")).WillReturnRows(rows)
	result = domainAccesslogInstance.GetAccessCount()
	assert.Equal(t, 0, result)

	// エラーになった時
	mock.ExpectQuery(regexp.QuoteMeta("select count(id) from access_log")).WillReturnError(fmt.Errorf("some error case"))
	result = domainAccesslogInstance.GetAccessCount()
	assert.Equal(t, 0, result)

}

func TestRegister(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	domainAccesslogInstance := getMockDomain(db)

	r, _ := http.NewRequest("GET", "http://example.com/hogehoge?a=a&b=b", nil)

	var result bool

	// 成功時
	mock.ExpectExec(regexp.QuoteMeta("insert into access_log(method,endpoint,query_string,user_agent) values(?,?,?,?)")).WithArgs(r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("USer-Agent")).WillReturnResult(sqlmock.NewResult(1, 1))
	result = domainAccesslogInstance.Register(r)
	assert.Equal(t, true, result)

	// エラー時（特になもしない）
	mock.ExpectExec(regexp.QuoteMeta("insert into access_log(method,endpoint,query_string,user_agent) values(?,?,?,?)")).WithArgs(r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("USer-Agent")).WillReturnError(fmt.Errorf("some error case"))
	result = domainAccesslogInstance.Register(r)
	assert.Equal(t, false, result)

}

func getMockDomain(db *sql.DB) *Accesslog {
	domainAccesslogInstance := GetInstance()
	domainAccesslogInstance.DBAccesslog.DBInstance = db
	return domainAccesslogInstance
}
