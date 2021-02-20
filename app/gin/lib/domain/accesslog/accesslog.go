package accesslog

import (
	"net/http"

	dbAccesslog "local.packages/game-information/lib/db/analysis/accesslog"
)

// Accesslog インスタンス
type Accesslog struct {
	DBAccesslog *dbAccesslog.Accesslog
}

// GetInstance インスタンス生成
func GetInstance() *Accesslog {
	return &Accesslog{DBAccesslog: dbAccesslog.GetInstance()}
}

// GetAccessCount アクセスカウンター
func (domain Accesslog) GetAccessCount() int {
	return domain.DBAccesslog.CountAll()
}

// Register アクセスログ記録
func (domain Accesslog) Register(r *http.Request) bool {
	return domain.DBAccesslog.Insert(r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("USer-Agent"))
}
