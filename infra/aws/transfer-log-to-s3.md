# LambdaのログをS3に保存する.md

## 概要

Webhookを受け付けるAWS Lambda関数で出力するログをS3に保存(転送)したい

## 目的と要件

- ログを永続化し、簡易分析したい
  - 分析時はなるべく手をかけずにログを取得したい
- 実装工数をあまりかけたくない(一時的に必要な機能であるため)
- コストを安く抑えたい

## 実現方法

### CloudWatch LogsからS3にエクスポート

- AWSマネージドコンソールのCloudWatch　Logsのアクションメニューからログデータのエクスポートを選択する
- 対象日時等でフィルターをかけた上で指定するS3バケットにログファイルを出力する
- フィルター条件や対象範囲にカスタマイズ性を持たせたい場合は有効そう

### CloudWatch LogsからS3にエクスポート(自動実行)

- 上記の一連の流れをEventBridge Schedulerで自動実行させる
- 手順がとても簡単そう
  - [CloudWatch LogsをS3に転送するためにEventBridge Schedulerを使用する](https://zenn.dev/fy0323/articles/0c2b5b556d5a0a)
- 今回の安く簡単にという要件に最もマッチしそう

### Lambda関数からリアルタイムにログをS3に転送

- Lambda関数の処理の中でログの出力先(保存先)をS3に保存する
 - プログラム内でAWS SDKを使ってPutObjectメソッドを実行するイメージ 
- Lambdaでリクエストを受け付ける度にS3アクセスが発生するため、レイテンシが懸念になる
- ログ毎にファイル生成するのでは検索性が悪いので、1つのファイルに蓄積させる方法は(？)
  -　ちょっと調べたが、GetObjectした後に更新したコンテントを新たにPutObjectするしかなさそう..？
  - S3はオブジェクトストレージなので、ファイルに追記とかはできなそう 
- Lambdaのみで処理が完結するのはメリット(?)ではあるが、上記の懸念があるためあまり有効な選択肢にはならなそう

### その他

- CloudWatchLogsのサブスクリプションフィルターにて、KinesisDataFirehoseを指定し、KinesisDataFirehose側ではターゲットとなるS3を指定するKinesis Data Firehoseを使用してログをストリームさせる方法もあるっぽい

## 参考
[CloudWatchLogsのログをS3に転送する方法の比較 #AWS - Qiita](https://qiita.com/Regryp/items/031141f8930c94378d5f)
