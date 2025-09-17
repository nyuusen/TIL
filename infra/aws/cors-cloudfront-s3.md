# CloudFront + S3構成でCORSを設定する

## 前提

- S3バケットに何かしらの資材を配置する
- S3バケットは外部公開しておらずCloudFront Distributionに設定(OAC)
- クライアント(ブラウザ) -- CloudFront(ディストリビューション) -- S3(バケット)という構成

## CORSおさらい

- https://github.com/nyuusen/TIL/blob/main/other/cors.md

## CloudFront + S3構成でCORSを設定する方法

実現したいことは「CloudFront + S3で返す内容にCORSの設定を追加したい」\
もう少し具体にすると「クライアントからのリクエストに対するレスポンスのヘッダに`Access-Control-Allow-Origin`とその値として該当オリジンを追加したい」となる。

そのためには以下のようにいくつかの手順が必要となる。

### STEP1: S3(オリジンサーバー)からAccess-Control-Allow-Originヘッダを返す

- バケットにCORS設定を作成する
- 詳細: [CORS 設定の要素 - Amazon Simple Storage Service](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManageCorsUsing.html)

### STEP2: CloudFrontディストリビューションからOrigin等のヘッダをオリジンサーバに転送するように設定する

- CloudFrontディストリビューションに、以下のヘッダをS3に転送するように設定する
  - Access-Control-Request-Headers
  - Access-Control-Request-Method
  - Origin
- 上記の設定は、マネージドポリシー「CORS-S3Origin」もしくは「CORS-CustomOrigin」に含まれているので、どちらかを選択すると良い

### STEP3: CloudFrontディストリビューションのキャッシュ動作を、HTTPリクエストのOPTIONSメソッドを許可する設定にする

- プリフライトリクエストが使用される場合は、OPTIONSメソッドを明示的に許可する必要がある
  - デフォルトでは、GETとHEADリクエストのみ許可されている

### STEP4: CloudFrontディストリビューションでレスポンスヘッダーポリシーを設定する

- CloudFrontディストリビューションのCORSを有効にするレスポンスヘッダーポリシーを追加する

### 参考

- [CloudFront からの「アクセス制御-オリジン許可ヘッダなし」エラーを解決する | AWS re:Post](https://repost.aws/ja/knowledge-center/no-access-control-allow-origin-error)
