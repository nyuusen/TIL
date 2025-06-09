# AWS Transfer Family For SFTP

## やりたいこと

- バッチ処理で出力したファイルをS3に置く
- そのファイルをSFTPで任意のPCからダウンロードする

## 登場人物

- バッチ処理環境
- S3バケット
- Transfer Family For SFTP（SFTPサーバー）
- SFTPクライアント

## 必要な設定

- S3バケット作成
  - バケットポリシー定義
- Transfer Family For SFTP作成
  - バックエンドはS3を設定
- SFTPユーザーの作成
  - ユーザーのSSH公開鍵を設定（作成する必要あり）
  - ホームディレクトリ: 転送を行うS3バケットのパス名
- SFTPクライアント設定（WinSCPとか）
  - 接続先
  - ユーザー名（上記で設定したユーザー名）
  - パスワード（秘密鍵）