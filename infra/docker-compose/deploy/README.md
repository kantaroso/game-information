## 共通

### docker起動
```shell
docker-compose -f infra/docker-compose/deploy/docker-compose.yml up -d
```

## フロントエンド

### コンテナに入る
```shell
docker exec -it deploy-frontend sh
```

### ログイン
```shell
# ログイン
firebase login
# 確認
firebase login:list
```

### テスト
```shell
firebase emulators:start
```

### デプロイ
```shell
firebase deploy --only hosting
```

### 非公開にする
```shell
firebase hosting:disable
```

## バックエンド

### コンテナに入る
```shell
docker exec -it deploy-api sh
```

### ibmcloudにログインする
* [ワンタイムコードを取得する](https://identity-1.ap-north.iam.cloud.ibm.com/identity/passcode)
* リージョンはus-south（ダラス）
```shell
# ログイン
ibmcloud login --sso

# ログイン確認
ibmcloud account list

# ログアウト
ibmcloud logout
```

### デプロイする
```shell
ibmcloud target --cf
ibmcloud cf push gameinfomation
```

### 開始 / 停止
```shell
ibmcloud cf start gameinfomation
ibmcloud cf stop gameinfomation
```

### 補足補足

```shell
# goのルートディレクトリを変更した場合は下記のコマンドを使ってdockerのvolumeを書き換える
sh infra/docker-compose/deploy/filter_api_path.sh
vi infra/docker-compose/deploy/docker-compose.yml
```
