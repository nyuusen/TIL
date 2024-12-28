# VPC Lambda

## VPC Lambdaとは

- VPCにアタッチしたLambda関数のこと
  - 各公式ドキュメントでは「接続」という言葉が使用されているが「アタッチ」と同義だと思っているのと、そっちの方がしっくりくる
  - ちなみにLambda関数は、デフォルトではLambdaマネージドVPCで実行される
- VPC内のリソースへのプライベートアクセスが可能になる点が最大のメリット(だと思われる)

## 構築に必要なもの

- ネットワークインターフェイス
- 実行ロールへのIAMポリシー(AWSLambdaVPCAccessExecutionRole)アタッチ
  - 上2つはどちらもLambdaがVPC内のリソースへのアクセスするために必要

## 構築手順

めちゃくちゃ端折ると..
- Lambda関数作成時に「VPCを有効化」を選択する
- Lambda関数の設定からVPCを選択し、サブネットを選択する

## 注意ポイント

- Lambda関数からインターネット接続したい場合はNATゲートウェイを経由する必要がある
  - これはサブネットがパブリックであっても同じ
  - 公式で記載されている情報を見ると、Lambda関数にはプライベートIPしか割り振られないためだと思われる
  - NATゲートウェイは高いので、代替案を考察している記事
    - [VPC内のLambdaからインターネット接続する方法 - Briswell Tech Blog](https://tech.briswell.com/entry/2023/10/03/202903)
    - 結論、VPC Lambda->VPC外 Lambdaでインターネットにアクセスする(VPCエンドポイント設定が必要)と行けるみたい。面白い。

## 参考

- [Lambda 関数に Amazon VPC 内のリソースへのアクセスを許可する - AWS Lambda](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-vpc.html)
- [インターネットアクセス可能な VPC Lambda を作成してみた | DevelopersIO](https://dev.classmethod.jp/articles/internet-access-vpc-lambda/)
- [VPC の Lambda 関数へのインターネットアクセスを許可する | AWS re:Post](https://repost.aws/ja/knowledge-center/internet-access-lambda-function)