# RDSのAuroraをアップデートする
## はじめに

業務で必要になったので、調べたことを雑にメモ。

## アップデート方法

### 1.インプレース(In-Place)

- 現行クラスターをそのままアップグレードする
- マネコンから1クリックで実行可能であり、エンドポイントも変わらないので作業工数が最もかからない
- 他の方法に比べダウンタイムが長い
  - データベース再起動が発生するため
- 切り戻しは、アップグレード前のスナップショットを用いて行う

### 2.ダンプ&リストア

- 移行先のDBクラスターを作成し、mysqldump等を利用してクラスター間でデータ移行させる
- データ量に応じて、ダンプにかかるエクスポート・インポート時間が発生する

### 3.レプリケーション

- 移行先のDBクラスターを作成し、binlogレプリケーションでクラスター間でデータ同期させた後、スイッチオーバーする
- データ量に関係なくダウンタイムが発生しない
- 手順が複雑
  - 現在では、RDSでBlue/Greenデプロイに対応しているので前よりは簡単にできるらしい
    - [より安全、簡単、迅速な更新のための Amazon RDS ブルー/グリーンデプロイを発表](https://aws.amazon.com/jp/about-aws/whats-new/2022/11/amazon-rds-blue-green-deployments-safer-simpler-faster-updates/)
    - [【衝撃】AWSのRDSがデータを失わないBlue/Greenデプロイに対応しました #reinvent | DevelopersIO](https://dev.classmethod.jp/articles/rds-bg-deploy/)

### 4.ブルー/グリーンデプロイ機能

- エンドポイント変わらない
- グリーンへの切り替え後に、ブルーへの逆方向レプリケーションが可能
  - 切り戻しが必要な時に迅速に対応が可能

#### 補足: binlogとは

- バイナリログのこと
- バイナリログには、テーブルやデータ変更等のイベントが記述される
- バイナリログをレプリカに転送することで、レプリカ側でトランザクションを再現して、同様のデータ変更が可能
- [MySQL :: MySQL 8.0 リファレンスマニュアル :: 5.4.4 バイナリログ](https://dev.mysql.com/doc/refman/8.0/ja/binary-log.html)


## 注意点等

- 更新にはDBの再起動が必要であるため、採用するアップデート方式によるが、通常20-30秒のダウンタイムが発生する
  - 再起動に必要な時間は、クラッシュ回復プロセス、再起動時のデータベースアクティビティ、および特定の DB エンジンの動作によって異なる
  - データベースアクティビティをできるだけ減らすこと(未完了のトランザクションのロールバックアクティビティを減らすことに繋がる)で短くできる
    - コミットされていないトランザクションがあると、それらをロールバックする処理が発生するからということだと思われる
    - つまり、システムをメンテナンスモードにしておく等しておき、トランザクションが滞留していない状態を作るのが良さげ

## 参考

- [Amazon Aurora の更新 - Amazon Aurora](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Updates.html)
- [Amazon Aurora MySQL 5.6を頑張らずに8.0へメジャーアップグレードしてみた | DevelopersIO](https://dev.classmethod.jp/articles/upgrade-aurora-mysql-5-6-to-8-0-simple-stupid/)

## 他社事例

### ユーザベース様

記事: [Aurora MySQL 2から3へのアップグレード - 安全性とコストを考慮した移行プロセス - Uzabase for Engineers](https://tech.uzabase.com/entry/2024/12/15/090000)

- 影響調査として、更新系クエリの動作確認はバイナリログレプリケーションを利用した
  - Auroraのブルーグリーンデプロイメントは、ブルー環境のブルー環境のクラスタに存在するDBインスタンスと同台数作成されてしまうのでコストが高いため
- アップグレード自体は、ブルー/グリーンデプロイ機能を利用した
  - 逆方向レプリケーションで切り戻しがしやすいため
- レプリケーションの方法が気になったので調べた
  - [Aurora MySQL のバイナリログレプリケーションの設定 - Amazon Aurora](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Replication.MySQL.SettingUp.html#AuroraMySQL.Replication.MySQL.RetainBinlogs)