
## 起動手順
```shell
# ローカル起動
docker-compose -f infra/docker-compose/docker-compose.yml up -d
```

## frontend
```shell
# vue起動
# docker にログイン
docker exec -it front /bin/sh
# 初回のみ
# npm install
npm run serve
```

## api
```shell
# コンテナログ確認
docker-compose -f infra/docker-compose/docker-compose.yml logs -f go

# コンテナ ログイン
docker exec -it api /bin/sh
# テスト実行
go test -cover game-information/lib/domain/...

# cli
## スプレッドシートからゲーム情報を登録する
go run cmd.go makemaster maker

## youtubeの動画データを更新する ※引数指定で特定のゲームだけ更新可能
go run cmd.go makevideo
go run cmd.go makevideo august
```

## 初回構築手順
```
cd /var/www
vue create front


Vue CLI v4.5.3
? Please pick a preset: Manually select features
? Check the features needed for your project: Choose Vue version, Babel, TS, Router, Linter
? Choose a version of Vue.js that you want to start the project with 2.x
? Use class-style component syntax? Yes
? Use Babel alongside TypeScript (required for modern mode, auto-detected polyfills, transpiling JSX)? Yes
? Use history mode for router? (Requires proper server setup for index fallback in production) Yes
? Pick a linter / formatter config: Standard
? Pick additional lint features: Lint on save
? Where do you prefer placing config for Babel, ESLint, etc.? In package.json
? Save this as a preset for future projects? Yes
? Save preset as:
? Pick the package manager to use when installing dependencies: NPM


cd front
npm install bootstrap-vue bootstrap

```

## デザイン参考

### BootstrapVue
* https://bootstrap-vue.org/
