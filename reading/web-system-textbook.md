# 作りながら学ぶWebシステムの教科書

## 概要

読書メモ。気づきや調べたことをまとめていく。

- URL: https://info.nikkeibp.co.jp/media/LIN/atcl/books/081600038/

## Chapter2 Webシステムの基本、HTTP／HTTPSプロトコルを理解する

- HTTP/1.1はテキストベースプロトコル(ASCII文字), HTTP/2はバイナリベースプロトコル
  - それ以外にも1つのTCP接続で複数のリクエスト/レスポンスをやり取りできるようになったり、
  - ヘッダの圧縮等がより、高速で効率的な通信を実現できるようになった
  - [そろそろ知っておきたいHTTP/2の話 #http2 - Qiita](https://qiita.com/mogamin3/items/7698ee3336c70a482843#1%E3%81%A4%E3%81%AEtcp%E6%8E%A5%E7%B6%9A)
- HTTP/3で特徴的なのはQUICと呼ばれるUDPベースの伝送制御上でHTTP通信を行い、高効率や通信を実現している
  - 現在ALBではHTTP/3に対応しておらず、 HTTP/3に対応するにはCloudFrontを経由させる必要がある(?)
