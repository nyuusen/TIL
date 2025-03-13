# Basic認証

- username:passwordをBase64エンコードしてAuthorizationヘッダにセットする
- なぜbase64エンコードするのか？
    - base64は暗号化ではなくデコードが容易
    - base64エンコードすることでASCII文字で表現できるため(ヘッダはASCII文字で構成される)
- サーバー側の認証(照合)処理では単純な文字列比較ではなく、タイミング攻撃を考え、以下のような実装にするのが推奨
    - https://zenn.dev/foxtail88/articles/constant-time-compare
- CloudFront Functionsで実装可能(Basic認証はライブラリとかも不要でサクッと簡易的に実装できるのが良い点)