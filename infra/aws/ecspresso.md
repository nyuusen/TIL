# ecspresso

## はじめに

ECSデプロイツールであるecspressoについてまとめてみた。

## ecspressoとは？

- 発音は「エスプレッソ」
- Amazon ECSのためのデプロイツール
- タスク定義とサービスをファイルで管理する
- Goで実装され、シングルバイナルとして動作するCLI
- ソースコード: https://github.com/kayac/ecspresso
  - OSS(MIT LICENSE)

## 必要なもの

ecspressoを用いてECSデプロイを行うには以下が必要である

- インフラリソース
  - ecspressoはインフラリソースを作成・管理する機能はないため
- ecspresso(CLI)のインストール
- ecspressoの設定情報を定義するYAMLファイル
  - 使用するリージョンやECSクラスター/サービス/タスク定義のJSONファイル名等を記述する
- タスク及びサービスの設定情報を定義するJSONファイル
  - AWS CLI互換のJSONファイル

上記を用意した上で、ecspressoコマンドを実行すると、デプロイ等を実行してくれる

## 設計思想・特徴

基本的には「ECSのデプロイに関わる最小限のリソースのみを管理するツール」が主な設計思想である。  
どのような問題を解決するのかも含め、詳細については以下に記載する。

- サービスとタスク定義の管理に特化されている
  - 他のリソース(ex:VPCやRDS等)に比べ、タスク定義とサービスはデプロイの度に更新される
  - これらを単一のツールで管理するとなると、意図しない変更による障害発生のリスクが上がってしまう
  - ライフサイクルが異なるから、それぞれ別のツールを使った方が安全
  - さらに、アプリケーションデプロイには、アプリケーション側の知識(環境情報等)が必要なので、インフラ担当者じゃなくても扱いやすくなるというメリットもある
- Amazon ECS専用のツールとして作られている
  - 抽象度が高くならないため、理解が容易＋機能追加への追従も容易(作者談)
- 設定ファイルを作成する手間を削減
  - ecspresso initコマンドを利用することで、既にECSサービスから情報を参照して、ベースとなるJSONファイルを生成してくれる

## なぜ必要か？

そもそもTerraformがあればこのようなツールじゃないの？みたいな疑問をまず持ったので、なぜ必要かをまとめる

## 参考リンク

- [Amazon ECS デプロイツール ecspresso 開発5年の歩み - Speaker Deck](https://speakerdeck.com/fujiwara3/amazon-ecs-depuroituru-ecspresso-kai-fa-5nian-nobu-mi?slide=29)
- [ecspresso handbook v2対応版](https://zenn.dev/fujiwara/books/ecspresso-handbook-v2)
