# HomebrewでGoバージョンを切り替える

## 概要

今まではgoenvとかを使ってローカルのGoバージョンを管理していたが、PC切り替えたタイミングでシンプルにHomebrewで管理してみることにした。\
Homebrewだと`brew update`で新しいバージョンが入るので良いが、バージョン切り替えに手こずったのでその方法を残しておく。

## やり方

1.23.4から1.23.6に上げる方法\
※ここではあえて、`brew update`ではなくgoのインストールから手動で行うで紹介

```sh
// 存在を確認
brew search go@1.23.6

// インストール
brew install "go@1.23"

// 現在リンクされているgoバージョンを解除
brew unlink go

// 1.23に設定
brew link go@1.23

// 出力されたパス通すコマンド実行
echo 'export PATH="..."' >> ...
source ~/.zshrc

go version
>> go version go1.23.6 darwin/arm64
```

注意点としては、brew switchを紹介している古い記事が結構あるが、switchコマンドは廃止されており、使用できなかった...
