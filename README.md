# このリポジトリについて

## 環境構築手順

1. ソースをクローン
2. .envをコピー
```bash
cp .devcontainer/.env.sample .devcontainer/.env
```
3. VsCodeでプロジェクトフォルダーを開く
4. Reopen in Containerを押下
  Ctrl Shift P → Reopen in containerと入力して実行

## 動作確認

### ローカル環境

```bash
task infra:install
task infra:deploy:local

# リソース削除
task infra:destroy:local

# dynamoDBにレコード作成
go run internal/main.go 
```

### 実AWS環境

1. /home/dev/backend/.devcontainer/.env.dev.sampleの内容を.envにコピペする
2. 下記項目をセットする。
  ```dotenv
  AWS_ACCESS_KEY_ID=
  AWS_SECRET_ACCESS_KEY=
  ```
3．下記コマンド実行
```bash
task infra:install
task infra:deploy:aws

# dynamoDBにレコード作成
go run internal/main.go 

# リソース削除 ※削除不可設定があるのでテーブルはコンソールから削除してください。
task infra:destroy:aws
```