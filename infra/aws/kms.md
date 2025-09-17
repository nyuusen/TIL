# KMS

## KMSとは

- 暗号化キーの作成と管理機能を提供するマネージドサービス
- [AWS Key Management Service - AWS Key Management Service](https://docs.aws.amazon.com/ja_jp/kms/latest/developerguide/overview.html)

## KMSの必要性

- データの暗号化に必要な暗号化キーを保護する必要がある
- キーの作成・管理・使用・削除がAWS KMS内でのみで行われるので、安全性が高い
  - 自前でこの辺りやるのは大変だけど、KMSで楽に安全にできるよということ

## KMSの料金

- 作成したキーに対して、1USD/月
- ローテーションすると2回目までは1USD/月
- APIリクエストあたり0.0x$
- [料金 - AWS Key Management Service | AWS](https://aws.amazon.com/jp/kms/pricing/)

## エイリアス

- 自動で生成される識別子は判別しづらいので、エイリアスを付与することで判別しやすくなる
- エイリアスを使用することで、キーをノーテーションしてもアプリケーション側での変更が不要

## KMSのポリシー

- KMS側にはキーポリシーがある
- キーポリシーでは暗号キーに対して、どのリソースに対して何をできるかを定義することで、特定のサービスからのみアクセスさせるという制御が可能
- 例えば、ECSタスクに割り当てるロールに、特定のKMSに対する暗号復号のアクションを許可するみたいなことができる
