
## 起動手順
```shell
# docker 起動
docker-compose -f infra/docker-compose/docker-compose.yml up -d

# vue起動
# docker にログイン
docker exec -it front /bin/sh
npm run serve

# go 起動
docker exec -it api /bin/sh
go run main.go

# admin api
wget http://localhost/admin/create/mkaer/video/3rdeye -O -
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
