package accesslog

import (
	"net/http"
)

// GetAccessCount アクセスカウンター
func GetAccessCount() int {
	return countAll()
}

// Register アクセスログ記録
func Register(r *http.Request) {
	insert(r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("USer-Agent"))
}
