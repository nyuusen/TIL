# OpenID Connect

## 登場人物

- ユーザー
- クライアントアプリ
- OpenIDプロバイダ

## OpenID Connectとは？

- クライアントアプリとOpenIDプロバイダ間の**IDトークンの要求と応答**を標準化したもの
- OIDCと略される
- OAuth2.0の拡張仕様である
  - そのため処理フローが似ている

## OIDCとOAuth2.0の違い

- OIDCは人がシステムにログインするための仕組み(目的がログイン=認証)
- OAuth2.0はシステムがシステムにアクセスするための仕組み(目的がリソースアクセス＝認可)
- OAuth2.0のIdentityレイヤーを追加したもの(ログイン要件を満たすために)がOIDC
  - 具体的にはIDトークンの発行する/しないが違う部分になる

## IDトークンとは？

- 以下のような`ヘッダー.ペイロード.署名`で構成されている
  ```
  eyJraWQiOiIxZTlnZGs3IiwiYWxnIjoiUlMyNTYifQ.ewogImlz
  cyI6ICJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwKICJzdWIiOiAiMjQ4
  Mjg5NzYxMDAxIiwKICJhdWQiOiAiczZCaGRSa3F0MyIsCiAibm9uY2UiOiAi
  bi0wUzZfV3pBMk1qIiwKICJleHAiOiAxMzExMjgxOTcwLAogImlhdCI6IDEz
  MTEyODA5NzAsCiAibmFtZSI6ICJKYW5lIERvZSIsCiAiZ2l2ZW5fbmFtZSI6
  ICJKYW5lIiwKICJmYW1pbHlfbmFtZSI6ICJEb2UiLAogImdlbmRlciI6ICJm
  ZW1hbGUiLAogImJpcnRoZGF0ZSI6ICIwMDAwLTEwLTMxIiwKICJlbWFpbCI6
  ICJqYW5lZG9lQGV4YW1wbGUuY29tIiwKICJwaWN0dXJlIjogImh0dHA6Ly9l
  eGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyIKfQ.rHQjEmBqn9Jre0OLykYNn
  spA10Qql2rvx4FsD00jwlB0Sym4NzpgvPKsDjn_wMkHxcp6CilPcoKrWHcip
  R2iAjzLvDNAReF97zoJqq880ZD1bwY82JDauCXELVR9O6_B0w3K-E7yM2mac
  AAgNCUwtik6SjoSUZRcf-O5lygIyLENx882p6MtmwaL1hd6qn5RZOQ0TLrOY
  u0532g9Exxcm-ChymrB4xLykpDj3lUivJt63eEGGN6DH5K6o33TcxkIjNrCD
  4XB1CKKumZvCedgHHF3IAK4dVEDSUoGlH9z4pP_eWYNXvqQOjGs-rDaQzUHl
  6cQQWNiDpWOl_lxXjQEvQ
  ```
  - これはJSON Web Signature(JWS)形式で、3項目をそれぞれBase64エンコードした値
    - なのでそれぞれをBase64デコードする値が確認できる
  - ヘッダー
  - ペイロード
  - 署名
  -
- わかりやすい記事
  - [IDトークンが分かれば OpenID Connect が分かる #OAuth - Qiita](https://qiita.com/TakahikoKawasaki/items/8f0e422c7edd2d220e06)

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
