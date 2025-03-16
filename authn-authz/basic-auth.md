# Basic認証

- username:passwordをBase64エンコードしてAuthorizationヘッダにセットする
- なぜbase64エンコードするのか？
  - base64は暗号化ではなくデコードが容易
  - base64エンコードすることでASCII文字で表現できるため(ヘッダはASCII文字で構成される)
- サーバー側の認証(照合)処理では単純な文字列比較ではなく、タイミング攻撃を考え、以下のような実装にするのが推奨される
  - https://zenn.dev/foxtail88/articles/constant-time-compare
- WWW-Authenticateで認証チャレンジを定義可能
  - 例えば、`WWW-Authenticate: Basic` とした場合は、Basic認証情報を入力させるダイアログを表示させて、ユーザーに入力をリクエストすることが可能
  - https://developer.mozilla.org/ja/docs/Web/HTTP/Reference/Headers/WWW-Authenticate