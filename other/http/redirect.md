# リダイレクトについて

## はじめに

リダイレクト3xx系のHTTPステータスコードについて調べた。
歴史的背景とかが絡んで、仕様と実態の乖離等が発生し、わかりづらいものになっているらしいので整理。

## 301 Moved Permanently

- リクエストされたリソースが**完全**にHTTPヘッダーのLocationで示されたURLに移動したことを示す
- ブラウザはこのURLにリダイレクトするが、検索エンジンはリソースへのリンクを更新**する**
- 一部の古いクライアントでは、不正にメソッドがGETに書き換えられる恐れがあるので、GETリクエストのみで使用する
- SEO的には最も優れている？
- 参考: [301 Moved Permanently - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/301)

## 302　Found

- リクエストされたリソースが**一時的**にHTTPヘッダーのLocationで示されたURLに移動したことを示す
- ブラウザはこのURLにリダイレクトするが、検索エンジンはリソースへのリンクを更新**しない**
- 一部の古いクライアントでは、不正にメソッドがGETに書き換えられる恐れがあるので、GETリクエストのみで使用する
- 参考: [302 Found - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/302)

## 303 See Other

- リダイレクトが新しくアップロードされたリソースではなく、 (確認ページやアップロード進捗ページのような) 別なページにリンクすることを示す
- PUTやPOSTの結果として送り返される
- リダイレクト先の表示にはGETを使用する
- 参考: [303 See Other - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/303)

## 307 Temporary Redirect

- リクエストされたリソースが**一時的**にHTTPヘッダーのLocationで示されたURLに移動したことを示す
- リクエストメソッドと本文が変更されないことが保証される
  -　使用されるメソッドをGETに変更したい場合は303を使用する
- 参考: [307 Temporary Redirect - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/307)

## 308 Permanent Redirect

- リクエストされたリソースが**完全**にHTTPヘッダーのLocationで示されたURLに移動したことを示す(301と同じ)
- リクエストメソッドと本文が変更されないことが保証される
- 参考: [308 Permanent Redirect - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/308)
 
## ステータスコード別のユースケース

ではどれを使えば良いのかというのを調べてみる

### 301 Moved Permanently

- サイトの移動
- GETリクエストを維持したい
- ex: コーポレートサイトのドメイン変更等

### 302　Found

- 一時的なページを表示する
- GETリクエストを維持したい
- ex: メンテナンスページを一時的に表示したい時等

### 303 See Other

- 一時的なページを表示する
- GET以外のリクエストを受け付けて、GETでリダイレクトさせる
- ex: POSTやPUTメソッドでフォーム送信やファイルアップロードを行い、その確認ページを表示したい時等

### 307 Temporary Redirect

- 一時的なリダイレクト処理
- GETメソッド以外で、そのメソッドを引き継ぎたい時
- ex: POSTメソッドで受け付けるAPIが一時的に別のAPIにPOSTメソッドでリダイレクトしたい場合

### 308 Permanent Redirect

- 恒久的なリダイレクト処理
- GETメソッド以外で、そのメソッドを引き継ぎたい時
- ex: POSTメソッドで受け付けるAPIが恒久的に別のAPIにPOSTメソッドでリダイレクトしたい場合

## リダイレクト手法

- サーバー側でステータスコードとLocationヘッダーを設定する
- meta refreshを設置する
  - `<meta http-equiv="refresh" content="秒数;URL=URL">`というような書き方で実現可能
  - よく「10秒後に移動します」みたいなサイトがあると思うが、多分この方法でリダイレクトさせているのではと思う
- JavaScriptで実装する(window.location.hrefなど)
  -　ただしこれだとSEO観点で検索エンジンがURL変更を認識できず宜しくないらしい
    - canonicalタグで代替可能か？(`<meta><link ref="canonical" href="new_url" />...</meta>`)
 
## 所感

これまではtoBサービスを開発することが多かったので、SEOとかあまり気にしていなかったけど、toCサービスだとSEO観点も含めて実現方法を検討する必要があるんだと感じた。APIなら301リダイレクトで、静的なHTMLならmeta refreshで実装するのが良さそうなのかな。
  
