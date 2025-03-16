## AWSで実装する場合

以下のどちらかで実装が可能

- WAF
- CloudFront Functions

それぞれを比較してみる（Lambda@Edgeも可能だが、CFFできるなら...と思い、選択肢から外した）

### WAF

- WebACLのルールグループにヘッダに関するStatementを記述する形で実装する
- Block時のカスタムレスポンスとして、ヘッダにWWW-Authenticateを返すことも可能

### CloudFront Functions

- https://iret.media/95931 みたいな感じで手軽に実装可能

### どっちが良いか？

- 手軽さと「CloudFront Functions」だと思う
- ただし、既にCFのViewer RequestにCFFを割り当てている場合やWAFを既に利用している環境ではWAFの利用もアリなのではと思った
  - 実際、私が業務で対応した際は、IP制限を既にWAFを利用しており、IP制限外からでもBasic認証情報が検証できればリクエストを通したいといった要件だったので、WAFを採用した（したい）