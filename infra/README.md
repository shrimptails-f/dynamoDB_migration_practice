# infra

実 AWS と `floci`(LocalStack) の両方に DynamoDB テーブルを作れる最小 Go CDK 構成です。

## 作成されるテーブル

Pocのため簡易的なバッチ履歴テーブルを想定して作成しています。
- Table name: `batch-history`
  - Partition key: `batch_name` (String)
  - Sort key: `executed_at` (String)
  - TTL attribute: `expires_at`
  - GSI: `status-index` (`status` + `executed_at`)
  - Billing mode: `PAY_PER_REQUEST`
  - Removal policy: `APP=local` のとき `DESTROY`、それ以外は `RETAIN`

## 環境構築時実行コマンド

```bash
cd ~/backend/infra

npm install
go mod tidy
```
## デプロイコマンド

### 実 AWS

AWS 認証情報を有効にした上で実行します。`AWS_PROFILE` を使うか、`AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_DEFAULT_REGION` を設定してください。

```bash
cd ~/backend/infra && npm run bootstrap && npm run deploy
```

`CDK_DEFAULT_ACCOUNT` と `CDK_DEFAULT_REGION` も実 AWS の値に合わせてください。`123456789012` のような LocalStack 用ダミー値のままだと意図した環境には出ません。

### LocalStack

ローカル環境ではAWS エミュレーターの `floci` に向くように環境変数を設定してから実行します。
特に `AWS_ENDPOINT_URL` は `http://floci.sandbox.localstack.cloud:4566`、`AWS_ENDPOINT_URL_S3` は `http://s3.floci.sandbox.localstack.cloud:4566` のように設定してください。`http://floci:4566` のままだと CDK のアセット公開時に `cdk-hnb659fds-assets-... .floci` のようなホスト名解決に失敗します。
この構成は `.devcontainer/docker-compose.yml` のネットワーク alias と対で動作し、`cdklocal` が path-style S3 を有効化する条件も満たします。
環境変数の例は .devcontainer/.env.sampleを参照してください

```bash
cd ~/backend/infra && npm run bootstrap:local && npm run deploy:local
```

テーブルを削除したい場合は以下です。

```bash
cd ~/backend/infra && npm run destroy
```

LocalStack 側を削除したい場合は以下です。

```bash
cd ~/backend/infra && npm run destroy:local
```
