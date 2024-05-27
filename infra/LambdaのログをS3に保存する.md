# LambdaのログをS3に保存する.md

## 概要

Webhookを受け付けるAWS Lambda関数内でリクエストログをS3に保存したい

## 目的と要件

- 対象はWebhookで受け付けたリクエストログ
- ログを永続化したい
- ログを簡易分析したい（処理の成功/失敗を見るレベル)
- S3に保存したい
- 実装工数をあまりかけたくない
- コストを安く抑えたい

## 実現方法

### CloudWatch LogsからS3にエクスポート

- AWSマネージドコンソールのCloudWatch　Logsのアクションメニューからログデータのエクスポートを選択する
- 対象日時等でフィルターをかけた上で指定するS3バケットにログファイルを出力する
- 自動化するにはこの処理をLambda等に書き起こし、EventBridgeで定期実行してあげる必要がある

### Lambda関数からリアルタイムにログをS3に転送

- Lambda関数の処理の中でログの出力先(保存先)をS3に保存する
- Lambdaのみで処理が完結する
- Lambdaでリクエストを受け付ける度にS3アクセスが発生するため、レイテンシは若干悪くなりそう

### その他

- CloudWatchLogsのサブスクリプションフィルターにて、KinesisDataFirehoseを指定し、KinesisDataFirehose側ではターゲットとなるS3を指定するKinesis Data Firehoseを使用してログをストリームさせる方法もあるっぽいがコスト観点と実装工数的な意味で選択肢から除外

## 参考
[CloudWatchLogsのログをS3に転送する方法の比較 #AWS - Qiita](https://qiita.com/Regryp/items/031141f8930c94378d5f)
