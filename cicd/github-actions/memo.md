# GitHub Actionsに関する調べたことを雑に書く

## $GITHUB_OUTPUT vs $GITHUB_ENV

- $GITHUB_OUTPUT
  - 特定のステップで使用する等出力先を細かく指定したい場合に使用する
- $GITHUB_ENV
  - 後続の(複数の)ステップで値を参照できるようにしたい場合に使用する
- 基本的には$GITHUB_ENVを使用するで良いらしい

- 参考
  - [Workflow commands for GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-environment-variable)
  - [Difference between environment variable and output parameters · community · Discussion #55294](https://github.com/orgs/community/discussions/55294)