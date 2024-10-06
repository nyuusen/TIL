# システムリプレイス

## はじめに

あるtoCサービスのリプレイスの話

## システム構成

- フロントエンド: HTML/JavaScript 
  - S3に静的ホスティング
- バックエンド: TypeScript/Fastify
  - ECS on Fargate
- データベース: RDS
- キャッシュ: CloudFront
- Firewall: AWS WAF

## リプレイスの概要

- 現在稼働しているシステムの運用会社の変更に伴うリプレイス
  - AWSアカウント自体も変更になる
- サービス自体は、別の外部サービスへの送客するだけの小さいサービス

## 進め方

- (事前)新環境の開発&デプロイ
  - 非公開状態
- 現環境をメンテナンスモード
    - 誤って旧環境でデータ作成等が行われないようにする
- 現環境→新環境にデータ移行
  - 事前にデータ移行用のバッチを作成
- 新環境の動作が問題ないかをテスト
  - 社内ネットワークのみに公開する
- 新環境のIP制限を解除し、外部に公開

## 勉強になったポイント

- データ移行バッチで扱うデータ量が多く、移行時にAWSマネージドサービスを使用する場合は、スロットリング制限を考慮して実装を行う
  - 今回はAWS KMSを使用して一部データを暗号化していたので、何も考えずに実装したらスロットリング制限に引っかかってしまった
- データ移行前に、現環境はメンテナンスモードにする
  - 万が一、データ移行中等に現環境の方にデータが作成されてしまったら、移行対象外となり、障害発生の可能性がある 
  - メンテナンスモードに切り替える際にCloud Frontのオリジンを変更する場合は、キャッシュ保持期間に気を付ける
    - キャッシュのInvalidateを行わないと、普通のページにアクセスできてしまう恐れがある
- (新環境リリース後)現環境メンテナンスモード内に配置するリンク等は新環境のものにしておく
    - オートリダイレクトにするのもあり
      - HTMLならmeta refreshでできる
        - `<meta http-equiv="refresh" content="リダイレクトするまでの秒数; URL=新URL">`
      - APIなら、HeaderのLocationフィールドに新URLをセットして30x系のリダイレクトのHTTPステータスコードを返す
      - JavaScriptを使用するのも可能だが、SEO的に微妙らしい