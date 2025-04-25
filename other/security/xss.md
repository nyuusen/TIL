# XSS(Cross-Site Scripting)

## XSSとは？

- 悪意あるスクリプト（JavaScriptなど）をユーザーのブラウザ上で実行させる攻撃
- 例えば、HTML入力フォームに、`<script>alert('hello')</script>`という内容を入力する形でJSを実行できてしまう
  - 例えばCookie, LocalStorageにあるセッション・トークン情報を盗んだり

## 対策

### 発動を防ぐ: CSP

- サーバ側でContent-Security-Policy（CSP）を設定する(レスポンスにヘッダをセットする)
  - CSPは、ブラウザに命令をするもの(インラインのJS実行禁止とか、同一オリジンのJSのみ実行許可するとか)
  - ブラウザがCSPに書かれている命令を読み取り、それに則り行動する
- 設定例:
  - `script-src 'self'`: `<script src="./script/app.js></script>`といった 外部スクリプトの同一オリジンのみ許可（インライン＝HTML内のJSやbutton等のHTMLタグのonclick属性等のJSもブロック）

[コンテンツセキュリティポリシー (CSP) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Guides/CSP)

### 埋め込みを防ぐ: 入力値サニタイズ

- ユーザー入力値をエスケープ処理する
  - 例：`<script> → &lt;script&gt;`

### 読み取りを防ぐ: Cookie HttpOnly属性

- Cookie HttpOnly属性をつけることでJSからトークンを盗まれないようにできる

### 対策まとめ

サーバーからブラウザに変なJSを実行しないでね＋ユーザー入力値はきちんと文字列として扱ってね＋もし万が一Cookieにアクセスされても盗まれないようにね