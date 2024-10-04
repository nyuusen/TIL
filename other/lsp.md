# Language Server Protocol(LSP)

## はじめに

職場の方が語っていたので興味本位で調べてみる。

## Language Server Protocol(LSP)とは？

- テキストエディタやIDE（統合開発環境）とプログラミング言語のサーバー（言語サーバー）との間で通信を行うためのプロトコルのこと
- コード補完やシンタックスハイライト等の機能提供してくれる
- Microsoftが提唱した

## LSPの必要性

- エディタに関係なく一貫した開発体験を得られる
  - 言語ごとに専用のプラグインやサポートを作成するのは手間がかかる

## 技術的な構成

- エディタやIDEと言語サーバーがLSPを通じて通信する
- 言語サーバー
  - 実体はプログラミング言語環境
  - VSCodeではGoという拡張機能をインストールするとgoplsというGoの言語サーバーがローカルにインストールされるっぽい(?)

## VSCode + Goだと..？

- Goの言語サーバーの実装がgoplsである
- VSCodeのsettings.jsonの`"go.useLanguageServer": true`でLSP設定をONにできる
