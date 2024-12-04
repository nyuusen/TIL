# VPCピアリング

## VPCピアリングとは？

> プライベート IPv4 アドレスまたは IPv6 アドレスを使用して 2 つの VPC 間でトラフィックをルーティングすることを可能にするネットワーク接続です。
> どちらの VPC のインスタンスも、同じネットワーク内に存在しているかのように、相互に通信できます。

[VPC ピア機能とは - Amazon Virtual Private Cloud](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/what-is-vpc-peering.html)

## 料金

- 接続の確立には料金かからない
- AZ内での接続でのデータ通信も料金かからない
- ただし、異なるAZやリージョン間でのデータ通信には料金が発生する
- [VPC ピア機能とは - Amazon Virtual Private Cloud](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/what-is-vpc-peering.html)

## VPCピアリング接続の構築手順

詳細は[VPC ピアリングのプロセス、ライフサイクル、制限 - Amazon Virtual Private Cloud](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/vpc-peering-basics.html)に記載されている。

アカウントAのVPCからアカウントBにVPCに接続する前提だと、ざっくりと以下の流れ。

- アカウントAでVPCピアリングを作成
  - この時に接続先としてアカウントBのVPCを選択する
  - 作成したら「承諾の保留中」というステータスになる
  - アカウントAのVPCはリクエスタという扱いになる
- アカウントBで該当のピアリング接続を選択し「リクエストを承諾」を選択する
- アカウントAの対象VPCのルートテーブルを更新する
  - ターゲットには「ピアリング接続」を選択し、送信先には接続先VPCのCIDRを指定する
- アカウントBの対象VPCのルートテーブルを更新する
  - アカウントAでの設定の逆を設定する
  - [VPC ピアリング接続のルートテーブルを更新する - Amazon Virtual Private Cloud](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/vpc-peering-routing.html)
- アカウントBで対象リソースのセキュリティグループを変更する
  - アカウントA側のリソースにアタッチしているセキュリティグループorアカウントAのVPCのCIDR
    - 前者の方が狭く絞れる分、セキュリティ性は高い
    - ちなみに、両方のVPCが同じリージョンにある場合のみ、相手のVPCのセキュリティグループを参照できる
- アカウントAのVPCのリソースから、アカウントBのリソースへ接続する
  - ルートテーブルに従って、ピアリング接続で接続できる

## 注意点

- 複数のVPCピアリング接続はできない(必ず1対1である必要がある/[参考](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-basics-multiple))
- プライベートIPで名前解決を行うので、それぞれのVPCでCIDRが重複してはいけない([参考](https://docs.aws.amazon.com/ja_jp/vpc/latest/peering/vpc-peering-basics.html#:~:text=%E3%82%A2%E3%82%AF%E3%82%BB%E3%83%97%E3%82%BF%20VPC%20%E3%81%AF%E3%83%AA%E3%82%AF%E3%82%A8%E3%82%B9%E3%82%BF%20VPC%20%E3%81%AE%20CIDR%20%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF%E3%81%A8%E9%87%8D%E8%A4%87%E3%81%99%E3%82%8B%20CIDR%20%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF%E3%82%92%E4%BF%9D%E6%8C%81%E3%81%99%E3%82%8B%E3%81%93%E3%81%A8%E3%81%AF%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%9B%E3%82%93%E3%80%82))

## 役立つシーン

- VPC跨ぎのデータ取得や移行など
