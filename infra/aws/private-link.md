# PrivateLinkを理解する

## PrivateLinkとは？

- サービスがVPC内にあるかのようにVPCをそれらのサービスにプライベートに接続できるようにするサービス
  - あるVPCから別のVPCのあるリソースをあたかも同一VPC内にホストされているかのように、プライベートIPアドレスを使用して接続することを可能にする
- 後述のVPCエンドポイントVPCエンドポイントサービスを使用して、上記の機能が実現することができ、その提供される機能（サービス）のことをPrivateLinkという
- IGW、NATデバイス、パブリックIPアドレス、DirectConnect、Site-to-Site VPN接続を使用する必要はない
  - セキュリティが向上するのはもちろん、NATゲートウェイ（インスタンス）やEIPが不要ということでコスト面でもメリットが大きい
- [AWS PrivateLink とは](https://docs.aws.amazon.com/ja_jp/vpc/latest/privatelink/what-is-privatelink.html)
- [AWS PrivateLink の概念](https://docs.aws.amazon.com/ja_jp/vpc/latest/privatelink/concepts.html)

## VPCエンドポイントとは？

- VPCと他サービス間でプライベート接続を可能にするコンポーネント
- **サービス利用側**のVPC内に作成する
- (インターフェイス型)VPCエンドポイントを作成するとENIが作成（アタッチ）される
  - VPCエンドポイントを作成するときは1つ以上のサブネットを選択するが、選択したサブネットそれぞれにENIが作成される

## VPCエンドポイントサービスとは？

- VPCと他サービス間でプライベート接続を可能にするコンポーネント
- **サービス提供側**のVPC内に作成する
- 実体はロードバランサー
  - NLB or GLBが必要
- VPCエンドポイントサービス側では、VPCエンドポイントの接続リクエストを承諾を行う必要がある

## VPCエンドポイントが存在するのはプライベートサブネットだけ？

## VPCエンドポイントの種類

### インターフェイスタイプ

- ENIが作成される

### ゲートウェイタイプ

- ルートテーブルで送信先に対するターゲットとして指定する

## インターフェースVPCエンドポイントの料金

- 時間課金: VPCエンドポイント1つあたり 0.014USD/1h
- データ処理量: 0.01USD/1PB

※ゲートウェイ型は、データ処理量のみ発生する

## ユースケース

## 参考

- [【初心者向け】VPCエンドポイントとAWS PrivateLinkの違いを実際に構築して理解してみた | DevelopersIO](https://dev.classmethod.jp/articles/aws-vpcendpoint-privatelink-beginner/)
