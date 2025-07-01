# URLパス設計

## 特定のリソースを指すURL設計ベストプラクティス

- パスパターン(ex: `article/123`)とクエリパターン(ex: `article/?id=123`)のどちらが良いか？

### APIの場合

- 結論、`article/:123`の方が良い
- RESTで言うと、articleのidのものとURIの設計として自然
- 「123の記事」というのが一目でわかる（可読性高い）
- クエリストリングはフィルターのような意味合いである

### 画面の場合

- こちらも結論は`article/:123`の方が良い
- URLからリソースが明確
- ブラウザのURLバーを見ても読みやすい
- SEOにも有利
  - 検索エンジンはクエリパラメータ付きURLを嫌う傾向があり、パス形式の方がクローラにフレンドリー
  - 以下のGoogleドキュメントに「URLは短い方が良い(パラメータも少ない方が良い)」と書かれている（それ以外にも参考となる情報が多く書かれている）
  - 参考: [Google 検索の URL 構造に関するベスト プラクティス | Google 検索セントラル | ドキュメント | Google for Developers](https://developers.google.com/search/docs/crawling-indexing/url-structure?hl=en&visit_id=638869248042537543-2501025600&rd=1)

### 補足: Slugについて

- Slugとは、WebサイトやブログのURLの末尾に付与される、記事の内容を表す短い文字列のこと（`article/seo-best-practice`というURLパスで言うと`seo-best-practice`の部分）
- Slugはユーザーからの見やすさとSEOに良い影響を与える
  - ユーザーからの見やすさ
    - `article/8b69af99-57ab-47f3-b8be-9b5ecfdf6f9f`と`article/seo-best-practice`のURLがあったとしたら、
      - どちらが見やすいか？
      - URLから内容が推測しやすいか？
      - どちらが信頼できるか？（怪しさがないか？）
  - SEOへの影響
    - GoogleはURLテキストをランキングシグナル(検索順位を決めるための要素)の1つとしているので、ランダム文字列ではなくわかりやすい(&短い)タイトルの方がSEOに有利