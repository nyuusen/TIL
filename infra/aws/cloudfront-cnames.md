## CloudFrontの代替ドメイン名の移動やそもそもについて調べてみた

### 移動方法

システムリプレイスなどで同一ドメインに対してCloudFrontディストリビューションを切り替えたい時は、ディストリビューション設定の代替ドメインを移動する必要がある（重複登録すると**CNAMEAlreadyExistsエラーが発生する）**

同じAWSアカウントでの移動の場合は、

- AWS CLIのassociate-alias or update-domain-associationコマンドを実行
- DNSレコード（Aレコード）の向き先を変える

参考：https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/alternate-domain-names-move-options.html

他にも既存ディストリビューションの代替ドメイン名をワイルドカードに変更してから、新規ディストリビューションの代替ドメイン名を正規ドメインに変える方法もある（スワップ）

クロスアカウント間での移動は少し手間が増える（所有権検証（TXTレコード）や旧側の無効化などが必要になる）

https://repost.aws/knowledge-center/resolve-cnamealreadyexists-error

この際、「AWS CLIのupdate-domain-associationコマンドを実行」した後にCloudFrontのデプロイが完了した時点で、リクエストの飛び先が変わるのは注意ポイント。

### そもそも代替ドメイン名とは？

まずCloudFrontディストリビューションのDNS名前解決として、最初にCloudFront共通のエッジサーバ→CloudFront内で対象ディストリビューション特定という流れになる。この時の「CloudFront内で対象ディストリビューション特定」するタイミングで、必要になるのがこの代替ドメイン名である。だから重複登録しようとするとエラーになる。具体的には共通のエッジサーバーでHostヘッダを見ている。

で、Aレコードは何を返しているかというとCloudFrontエッジサーバーの共通のIP。だからAレコードの向き先を変える前から、CloudFront同士の切り替えであれば、代替ドメイン名を変えた時点でリクエストが届く先も変わってしまう。

ちなみにワイルドカードにした場合は完全一致で見つからなかった時のフォールバックされるドメインみたいな扱いになる。

このことを検証している記事もあった。

https://dev.classmethod.jp/articles/amazon-cloudfront-cname-and-host-header-test/
