# TiDB

業務でTiDBを利用することになりそうなので[ドキュメント](https://docs.pingcap.com/ja/tidbcloud/tidb-cloud-intro/)を読み、気になることなどをメモする

- TiDB Cloudのプラン別の機能比較 https://docs.pingcap.com/ja/tidbcloud/features/
  - 商用利用するなら基本的にはDedicatedになりそうだが
  - 「ワークロードに基づいた自動スケーリング」は非対応なのか..
- MySQLの互換性 https://docs.pingcap.com/ja/tidbcloud/mysql-compatibility/
  - サポートされていない機能のうち、トリガーとストアドあたりは要注意
  - 違い
    - 自動増分ID
      - 複数のTiDBノード間でIDを増分するにはAUTO_INCREMENT MySQL互換モードを使用する必要がある
      - デフォルトだと、1, 30001, 30002, 2のように番号が飛ぶ可能性あり
        - これは各TiDBノードが一定の数をキャッシュしているため（そのキャッシュする数のデフォルト値が30000）
        - TiDBノードが再起動すると、使われていない番号があっても、その番号は捨てられて、新たにキャッシュされる
      - ノードが保持する数を1に(`AUTO_ID_CACHE = 1`) にすると従来のMySQLのように連番になる
        - [https://zenn.dev/shigeyuki/articles/7ed78672a92061#v8.1.0以後-(2025%2F03%2F07追記)](https://zenn.dev/shigeyuki/articles/7ed78672a92061#v8.1.0%E4%BB%A5%E5%BE%8C-(2025%2F03%2F07%E8%BF%BD%E8%A8%98))
    - 組み込み関数
      - 大体サポートしているが、`SHOW BUILTINS;`でリスト表示可能
    - コレーション

      ```go
      デフォルトの照合順序:
      TiDB のデフォルトの照合順序はutf8mb4_binです。
      MySQL 5.7のデフォルトの照合順序はutf8mb4_general_ciです。
      MySQL 8.0 のデフォルトの照合順序はutf8mb4_0900_ai_ciです。
      ```
      →現行、utf8mb4_bin を使っているみたいなのでOKそう

- 移行
  - https://docs.pingcap.com/ja/tidbcloud/migrate-from-mysql-using-data-migration/
  - 移行方法
    - Data Migration
      - 無停止で、既存データを移行＆binlogで増分データ移行
    - Dumpling
      - 一括エクスポート
  - 1TiB未満の場合は、デフォルトの論理モードが推奨
    - 論理モード：ソースDBからデータをSQL文としてエクスポートし、それをTiDBで実行
    - 物理モード：ソースDBからデータをエクスポートし、KVペアとしてエンコードしてTiIKVに直接書き込む
    - 物理モードの方がパフォーマンスは良い
  - 移行手順 Aurora→TiDB Cloud
    - Aurora側の準備
      - バイナリログ有効化必要
        - パラメータグループの設定変更＆再起動が必要
    - Aurora側のデータをエクスポート
      - dumplingというツール
    - TiDB側でインポート
      - TiDB Lightningというツール
    - TiDB側で増分
      - binlogを受け取り、TiDB Cloudへの適用
  - 移行の方法
    - Dumpling vs DataMigration
    - 一定時間（容量次第だけど数時間ほど）無停止許容ならDumpling&Lightning
    - NoならDM
  - 移行時は移行用のMySQLユーザーを作成した方が良さそう
  - 既存環境からの接続方法 プライベートエンドポイント vs VPCピアリング or パブリック
    - プライベートエンドポイント
      - NLBを立てる必要がある

      ```go
      AWS は RDS またはAuroraへの直接の PrivateLink アクセスをサポートしていません。そのため、ネットワークロードバランサー (NLB) を作成し、ソース MySQL インスタンスに関連付けられたエンドポイントサービスとして公開する必要があります。
      ```

    - VPCピアリング
      - クラスタ作成時にCIDR設定が必要
      - これは既存のVPCのCIDRと重複させてはいけない

      **→今回立てるTiDBクラスタには複数のプロダクトのDBも引っ越ししてくる想定なので、VPCピアリングは破綻しそうなので、プライベートエンドポイントを選定する**
      →と思ったけど、そもそも接続が必要なのは移行時だけなので、既存VPCのCIDRと被らないようにTiDBクラスタを立ててVPCピアリングで良さげ？

  - 参考
    - [https://speakerdeck.com/staffrecruiter/yun-yong-zhong-taitorudemoyi-xing-dekiru-gemuye-jie-chu-aurorakaratidbhenogemudetayi-xing-nowu-tai-li](https://speakerdeck.com/staffrecruiter/yun-yong-zhong-taitorudemoyi-xing-dekiru-gemuye-jie-chu-aurorakaratidbhenogemudetayi-xing-nowu-tai-li?slide=90)
