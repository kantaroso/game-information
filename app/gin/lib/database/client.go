package lib/database

import (
	"github.com/kantaroso/game-information/config"
)

type DatabaseClient interface {
	query()
}

type DatabaseClient struct {
	DatabaseClient
	config config.Database
}

func GetDatabaseClient(){

}
