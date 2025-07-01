# WebAPI The Good Parts

# 1章 WebAPIとは何か

- 本書でのWebAPIの定義は「HTTPプロトコルを利用してネットワーク越しに呼び出すAPI」

# 2章 エンドポイントの設計とリクエストの形式

- 良いエンドポイント設計
  - **覚えやすくどんな機能を持つURIなのかが一目で分かる**
  - 具体的には、
    - 短い
    - 人間が読める
    - 省略形は使わない
    - 大文字小文字混在させない
    - 単数系複数形、パラメータ指定方法を統一する
- HTTPメソッド
  - POST
    - 新しいリソースを送信する・登録する
    - 既存の更新・修正・削除はPUT・DELETEが本来は正しい
  - PUT
    - 既存リソースの完全上書き
  - PATCH
    - 既存リソースの一部を変更する
  - X-HTTP-Method-Overrideヘッダ
    - 古いブラウザや一部のライブラリなどではPATCHやDELETEなどのメソッドに対応していないケースがある
    - そういったケースに対する対応策としては、以下の2つがある
      - X-HTTP-Method-Overrideヘッダ
        - リクエスト自体はPOSTにして、X-HTTP-Method-Overrideヘッダに本来使用した値(DELETEなど)を設定する
      - _methodパラメータ
        - Formのパラメータの1つとして、application/x-www-form-urlencodedというContent typeで表されるデータの一部として送信される
          - `user=test&_method=PUT`
    - X-HTTP-Method-Overrideヘッダの方が好ましい（理由：_methodはContentTypeが限られるのと、リクエスト本文に関係のないメタ情報が入ってしまう点が微妙）
    - ちなみにX-HTTP-Method-Overrideヘッダを使用すると、一旦ブラウザからはPOSTで送信するけど、サーバー側ではDELETE等で処理を行うことを期待すると言う意味合い
      - なのでサーバー側ではこのヘッダがセットされた時の実装を考慮する必要がある
- URLとメソッド
  - `user/`: 一覧系
  - `user/:id`: 詳細系
  - 上記のエンドポイントに対して、GETなら取得、POSTなら登録、DELETEなら削除といったように、エンドポイントで「あるデータの集合」や「個々のデータ」表現し、HTTPメソッドで操作を表すのがWebAPI設計の基本中の基本となる
- URL設計に関しては少しだけ深掘りしてメモを残した: https://github.com/nyuusen/TIL/blob/f8c0a5e7b01f9d2bd9b5e9fb66f658e0e8147c2f/system-design/url-path.md