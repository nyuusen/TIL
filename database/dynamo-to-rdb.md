# DynamoからRDBへの移行

## 背景

業務でDynamoDBをRDB(TiDB)に移行する時の調査メモ。  
特に知りたいのがデータモデリングの

## モデリング

- 現行のDynamoDBのプライマリキーからRDBのプライマリキーを設計する
  - Partition Keyのみ → そのままPKへ
  - Partition Key + Sort Key → 複合PK
- インデックス
  - GSI,LSIどちらも構造そのままインデックスへ
    - ただしTiDBでは内部的にインデックスはKey-Valueをそのまま持つので無駄なインデックスは排除したい
- データ型
  - 文字列型や数値型は実際に格納される値からVARCHARやINT型にする