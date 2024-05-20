# OpenAPI仕様書生成&デプロイ

## 概要

以下を実現するGitHub Actionsのワークフローを作成した。

- OpenAPIで記述されたyamlファイルをredocly(npmパッケージ)でhtmlを生成
- gh-pagesブランチにコミット&プッシュするワークフローを作成

前提として、gh-pagesブランチへのコミット&プッシュされることでGitHubPagesにデプロイされる設定をしている

TODO: キャプチャ

## 実装したコード

実装したコードの最終形は以下。

```yml
name: Publish API doc

# トリガーは以下の1もしくは2どちらかを有効(コメント解除)する
# 1.手動実行のトリガーを設定
on: workflow_dispatch

# 2.mainへのmergeとopenapi.ymlの変更があれば実行されるトリガーを設定
# on:
#   push:
#     branches:
#       - main
#     paths:
#       - 'openapi.yml'

jobs:
  publish-api-docs:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout branch
        uses: actions/checkout@v4
      
      - name: Setup node
        uses: actions/setup-node@v4

      - name: Generate API doc
        run: npx @redocly/cli build-docs ./openapi.yml -o docs/api.html

      - name: Publish API doc
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git checkout -b gh-pages
          git add -f ./docs/api.html
          git commit -m "publish api doc"
          git push -f -u origin gh-pages


```

※本コードを実装したリポジトリ: [udemy-github-actions](https://github.com/nyuusen/udemy-github-actions)

## ポイント

### npxでredocly@cliコマンドを実行する

package.jsonが存在しないプロジェクトであることもあり、node環境をセットアップし、npxコマンドで実行するようにした

### コミットするgitユーザーどうしようか問題

前提として、自動ビルドしたhtmlをコミットするので、特定の実ユーザーではなく、bot等にコミットさせたい(同時にunknownユーザー等になってしまうことも避けたい)
結論、GitHub側で用意しているbotユーザーを使用することにした

```
git config user.name "github-actions[bot]"
git config user.email "github-actions[bot]@users.noreply.github.com"
```

TODO: キャプチャ

### git push時に403エラーが発生する

ワークフロー内で実行されるgit push時に以下のエラーが発生する

```
Permission to nyuusen/udemy-github-actions.git denied to github-actions[bot].
fatal: unable to access 'https://github.com/nyuusen/udemy-github-actions/': The requested URL returned error: 403
```

GUI上から以下で権限付与することで解決した
`settings -> Actions -> General -> Workflow permissions -> Read and write permissionsにチェックを入れる`

ワークフロー毎に権限を付与することも可能らしい
参考: [権限をジョブに割り当てる - GitHub Enterprise Cloud Docs](https://docs.github.com/ja/enterprise-cloud@latest/actions/using-jobs/assigning-permissions-to-jobs)

### git push時にリモートにある変更がローカルにないエラーが発生する
エラー全文は下記の通り。

```
To https://github.com/nyuusen/udemy-github-actions
 ! [rejected]        gh-pages -> gh-pages (fetch first)
error: failed to push some refs to 'https://github.com/nyuusen/udemy-github-actions'
hint: Updates were rejected because the remote contains work that you do not
hint: have locally. This is usually caused by another repository pushing to
hint: the same ref. If you want to integrate the remote changes, use
hint: 'git pull' before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
Error: Process completed with exit code 1.
```

gh-pagesブランチは、GitHubActionsワークフローからしかコミットされないため、
コマンドを`git push origin gh-pages`から`git push -f -u origin gh-pages`に変更することで解決。
