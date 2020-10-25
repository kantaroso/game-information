build-vue-base:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/vue/base/Dockerfile -t kantaroso/game-information-vue-base:latest .
	docker push kantaroso/game-information-vue-base:latest
build-vue-prod:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/vue/laravel/Dockerfile -t kantaroso/game-information-vue-prod:${TAGDATE} -t kantaroso/game-information-vue-prod:latest .
	docker push kantaroso/game-information-vue-prod:${TAGDATE}
	docker push kantaroso/game-information-vue-prod:latest
build-go-pord:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker run -it -v ${PWD}/app/gin:/go/src/game-information -w=/go/src/game-information golang:1.13.8-alpine3.11 go build main.go
	docker build -f ./infra/docker/go/gin/Dockerfile -t kantaroso/game-information-go-prod:${TAGDATE} -t kantaroso/game-information-go-prod:latest .
	docker push kantaroso/game-information-go-prod:${TAGDATE}
	docker push kantaroso/game-information-go-prod:latest
