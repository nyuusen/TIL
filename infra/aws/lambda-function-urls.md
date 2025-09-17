# Lambda Function URLs

## カスタムドメイン

- Route53でAレコードにエイリアスレコードとして設定できないのか？
  - 結論から言うと不可
  - Lambdaは、背後でAWSが負荷分散するからIPが固定ではない
  - CloudFront or ALBを前段に立てる必要あり
