module github.com/kantaroso/game-information

go 1.13

require (
	github.com/gin-gonic/gin v1.6.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/kantaroso/game-information/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/kantaroso/game-information/controllers v0.0.0-00010101000000-000000000000
	github.com/kantaroso/game-information/lib/accesslog v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
)

replace (
	github.com/kantaroso/game-information/config => ./config
	github.com/kantaroso/game-information/controllers => ./controllers
	github.com/kantaroso/game-information/lib/accesslog => ./lib/accesslog
)
