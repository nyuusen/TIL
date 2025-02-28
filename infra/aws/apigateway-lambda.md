# API Gateway + Lambda

## 概要

- API Gateway（Lambda統合）をterraformで実装したいので色々調査

## やりたいこと

- API GatewayでPOSTリクエストを受け付け、Lambda関数をトリガーする
- Lambda関数ではリクエストの一部フィールドをdumpしたい
- dump先はCloudWatchLogs
- Lambda関数はPythonで書きたい
- PythonソースコードはS3バケットにアップロードしたい

## 必要なリソースまとめ

- API Gateway本体
  - IAMロール（lambda:InvokeFunctionポリシーをアタッチ）も必要
  - パスやスキーマ定義用にOpenAPI（YAML or JSON）も必要
- Lambda関数
  - IAMロール（AWSLambdaBasicExecutionRoleポリシーをアタッチ）も必要
    - LambdaからCloudWatchLogsへの書き込みに必要
- Lambda関数を置くS3バケット
- Pythonコード群

## 調査メモ

### HTTPリクエストを使用してLambda関数を呼び出す方法
- API Gatewayの他にLambdaURLがある
  - [HTTP リクエストを使用して Lambda 関数を呼び出す方法を選択する - AWS Lambda](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/apig-http-invoke-decision.html)
  - シンプル・コスト効率を意識する場合は、LambdaURLで推奨
  - 大規模やOpenAPI Descriptionサポート、認証オプション、カスタムドメイン名..などの高度な機能が必要な場合にはAPI Gatewayが適している

### API GatewayのAPIタイプ

> HTTP API: 軽量で低レイテンシーの RESTful API。
REST API: カスタマイズ可能で機能豊富な RESTful API。
WebSocket API: 全二重通信のためにクライアントとの永続的な接続を維持するウェブ API。

- HTTPとRESTどっちを選べば良いか問題
  - [Choose between REST APIs and HTTP APIs - Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vs-rest.html)
  - RESTは機能が豊富、HTTPがシンプル(その分低価格)

### API GatewayのパスやスキーマはOpenAPIで管理が可能

- OpenAPI定義ファイルをインポートすることで、API Gatewayのパスやスキーマの管理が可能
- [API Gateway で OpenAPI を使用して REST API を開発する - Amazon API Gateway](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/api-gateway-import-api.html)
- リクエストのバリデーションも設定可能
  - [x-amazon-apigateway-request-validator プロパティ - Amazon API Gateway](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/api-gateway-swagger-extensions-request-validator.html)

### Lambda関数の更新方法

- S3バケットにコードを配置する場合、source_code_hashで変更有無と更新が可能
- つまり、S3バケットに更新後のコード(zip)を配置した後、terraformを流せば、Lambda関数の更新が行われる

### アクセス許可について

- https://techblog.kayac.com/aws-lambda-iam

## 参考

- [Amazon API Gateway エンドポイントを使用した Lambda 関数の呼び出し - AWS Lambda](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/services-apigateway.html)