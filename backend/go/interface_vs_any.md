# interface{} vs any

## はじめに

空のインターフェイスを表現するものとして、interface{}とanyがある<br>
両者の違いとどちらを使うべきかを調べた際のメモ

## interface{}

- Goの初期バージョンから存在する
- 全ての型はinterface{}を実装しているので任意の型を格納できる

## any

- Go1.18で追加されたinterface{}の型エイリアス
- 可読性向上のために導入された

## まとめ

anyを使うべし

## 参考

[Go 1.18 で interface{} の代わりに any が使えるようになる話](https://zenn.dev/syumai/articles/c6q5un1j0msim0aj0ca0)
