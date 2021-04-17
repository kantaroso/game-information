build-vue-base:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/vue/base/Dockerfile -t kantaroso/game-information-vue-base:latest .
	docker push kantaroso/game-information-vue-base:latest
build-vue-prod:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/vue/laravel/Dockerfile -t kantaroso/game-information-vue-prod:${TAGDATE} -t kantaroso/game-information-vue-prod:latest .
	docker push kantaroso/game-information-vue-prod:${TAGDATE}
	docker push kantaroso/game-information-vue-prod:latest
build-go-dev:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/go/gin/dev/Dockerfile -t kantaroso/game-information-go-dev:${TAGDATE} -t kantaroso/game-information-go-dev:latest .
	docker push kantaroso/game-information-go-dev:${TAGDATE}
	docker push kantaroso/game-information-go-dev:latest
build-go-pord:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker run -it -v ${PWD}/app/gin:/go/src/game-information -w=/go/src/game-information kantaroso/game-information-go-base:latest go build main.go
	docker build -f ./infra/docker/go/gin/prod/Dockerfile -t kantaroso/game-information-go-prod:${TAGDATE} -t kantaroso/game-information-go-prod:latest .
	docker push kantaroso/game-information-go-prod:${TAGDATE}
	docker push kantaroso/game-information-go-prod:latest
