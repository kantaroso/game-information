build-docker-vue-base:
	docker build -f ./infra/docker/vue/base/Dockerfile -t kantaroso/game-information-vue-base:latest .
	docker push kantaroso/game-information-vue-base:latest
build-docker-firebase:
	docker build -f ./infra/docker/firebase/Dockerfile -t kantaroso/game-information-firebase:latest .
	docker push kantaroso/game-information-firebase:latest
build-docker-ibmcloud:
	docker build -f ./infra/docker/ibmcloud/Dockerfile -t kantaroso/game-information-ibmcloud:latest .
	docker push kantaroso/game-information-ibmcloud:latest
build-docker-go-dev:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/go/gin/dev/Dockerfile -t kantaroso/game-information-go-dev:${TAGDATE} -t kantaroso/game-information-go-dev:latest .
	docker push kantaroso/game-information-go-dev:${TAGDATE}
	docker push kantaroso/game-information-go-dev:latest
build-vue-assets-prod:
	docker-compose -f ./infra/docker-compose/docker-compose.yml run --rm vue npm run build
build-go-main:
	docker run -it -v ${PWD}/app/gin:/go/src/game-information -w=/go/src/game-information kantaroso/game-information-go-dev:latest go build main.go
