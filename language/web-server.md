# Webサーバー

## 概要

PHP/LaravelでWebアプリケーションを開発する場合、ApacheやNginx等のWebサーバーが必要になる一方で、\
Node.jsの場合は不要である。この違いが気になったのと、そもそもWebサーバーって何をやってくれているのか気になるので調べてみる。

## PHP/LaravelにおけるApacheやNginxの必要性

- リクエストのルーティング
  - pacheやNginxはクライアント（ブラウザなど）からのHTTPリクエストを受け取り、リクエストされたURLとHTTPメソッドに基づいて、PHPアプリケーションを振り分けている
- PHPエンジンとの連携
  - ApacheやNginxはリクエストをPHPエンジンに渡している
  - Apacheではmod_phpモジュール、Nginxでは、PHP-FPM（FastCGI Process Manager）を使用している
- 静的ファイルの配信
- 負荷分散やリバースプロキシ機能

## Node.jsにおけるWebサーバーが不要である理由

- 理由：内臓のHTTPサーバー機能を持っており、Node.js自体がHTTPリクエストを直接処理可能であるため
  - httpやexpressなどのモジュールを使用して直接HTTPリクエストを受け取り、レスポンスを返す

## PHP/LaravelとECS on Fargateにデプロイしたい時はどうするの？

- HTTPリクエスト処理を行うNginx/ApacheコンテナとPHP-FPM（PHP FastCGI Process Manager）コンテナに分割する
  - 同居も可能ではあるが、リソース効率やスケールの柔軟性を考慮すると分割するのが推奨らしい
- (Nginxを仮定し、前面にALBを配置したケースを例)具体的には、
  - ALBのターゲットをNginxコンテナの80番ポートに転送する
  - NginxコンテナからPHPリクエストをFastCGIでPHP-FPMコンテナに転送するように設定(nginx.conf)
