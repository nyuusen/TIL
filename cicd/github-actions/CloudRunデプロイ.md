# CloudRunデプロイ

## はじめに

個人で開発したCloud Runで動いているGoアプリケーションデプロイをGitHub Actionsでパイプラインを構築し、
デプロイ作業を簡略化したい

## デプロイの流れ

- Google Cloud認証
- Google Cloud SDKセットアップ
- 最新のコミットハッシュを取得(①)
- Dockerイメージビルド(タグ名に①を埋め込む)&ArtifactRegistryへイメージプッシュ
- Cloud Runでプッシュされたイメージ(タグ名)を指定してデプロイ

## 手順毎に調べながら進める


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

### Google Cloud SDKセットアップ

`google-github-actions/setup-gcloud@v1`が使える

### 最新のコミットハッシュを取得

以下のコマンドで取得できる

```sh
git log --pretty=%H -1
```

$GITHUB_ENVにセットすることで変数として以降のステップで使用可能になる

```sh
echo "image_hash=$(git log --pretty=%H -1)" >> "$GITHUB_ENV"
```

### Dockerイメージビルド&ArtifactRegistryへイメージプッシュ

ローカルで使用したコマンドを指定すれば良いが、プロジェクトID等の情報はGitHubのSecretsで管理する
(正直な所、どの値をSecretsで管理すれば良いか分かっていないので過剰にSecretsで管理しすぎている感がある)

```yml
- name: Docker build and push
  run: |
    docker build -t asia-northeast1-docker.pkg.dev/${{ secrets.PROJECT_ID }}/${{ secrets.REPOSITORY_NAME }}/api:${image_hash} -f deploy/api/Dockerfile --platform linux/amd64 .
    docker push asia-northeast1-docker.pkg.dev/${{ secrets.PROJECT_ID }}/${{ secrets.REPOSITORY_NAME }}/api:${image_hash}
```

上記のコードで以下のエラーが発生した。

```sh
denied: Unauthenticated request. Unauthenticated requests do not have permission "artifactregistry.repositories.uploadArtifacts" on resource "projects/***/locations/asia-northeast1/repositories/***" (or it may not exist)
```

リソースが存在していない？  
→ローカル実行すると通るので存在してそう

アタッチしているIAMの問題か？  
→該当IAMには`artifactregistry.repositories.uploadArtifacts`の権限が付与されていることを確認

雑にエラー文でググってみたら以下の記事がヒットし、同じ方法で解決  
[コンテナをpushすると発生した denied: Permission "artifactregistry.repositories.uploadArtifacts" denied on resource - show log include yuh](https://yunabe.hatenablog.com/entry/2023/06/04/220618)

具体的には、docker push前に以下のコードを追加した  

```yml
- name: Configure Docker
  run: gcloud auth configure-docker asia-northeast1-docker.pkg.dev
```

恐らくだが、ログイン済のgloudユーザーとdockerの紐付け？設定？が必要だったのかなと思われる

### Cloud Runでプッシュされたイメージを指定してデプロイ

gcloud run deployコマンドを実行するだけ。  
[gcloud run deploy  |  Google Cloud CLI Documentation](https://cloud.google.com/sdk/gcloud/reference/run/deploy)