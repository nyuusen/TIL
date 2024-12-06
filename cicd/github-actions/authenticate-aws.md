# GitHubActions上におけるAWSの認証認可

## どのようにするのが良いか？

> AWS への API 呼び出しは認証情報で署名する必要があるため、AWS SDK または AWS ツールのいずれかを使用する場合は、AWS 認証情報と AWS リージョンを提供する必要があります。GitHub Actions でこれを行う 1 つの方法は、IAM 認証情報でリポジトリシークレットを使用することですが、これは長期認証情報の使用に関するAWS セキュリティガイドラインに準拠していません 。代わりに、長期認証情報または JWT を使用して一時的な認証情報を取得し、それをツールで使用することをお勧めします。この GitHub Action はまさにそれを実現します。

AWS公式が運用・公開しているGitHubActionsワークフローには、上記のように記載がある。
参考: [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file)

つまり、GitHubのSecretsに長期的なアクセス情報を保持するのは適切ではないので、**一時的な認証情報をやり取りするのが良い**というようなことが書かれている。

## 一時的な認証情報はどうやって？

結論、AssumeRoleを使用してやり取りする。

## ざっくりな手順(理解用)

- GitHubとAWSが認証情報をやり取りできるように設定
  - 具体的にはOpenID Connect
- AWS側でIAMロールを作成する
- GitHubActions側でIAMロールを指定してAssumeRoleを実行する
  - AWS CLIを実行してとかではなく、[aws-actions/configure-aws-credentials]という便利なワークフローを使用する

## 具体的な手順

- 基本的なことは以下に書かれている。
  - [アマゾン ウェブ サービスでの OpenID Connect の構成 - GitHub Docs](https://docs.github.com/ja/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services#adding-the-identity-provider-to-aws)

- 少し補足すると、AWS側では以下の設定が必要になる。
  - IDプロバイダとしてGitHubを登録する(OpenID Connect)
  - GitHubActionsが引き受けるIAMロール(それに紐付けるIAMポリシーも)
    - 信頼関係の条件には、`"token.actions.githubusercontent.com:sub": "repo:octo-org/octo-repo:ref:refs/heads/octo-branch"`のように指定することで、GitHubActionsからの引き受けを制限できる

## 参考

- [アマゾン ウェブ サービスでの OpenID Connect の構成 - GitHub Docs](https://docs.github.com/ja/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services#prerequisites)
- [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file)
- [OpenID Connect を使ったセキュリティ強化について - GitHub Docs](https://docs.github.com/ja/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect)
  - OpenID Connectについて詳しく書かれている。