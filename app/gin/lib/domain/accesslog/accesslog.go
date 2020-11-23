package accesslog

import (
	"net/http"

	dbAccesslog "local.packages/game-information/lib/db/accesslog"
)

// GetAccessCount アクセスカウンター
func GetAccessCount() int {
	return dbAccesslog.CountAll()
}

// Register アクセスログ記録
func Register(r *http.Request) {
	dbAccesslog.Insert(r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("USer-Agent"))
}
