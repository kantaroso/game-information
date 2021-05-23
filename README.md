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

## ローカル構築手順

* [こちら](infra/docker-compose/README.md)


## クエリ

* [こちら](https://github.com/kantaroso/game-information/tree/master/documents/sql)

## dockerイメージのデプロイ

* ログイン
```
docker login
```

* [makeファイル](Makefile)で実行

## デプロイ

* [こちら](infra/docker-compose/deploy/vue/README.md)
