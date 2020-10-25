## 仕様

### トップページ
* アクセスカウンター
  * PV
* メニュー

### メーカー詳細ページ
* 動画一覧


## ディレクトリ構成

## goのモジュール関連の使い方

```sell
# モジュールについて
# https://qiita.com/tana6/items/df9a48eecb84576f618d

# gopath以下で作業していないので必ずpackage名を明記する
# https://castaneai.hatenablog.com/entry/2019/02/22/151213
cd app/gin
go mod init github.com/kantaroso/game-information

# go get
# 必要なライブラリは都度 go get
go get github.com/gin-gonic/gin

# ローカルモジュールの設定
# app/gin/go.mod に replaceを記載
# https://qiita.com/moto_pipedo/items/c5859e9c4ad81cdbe750
cd app/gin/controllers
go mod init github.com/kantaroso/game-information/controllers

# go 仮実行
cd app/gin
go run main.go

```
##


## ローカル構築手順

### go

* [docker-compose](infra/docker-compose/gin/README.md)

### vue

* [docker-compose](infra/docker-compose/vue_gin/README.md)


## クエリ

* [こちら](https://github.com/kantaroso/game-information/tree/master/documents/sql)
