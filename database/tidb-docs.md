https://docs.pingcap.com/ja/tidbcloud/

## トランザクション

- デフォルトは、悲観的トランザクションモード
- トランザクション
  - トランザクション分離レベル（REPETABLE READなど）
    - 読み取り
    - どの時点のデータを正として見るかの設定
    - TiDBでの設定
      - デフォルトは`REPEATABLE-READ`
        - トランザクション開始時点のスナップショットを見続ける
      - v4移行であれば常に「コミット済みの最新」を見に行く`READ-COMMITTED`が使用可能
  - トランザクションモード（楽観的・悲観的など）
    - 書き込み
    - 他者と衝突しそうな時、先にロックしておくか・最後に確認するか(衝突した場合はロールバック)の設定
    - TiDBでの設定
      - デフォルトは悲観的
  - FOR UPDATE
    - 更新前に、事前読み取りが必要かどうかをクエリごとに制御

  > TiDBはスナップショット分離（SI）一貫性を実装しており、MySQLとの互換性のために`REPEATABLE-READ`として宣伝されています。これはSI [ANSI繰り返し読み取り分離レベル](https://docs.pingcap.com/ja/tidbcloud/transaction-isolation-levels/#difference-between-tidb-and-ansi-repeatable-read)やSI [MySQL 繰り返し読み取りレベル](https://docs.pingcap.com/ja/tidbcloud/transaction-isolation-levels/#difference-between-tidb-and-mysql-repeatable-read)とは異なります。
  >
  > 詳細については[TiDBトランザクション分離レベル](https://docs.pingcap.com/ja/tidbcloud/transaction-isolation-levels/)参照してください。

## SQL

> 現在の`AUTO_INCREMENT`列が主キーで、型が`BIGINT`場合、 `ALTER TABLE t MODIFY COLUMN id BIGINT AUTO_RANDOM(5);`ステートメントを実行して`AUTO_INCREMENT`から`AUTO_RANDOM`に切り替えることができます。

→AUTO_INCREMENT BIGINTなら、AUTO_RANDOMへの切り替え可能

## 拡張性

> 各TiDB Cloud Dedicatedクラスタには少なくとも2つのTiDBノードを配置することをお勧めします。

→TiDBは最低2台が推奨

> TiKVは行ベースのデータの保存を担います。TiKVのノード数、vCPU、RAM、storageを設定できます。TiKVノードの数は少なくとも1セット（3つの異なる利用可能なゾーンに3ノード）で、3ノードずつ増加する必要があります。

→TiKVは1セットが3ノード

## 監視・アラート

- https://docs.pingcap.com/ja/tidbcloud/monitor-built-in-alerting/
  - CPUやメモリ使用率などでアラート発報が可能
  - サードパーティ製メトリクスサービスとアラート等の連携ができるっぽいが、CloudWatchは無し

## ストリームデータ

https://docs.pingcap.com/ja/tidbcloud/changefeed-sink-to-mysql/

- 「Sink to MySQL changefeed」という機能を使えば、MySQLにデータストリーミングできそう
  - TiDB移行後に性能面など何か致命的な問題が生じて元のDBに切り戻したい時に使えそう使えそう？

## バックアップ

https://docs.pingcap.com/ja/tidbcloud/backup-and-restore/#turn-on-auto-backup

> デフォルトでは、スナップショット バックアップは自動的に作成され、バックアップ保持ポリシーに従って保存されます。

- スナップショットバックアップはデフォルト有効化
- 頻度は日次 or 週次

## Security

https://docs.pingcap.com/ja/tidbcloud/security-concepts/

> **主要コンポーネント**
>
> - **アイデンティティおよびアクセス管理 (IAM )** : TiDB Cloudコンソールとデータベース環境の両方に対するセキュリティで柔軟な認証および権限管理。
> - **ネットワーク アクセス制御**: プライベート エンドポイント、VPC ピアリング、TLS 暗号化、IP アクセス リストなどの構成可能な接続オプション。
> - **データ アクセス制御**: 保存中のデータを保護するための顧客管理暗号化キー (CMEK) などの高度な暗号化機能。
> - **監査ログ**: コンソールアクションとデータベース操作の両方に対する包括的なアクティビティ追跡により、説明責任と透明性が確保されます。

- TiDB Cloudユーザーアカウント
  - 組織のSSO認証（エンタープライズ利用ならこれ一択だと思う）
    - OIDC/SAMLプロトコルで、企業IDプロバイダーと統合可能
    - MFA適用やパスワード有効期限ポリシーの設定も可能
    - 手順：https://docs.pingcap.com/ja/tidbcloud/tidb-cloud-org-sso-authentication/
      - コンソールURLは、専用のカスタムURLを使用する必要がある（組織設定 > 認証で設定）
- DBアクセス制御

  > データベース ユーザー アカウントは`mysql.user`システム テーブルに保存され、ユーザー名とクライアント ホストによって一意に識別されます。

  https://docs.pingcap.com/ja/tidb/stable/user-account-management/

- 実際の企業での運用想定
  - 組織SSO必須
  - DBユーザー・ロールを最小権限で作成
    - アプリケーションと開発者(運用担当)用に特定のDB配下のみなど
  - TiDBユーザー管理
    - 組織SSO
    - IPを社内VPNのみ可能
    - ロールを用意して、SQL Editor使用制限など
- ネットワークアクセス制御
  - PrivateLink利用したプライベートエンドポイント
  - TLS接続
    - アプリケーション — DB間をTLSを接続にする
    - DB側で証明書を発行し、CA証明書をアプリケーションに設定し、DBへの接続時にDBの証明書が本物かどうかを検証
    - その後、アプリケーション側は一時的な暗号鍵を作成＆DBの公開鍵で暗号化したデータ送信→DB側で復号

## クラスタを計画する

### TiDBノード

**※TiKVノードも同様だった**

> **4 vCPU、16 GiB**の TiDB は、学習、テスト、およびトライアル用途向けに設計されています。
>
> TiDB の vCPU と RAM サイズが**4 vCPU、16 GiB**に設定されている場合、次の制限に注意してください。
>
> - TiDB のノード数は 1 または 2 にのみ設定でき、TiKV のノード数は 3 に固定されています。
> - 4 vCPU TiDB は 4 vCPU TiKV でのみ使用できます。

- 8 vCPU、16 GiB or 16 vCPU、32 GiB あたりが良いのか？
  - サイジング基準がわからん
    - →https://docs.pingcap.com/ja/tidbcloud/tidb-cloud-performance-reference/
    - ベーシックなサイズを選定して負荷かけながら

> 一般的に、TiDBのパフォーマンスはTiDBノード数に比例して増加します。

## TiDBへの接続(プライベートエンドポイント)

https://docs.pingcap.com/ja/tidbcloud/set-up-private-endpoint-connections/

- 自VPC側にプライベートエンドポイントを作成し、TiDB側にNLBが配置され、NLBからTiDBクラスタに接続される構成
  - 安全かつプライベートで、データがパブリックインターネットに公開されることはない
  - プライベートエンドポイント＝NLB（とその先にあるTiDBクラスタへの接続設定など）という解釈で良さそう
  - プライベートエンドポイントとTiDBクラスタは同一リージョンにある必要がある

## クラスターの停止

- 一時停止すると、監視収集とコンピュートリソースのコストがかからなくなる
- 最大停止期間は7日間で、7日経過後は自動的に再開される

## メンテナンスウィンドウ

https://docs.pingcap.com/ja/tidbcloud/configure-maintenance-window/

## パフォーマンスの分析とチューニングの概要

https://docs.pingcap.com/ja/tidbcloud/tidb-cloud-tune-performance-overview/

まずは問題の原因がTiDBクラスタ内か外かを判断する

> 「概要」タブでレイテンシ(P80)を確認します。この値がユーザー応答時間のP80値よりも大幅に低い場合、主なボトルネックはTiDBクラスター外にある可能性があると判断できます。

- TiDBクラスタ内のボトルネック
  - 遅いSQLクエリを最適化します。
  - ホットスポットの問題を解決します。
  - 容量を拡張するには、クラスターをスケールアウトします。

## SQLチューニング

https://docs.pingcap.com/ja/tidbcloud/tidb-cloud-sql-tuning-overview/

- 重要なのは
  - 検索範囲を絞る
  - インデックスを貼る
  - 適切な結合アルゴリズムの選択（通常は自動で）
- TiDBコンソールの診断ページでプランが見れる

### 性能チューニングのベストプラクティス

https://docs.pingcap.com/ja/developer/dev-guide-optimize-sql-best-practices/#dml-best-practices

#### DML

- 複数行のステートメントを利用する（Bulk Insertなど）
  - ネットワークレイテンシやSQLパース回数・Redoログなどの発行回数が低減されるため
- 複数回実行するステートメントはPREPAREを利用する
  - SQL構文を解析することによるオーバーヘッドを低減
- 必要な列のみを選択する
  - ネットワークやディスクI/Oが余計にかかる
    - 特にBLOB型やTEXT型などが入っていると
  - インデックスを貼っている列だけであればカバリングインデックスが効く
- 一括削除・更新を利用する
  - 長時間排他ロックがかかる
  - トランザクションログが大量に発生する
  - LIMIT1000をつけるなどして更新・削除を繰り返し行うようにする
