# CORSとは？

## 概要

- (前提)SOP(Same Origin Policy)
  - ブラウザの機能として、SOPという仕組みがある
    - これはあるオリジンから別のオリジンにあるリソースへのアクセスを制限するというもの
      - オリジン＝スキーム（プロトコル）、ホスト、ポートの3つの組み合わせ(どれか異なるだけでも別オリジンと見なされる)
      - 例えば悪意あるスクリプトが仕込まれたWebページを読み込んだ時に、JSであらゆるリソースにアクセスできたらセキュリティ的な問題が生じうるので、それを防ぐための仕組み
    - ここで制限の対象となるのは、主にJavaScriptの話（fetch()とか、scriptタグでのフォントインストールとか）
      - なので、HTMLのimgタグのsrc属性に別オリジンとなるURLが設定されていても、SOPによる制限は受けない
    - これはあくまでブラウザの機能なので、curlとかだとSOPの仕組みはない（CORSエラーは発生しない）
- CORSは、異なるオリジンからのリソース要求を許可する機能(仕組み)
- `Access-Control-Allow-Origin`ヘッダーに許可したいドメインを設定することで、そのオリジンからの要求を許可する仕組み
  - 例えば、APIで`Access-Control-Allow-Origin`ヘッダーの値にフロントエンドドメインを設定することで、フロントエンドからのAPIアクセスが可能になる

（補足）
- ちょっと混同しがちなのは、CORSは異なるオリジンからのアクセスを制限する意味ではない
  - Cross Origin Request **Sharing**というくらいなので、異なるオリジンからのアクセスを**許可**するための仕組みのことを指す

### CORSエラーはどのようにして発生するか

CORSエラーが発生するタイミングは、リクエストの種類によって異なるので整理する

#### 1.シンプルリクエストの場合

- シンプルリクエストとは？（＝プリフライトリクエストが発生しないリクエスト）
  - HTTPメソッドはGET、POST、HEADのいずれか
  - リクエストヘッダーが以下のみ
    - Accept
    - Accept-Language
    - Content-Language
    - Content-Type（ただしapplication/x-www-form-urlencoded,multipart/form-data,text/plainのみ）
      - Authorizationやカスタムヘッダ(X-API-KEYなど)を含む場合は、非シンプルリクエストになる
  - リクエストボディにBlobやArrayBufferなどを含まない（通常のテキストデータのみ）
- ブラウザがサーバーにリクエストを送信し、ブラウザがレスポンスを受け取った後に`Access-Control-Allow-Origin`ヘッダーをチェックする
  - 許可されていない場合は、ブラウザがエラーを発生させ、レスポンスをJavaScript側に渡さない

#### 2.プリフライトリクエスト(が必要なリクエスト)の場合

- ブラウザがリクエスト送信前にOPTIONSメソッドでプリフライトリクエストを送る
- サーバーが許可しているかをレスポンスヘッダーで返す(ここで判定)
  - 許可されていれば、本リクエストを送信する
  - 許可されていなければ、本リクエストは送信されず、エラーが発生する

### 参考

- [オリジン間リソース共有 (CORS) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)
- [同一オリジンポリシー - ウェブセキュリティ | MDN](https://developer.mozilla.org/ja/docs/Web/Security/Same-origin_policy)