# このリポジトリについて

ローカルでAWSエミュレーターへのDynamoDBテーブル作成・管理を本番運用に向けてIaC管理として検討したPoCです。  
エミュレーターにはライセンス料金の観点から[floci](https://github.com/floci-io/floci)を採用しています。  
また、IaCにはAWS CDK/Goを採用しバックエンドとインフラ構成管理を兼任したリポジトリとしています。  

DynamoDBにおけるマイグレーションの検討については[こちら](./docs/dynamodb_spec.md)を参照してください。

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

1. テーブル作成
```bash
# テーブル作成
task infra:install
task infra:deploy:local

# サンプルデータ投入
task seed

# dynamoDBにレコード作成
go run internal/main.go 

# リソース削除をする場合
# task infra:destroy:local
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