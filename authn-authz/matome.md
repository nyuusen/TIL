# 現代のWebアプリケーション(HTTP)における認証認可の技術まとめ

## はじめに

MDNに良い感じにまとめてあった。
[HTTP 認証 - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Authentication)

## 認証認可を実現する技術(その仕様)

- ID/パスワード
- SAML
- OAuth
- OpenID Connect
- JWT
- MFA

### Basic認証

- IDとパスワードをBASE64エンコードしてHTTPヘッダのAuthorizationに設定する
- 実装が簡単
- しかし盗聴されるとBase64デコードすることでセキュリティ性は低くない
  - というわけでHTTPS化することが推奨される

### ダイジェスト認証

- Basic認証の改良版
- サーバが生成したランダムの文字列をパスワードに付与し、IDとパスワードをMD5でハッシュ化して送信
- サーバ側でも登録済みの認証情報をハッシュ化し、照合することで認証する

### セッションベースの認証

- リクエストの度にパスワード等を送るのは


- トークンベースの認証

## サーバー側で認証認可に関連する情報を保持する方法

- セッション
  - データベースやRedisなどの外部ストレージに保存する

## クライアント側で認証認可に関連する情報を保持する方法

- LocalStorage
  - 実装は楽(setItemとかremoveItemなどのAPIを実行するだけ)
  - 一方、JavaScriptを使用するのでXSSで盗まれる可能性がある
- Cookie
  - HTTP Onlyを属性を付加することでJavaScriptからの操作を無効にできる
- インメモリ
  - JavaScriptでブラウザのメモリ内に保存する(クロージャに入れるとかするイメージ)
  - リロードするとログアウト状態になる、タブ間でログイン状態が共有されないなどの問題がある
- Auth0
  - Silent Authenticationという仕組みがある
    - インメモリ形式で実装しているが、リロード時にログアウトしてしまう問題がないのが利点
    - 実際の仕組みとしては、HTMLのiframeタグを駆使している
  - 認証基盤のロックインや一定規模が超えた場合に課金が発生するという課題はある
- 参考
  - [認証用トークン保存先の第4選択肢としての「Auth0」 | ログミーBusiness](https://logmi.jp/main/technology/324349)

- 読みたい記事や見たい動画
  - [SPAセキュリティ入門～PHP Conference Japan 2021 | ドクセル](https://www.docswell.com/s/ockeghem/ZM6VNK-phpconf2021-spa-security)
  - [Cookieにまつわるセキュリティ - YouTube](https://www.youtube.com/playlist?list=PLWiFLcGkQgLx8lbno3zZEinqu5C19hwET)

## 認証技術

### 

## 参考

- [HTTP 認証 - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Authentication)
