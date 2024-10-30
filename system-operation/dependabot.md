# Dependabot

## Dependabotとは？

- GitHubによって提供される、プロジェクトの依存関係を自動的に管理・更新するツール
- 依存関係の脆弱性アップデートを通知してくれたり、自動的に提案するプルリクエストを作成してくれたりする
- 詳細は：[Dependabotとは](https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/about-supply-chain-security#what-is-dependabot)

## 使い方

- GitHubリポジトリの設定から有効化できる
  - Security > Code security and analysis
- .github/dependabot.ymlにチェック頻度や対象範囲を定義できる
- [Dependabot を使う - GitHub Docs](https://docs.github.com/ja/code-security/dependabot/working-with-dependabot)

## GitHubActionsワークフローを用いて効率的に運用する

- 例えば、
  - プルリクエスト作成時に自動でテスト実行する等のワークフローを組み込んでおく
  - dependabot/fetch-metadata アクションを用いてアップデート情報を取得する
    - 取得した結果をPRにコメントとしてぶら下げるみたいなこともできそう
- ただし、プルリクエストは人間の目で見て、マージすることが推奨されている(そりゃそう)
- [GitHub ActionsでのDependabotの自動化 - GitHub Docs](https://docs.github.com/ja/code-security/dependabot/working-with-dependabot/automating-dependabot-with-github-actions) 

## 注意点

### Docker

DependabotにはDockerfileのベースイメージ(FROMフィールド)に記述されたバージョンアップの対象とすることができるが、  
DockerComposeファイル(compose.yaml)に記述されたイメージ指定(imageフィールド)を対象とすることは現状できない。  

対策としては、DockerComposeにイメージをベタ書きせずにDockerfileに切り出すとかRenovateにするかとか...

参考：[Docker Compose files対応が待ち望まれているIssue](https://github.com/dependabot/dependabot-core/issues/390)