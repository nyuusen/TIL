# Basic認証

- username:passwordをBase64エンコードする
- なぜbase64エンコードするのか？
    - base64は暗号化ではなくデコードが容易
    - base64エンコードすることでASCII文字で表現できるため(ヘッダはASCII文字で構成される)
- サーバー側の認証(照合)処理では単純な文字列比較ではなく、タイミング攻撃を考え、以下のような実装にするのが推奨
    - https://zenn.dev/foxtail88/articles/constant-time-compare