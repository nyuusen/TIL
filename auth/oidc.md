# OpenID Connect

## 登場人物
- ユーザー
- クライアントアプリ
- OpenIDプロバイダ

## OpenID Connectとは？
- クライアントアプリとOpenIDプロバイダ間のIDトークンの要求と応答を標準化したもの
- OAuth2.0の拡張


## 流れ
- クライアントアプリが署名付きの認証情報をOpenIDプロバイダに渡す
- OpenIDプロバイダが署名付きの認証情報の発行元に公開鍵を依頼し、署名を検証する
- 検証に成功したら、OpenIDプロバイダがIDトークンをクライアントアプリに発行する
  - 発行時にOpenIDプロバイダからユーザーに「本人ですか？発行しますか？」の確認を行う(=認証)
- 

## SAMLとの違い

## 感想

## 参考
- [一番分かりやすい OpenID Connect の説明 #OAuth - Qiita](https://qiita.com/TakahikoKawasaki/items/498ca08bbfcc341691fe)
