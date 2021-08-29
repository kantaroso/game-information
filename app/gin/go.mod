module github.com/kantaroso/game-information

go 1.13

replace (
	// commands
	local.packages/game-information/commands => ./commands
	// config
	local.packages/game-information/config/database => ./config/database
	// controllers
	local.packages/game-information/controllers => ./controllers
	// debug
	local.packages/game-information/debug => ./debug
	// db
	// analysis
	local.packages/game-information/lib/db/analysis => ./lib/db/analysis
	local.packages/game-information/lib/db/analysis/accesslog => ./lib/db/analysis/accesslog
	// master
	local.packages/game-information/lib/db/master => ./lib/db/master
	local.packages/game-information/lib/db/master/maker => ./lib/db/master/maker
	local.packages/game-information/lib/db/master/makerdetail => ./lib/db/master/makerdetail
	local.packages/game-information/lib/db/master/makervideo => ./lib/db/master/makervideo
	// domain
	local.packages/game-information/lib/domain/accesslog => ./lib/domain/accesslog
	local.packages/game-information/lib/domain/log => ./lib/domain/log
	local.packages/game-information/lib/domain/maker => ./lib/domain/maker
	local.packages/game-information/lib/domain/spreadsheet => ./lib/domain/spreadsheet
	local.packages/game-information/lib/domain/youtube => ./lib/domain/youtube
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	local.packages/game-information/commands v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/config/database v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/controllers v0.0.0-00010101000000-000000000000
	local.packages/game-information/debug v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/analysis v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/analysis/accesslog v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/master v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/master/maker v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/master/makerdetail v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/master/makervideo v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/accesslog v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/log v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/maker v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/spreadsheet v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/youtube v0.0.0-00010101000000-000000000000 // indirect
)
