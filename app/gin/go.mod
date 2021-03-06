module github.com/kantaroso/game-information

go 1.13

replace (
	// config
	local.packages/game-information/config/database => ./config/database
	// controllers
	local.packages/game-information/controllers => ./controllers
	// debug
	local.packages/game-information/debug => ./debug
	// db
	local.packages/game-information/lib/db/accesslog => ./lib/db/accesslog
	local.packages/game-information/lib/db/maker => ./lib/db/maker
	local.packages/game-information/lib/db/makerdetail => ./lib/db/makerdetail
	local.packages/game-information/lib/db/makervideo => ./lib/db/makervideo
	// domain
	local.packages/game-information/lib/domain/accesslog => ./lib/domain/accesslog
	local.packages/game-information/lib/domain/maker => ./lib/domain/maker
	local.packages/game-information/lib/domain/youtube => ./lib/domain/youtube
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	local.packages/game-information/config/database v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/controllers v0.0.0-00010101000000-000000000000
	local.packages/game-information/debug v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/accesslog v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/maker v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/makerdetail v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/db/makervideo v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/accesslog v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/maker v0.0.0-00010101000000-000000000000 // indirect
	local.packages/game-information/lib/domain/youtube v0.0.0-00010101000000-000000000000 // indirect
)
