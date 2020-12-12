package database

import (
	"os"
	"strconv"
)

// dbはに標準ライブラリを使う
// https://blog.p1ass.com/posts/go-database-sql-wrapper/
// 理由
//   ・勉強のため
//   ・そこそこの機能が載っているらしい
// 参考
// https://rabbitfoot141.hatenablog.com/entry/2019/03/05/000551

// ConnectionOption sql.Open option
const ConnectionOption string = "%s:%s@tcp(%s:%d)/%s?parseTime=true"

// Database db接続情報
type Database struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// getDatabaseDefault db接続情報の取得 デフォルト
func newDatabase() *Database {
	conf := new(Database)
	conf.User = "root"
	conf.Password = "password"
	conf.Host = "mysql"
	conf.Port = 3306
	conf.Name = "default"
	return conf
}

// GetMaster masterの接続先
func GetMaster() *Database {
	conf := newDatabase()
	conf.Name = "master"
	if len(os.Getenv("DB_USER_MASTER")) > 1 {
		conf.User = os.Getenv("DB_USER_MASTER")
	}
	if len(os.Getenv("DB_PASSWORD_MASTER")) > 1 {
		conf.Password = os.Getenv("DB_PASSWORD_MASTER")
	}
	if len(os.Getenv("DB_HOST_MASTER")) > 1 {
		conf.Host = os.Getenv("DB_HOST_MASTER")
	}
	if len(os.Getenv("DB_PORT_MASTER")) > 1 {
		conf.Port, _ = strconv.Atoi(os.Getenv("DB_PORT_MASTER"))
	}
	return conf
}

// GetAnalysis analysisの接続先
func GetAnalysis() *Database {
	conf := newDatabase()
	conf.Name = "analysis"
	if len(os.Getenv("DB_USER_ANALYSIS")) > 1 {
		conf.User = os.Getenv("DB_USER_ANALYSIS")
	}
	if len(os.Getenv("DB_PASSWORD_ANALYSIS")) > 1 {
		conf.Password = os.Getenv("DB_PASSWORD_ANALYSIS")
	}
	if len(os.Getenv("DB_HOST_ANALYSIS")) > 1 {
		conf.Host = os.Getenv("DB_HOST_ANALYSIS")
	}
	if len(os.Getenv("DB_PORT_ANALYSIS")) > 1 {
		conf.Port, _ = strconv.Atoi(os.Getenv("DB_PORT_ANALYSIS"))
	}
	return conf
}
