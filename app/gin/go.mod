module github.com/kantaroso/game-information

go 1.17

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
	local.packages/game-information/commands v0.0.0-00010101000000-000000000000
	local.packages/game-information/controllers v0.0.0-00010101000000-000000000000
)

require (
	cloud.google.com/go v0.72.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.2.0 // indirect
	go.opencensus.io v0.22.5 // indirect
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20201201145000-ef89a241ccb3 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/api v0.36.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201201144952-b05cb90ed32e // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	local.packages/game-information/config/database v0.0.0-00010101000000-000000000000 // indirect
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
