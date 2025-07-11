# ECSのロギング設定

- タスク定義でログドライバ(`logConfiguration.logDriver`)を設定可能
- シンプルにCloudWatchLogsに吐かせるなら`awslogs`を設定するだけで良いが、例えば特定の監視ツールにも転送したいとかなると、FluentBitのようなログルータ的なものが必要になる
  - FluentBitは、ログを集約・整形して転送してくれるログの便利屋さん
- FluentBitの動かし方？としてはいくつか選択肢がある

## サイドカー

- 同一タスク内に別コンテナとしてサイドカーとしてFluentBitコンテナを動かす
- アプリケーションコンテナからは標準出力 or 特定ファイルにログを出力し、ホストとそれぞれが同一のボリュームマウントすることで、FluentBitコンテナを内部的にはそのファイルをtailするなどすることでログを収集してくれる
- ECS設定的にはcontainerDefinitions(配列)に`fluent/fluent-bit:latest`のような形でタスクを定義する

## Daemonサービス

- データプレーンとしてEC2を採用している場合、ホストに別プロセスとしてFluentBitを常駐させておく
  - 内部的には、コンテナで動作するアプリケーションが標準出力したログを出力内容をキャプチャし、`logConfiguration.logDriver`(大抵の場合は`json-file`)に出力、それを別プロセスで動いているFluentBitがtailコマンドなりでリアルタイムに監視する
- 1コンテナ1プロセスの原則に反する
  - 外から見た時に片方が落ちていることも気づきにくいので可観測性が落ちる（ので非推奨）

## FireLens

- AWS推奨＆公式ECSロギング機構
- タスク定義のログドライバに`awsfirelens`を設定するだけ
- 内部的には背後にFluentBitがサイドカーとして自動で立ち上がる

## 業務での設定例

- ECSのログドライバとしてCloudWatch Logsを設定
- CloudWatch Logsにログが吐かれたらLambdaを実行し、LambdaでNewRelicにログ転送
  - Lambdaトリガーはサブスクリプションフィルターを使って、特定のロググループに新しいログが出力されるたびにLambdaにストリーミングする設定
- メトリクスは、メトリクスストリームを設定して、NewRelicで収集