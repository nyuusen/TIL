## JavaScriptのDOMContentLoadedの話 by tanaka-san

- DOMContentLoadedだとブラウザバックした時に値が拾えない
  - 値を拾う＝キャッシュから値を取得するタイミングがDOMContentLoadedより遅いタイミングで実行されるため
  - pageShowイベントを実行することで解決

## 個人開発でリアルタイム編集 by koike-san

- リアルタイム編集機能をFirebaseとFlutterのStreamBuilderで実装
  - StreamBuilderを使うことでデータベース側の変更を検知できる
- リアルタイム編集機能は考慮ポイントが多い...

## IPv4 VPC CIDRブロック by irikawa

- 話した内容: https://github.com/nyuusen/TIL/blob/main/infra/aws/vpc-ipv4-cidr-block.md
  - Q. ブロックサイズを小さくする(ホスト部に割り当てられるIPアドレス多くする)ことによるデメリットは？
    - IPアドレスの無駄遣いにより、VPCピアリング接続時のアドレス空間重複の可能性が高まる
    - 未使用のIPアドレスが多く残り、非効率であり、管理の複雑性も増す
    - セキュリティグループやネットワークACLでの管理が複雑になり、設定ミスやアクセス制御の甘さが生じる可能性がある