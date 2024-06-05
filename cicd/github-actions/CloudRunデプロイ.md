# CloudRunデプロイ

## はじめに

個人で開発したCloud Runで動いているGoアプリケーションデプロイをGitHub Actionsでパイプラインを構築し、
デプロイ作業を簡略化したい

## デプロイの流れ

- 最新のコミットハッシュを取得(①)
- Google Cloud認証
- Dockerイメージビルド(タグ名に①を埋め込む)
- Google Cloud SDKセットアップ
- ArtifactRegistryへイメージプッシュ
- Cloud Runでプッシュされたイメージ(タグ名)を指定してデプロイ

## 手順毎に調べながら進める

### 最新のコミットハッシュを取得

以下のコマンドで取得できる

```
git log --pretty=%H -1
```

### Google Cloud認証

以前はデプロイロールを紐づけたサービスアカウントを作成し、サービスアカウントキーを発行するのが一般的だったみたいだが、
サービスアカウントキーは強力な認証を持つため漏洩時のリスクが大きいという課題があるらしく、
Workload Identity連携で行うのが推奨とされているらしい。

以下の記事の通りに進めた。  
[Workload Identity 連携を利用して GitHub Actions を動かす](https://zenn.dev/cloud_ace/articles/7fe428ac4f25c8)  

**(そもそも)サービスアカウントとは？**
- 人以外のリソース（CloudRun等）が使用するアカウントであり、そのリソースが持つ認証情報である
- リソースにサービスアカウントをアタッチして、そのサービスアカウントに色々なロールをすることで、Google Cloud内部の他リソースに対するAPI等を実行することができるようになる

**Workload Identity連携とは？**

- Workload Identity 連携は、外部サービスとの認証方法の1つ
  - SAMLとOIDCをサポートしている
- サービスアカウントキーを使用する方法よりセキュアである
  - Workload Identity プールは、外部IDを管理するエンティティ
- GitHubと連携したい場合は、Workload Identity プール プロバイダとしてGitHubを選択する
- 参考
  - [Workload Identity 連携  |  IAM のドキュメント  |  Google Cloud](https://cloud.google.com/iam/docs/workload-identity-federation?hl=ja)

**参考**    
[デプロイメント パイプラインとの Workload Identity 連携を構成する](https://cloud.google.com/iam/docs/workload-identity-federation-with-deployment-pipelines?hl=ja#impersonation)  
[Google Cloud Platform での OpenID Connect の構成 - GitHub Docs](https://docs.github.com/ja/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-google-cloud-platform#adding-a-google-cloud-workload-identity-provider)

### Dockerイメージビルド(タグ名に①を埋め込む)
### Google Cloud SDKセットアップ
### ArtifactRegistryへイメージプッシュ
### Cloud Runでプッシュされたイメージ(タグ名)を指定してデプロイ
