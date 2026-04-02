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
    - 組み込み関数
      - 大体サポートしているが、`SHOW BUILTINS;`でリスト表示可能