### 秘密鍵周りの処理を深掘り

- 題材は以下のgosnowflakeのプライベート関数で定義されている秘密鍵パース処理の一部を改変したもの

  ```go
  // parsePrivateKey PEM形式の秘密鍵ファイルを解析してrsa.PrivateKeyを返す
  func parsePrivateKey(key string) (*rsa.PrivateKey, error) {
  	block, _ := pem.Decode([]byte(key))
  	if block == nil {
  		return nil, errors.New("PEMブロック解析失敗")
  	}
  	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
  	if err != nil {
  		return nil, err
  	}
  	pk, ok := privateKey.(*rsa.PrivateKey)
  	if !ok {
  		return nil, errors.New("rsa.PrivateKeyへの型アサーション失敗")
  	}
  	return pk, nil
  }
  ```

- `block, _ := pem.Decode([]byte(key))`
  - PEM(Privacy Enhanced Mail)形式は、`-----BEGIN...-----` というヘッダーで始まり、末尾が `-----END...-----` で終わる
  - これらの境界の中にある値をBase64デコードする
    - つまりPEMはBase64エンコードされている値
  - PEMをデコードすることで、生のバイナリ(DER:Distinguished Encoding Rules形式)を取り出す
  - なぜBase64エンコードしているか？
    - 異なるシステム間（メール、HTTP、ファイルシステム）でバイナリを安全に転送するため（テキストによるカプセル化を行っている）
    - バイナリデータは、00からFFまでの0〜255の全ての値を取るが、この中には通信プロトコルやOSにとって特別の意味を持つ制御文字（NULL文字や改行コード,EOFなど）が含まれている
    - そこでバイナリを安全な印刷可能な64種類の文字だけに変換するのがBase64
- `privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)`
  - 取り出された値はASN.1 (Abstract Syntax Notation One)という形式のバイナリ（どのデータが何を表すか、アルゴリズムの種類や鍵の数値などをバイナリで定義する標準規格）
  - `PKCS#8`は、「秘密鍵をどう並べるか」というフォーマットの規格で、ここでの処理でバイナリをスキャンして「これはRSAアルゴリズムの鍵だ」と特定し、適切な内部構造へマッピングしている
- `pk, ok := privateKey.(*rsa.PrivateKey)`
  - privateKeyはany型なので型アサーションを行っている

| **層**      | **形式**                               | **役割**                                                                |
| ----------- | -------------------------------------- | ----------------------------------------------------------------------- |
| **Layer 1** | **PEM (Privacy Enhanced Mail)**        | Base64エンコードされたテキスト形式。人間が読める「封筒」の役割。        |
| **Layer 2** | **DER (Distinguished Encoding Rules)** | ASN.1規格に基づくバイナリ形式。コンピュータが解析する「中身」。         |
| **Layer 3** | **rsa.PrivateKey (Struct)**            | メモリ上の構造体。$e, d, n, p, q$ といったRSAの数学的パラメータを保持。 |

- PEMとは？
  - バイナリデータをテキストとして安全に扱うための『標準的な封筒』の規格
  - 以下のような構造になっている
    - Header
      - BEGINとデータの種類を表す
      - 例：`----------BEGIN RSA PRIVATE KEY----------`
    - Body
      - バイナリをBase64エンコードした値
      - 1行あたり64文字で改行
    - Footer
      - データの終わり（`-----END ...-----`）
  - PEMのメリット
    - バイナリだけだとそれが何かわからないが、PEMだとヘッダーを見たらそれが秘密鍵なのか公開鍵なのか証明書なのかが判別できる
    - Base64エンコードすることで、データの完全性を維持できる
