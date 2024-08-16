# Environments

## Environmentsとは

- GitHubActionsの機能であり、リポジトリに設定できるもの
- 環境ごとにプロテクションルールや変数、シークレットを保持することができ、環境を指定して、ワークフローを指定することで、必要な権限のみでデプロイを実行できたりする
- [デプロイに環境の使用 - GitHub Docs](https://docs.github.com/ja/actions/managing-workflow-runs-and-deployments/managing-deployments/managing-environments-for-deployment)

## 使い方

- Environmentsの設定自体はGitHubの管理画面のSettingsから設定が可能
- ワークフローからはenvironmentフィールドにEnvironmentsの環境名を指定する

## ユースケース

- 環境ごとに必要な変数やシークレットを使用する
  - 例えば環境ごとに適切なAWSのロールを設定することで、セキュリティ性を向上させることができる

## 参考

- [[GitHub Actions] ブランチごとにジョブの実行を制御できる Environments を試してみた | DevelopersIO](https://dev.classmethod.jp/articles/github-actions-environment-secrets-and-environment-variables/)
