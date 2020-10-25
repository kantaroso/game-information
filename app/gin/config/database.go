package config

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

// Database db接続情報
type Database struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// GetDatabase db接続情報の取得
func GetDatabase() Database {

	var conf Database

	conf.User = "root"
	conf.Password = "password"
	conf.Host = "mysql"
	conf.Port = 3306
	conf.Name = "kanta_workspace"
	if len(os.Getenv("DB_USER")) > 1 {
		conf.User = os.Getenv("DB_USER")
	}
	if len(os.Getenv("DB_PASSWORD")) > 1 {
		conf.Password = os.Getenv("DB_PASSWORD")
	}
	if len(os.Getenv("DB_HOST")) > 1 {
		conf.Host = os.Getenv("DB_HOST")
	}
	if len(os.Getenv("DB_PORT")) > 1 {
		conf.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	}
	if len(os.Getenv("DB_DATABASE")) > 1 {
		conf.Name = os.Getenv("DB_DATABASE")
	}

	return conf
}
