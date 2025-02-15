# GitHubActionsでジョブ実行の条件に指定できるステータスの動作検証

## はじめに

GitHubActionsでジョブ実行の条件に指定できるが幾つがあるが、RequiredReviewが設定したEnvironmentにおいて、承認・承認却下した時の該当ステップの出力結果を検証したい。

## 検証リポジトリ

- https://github.com/nyuusen/github-actions-sandbox

## 検証対象

- success
- failure
- cancel

## 検証結果

### ワークフロー実行画面で「Cancel Workflow」を押下する

- success: false
- failure: false
- cancel: true

### 承認する

- success: true
- failure: false
- cancel: false

### 却下する

- success: false
- failure: true
- **cancel: false**

## まとめ

却下した時に、canceledがtrueにならない点は注意したほうが良さそう。
