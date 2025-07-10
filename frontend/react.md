# React

## はじめに

以下の資料を見てためになったことをメモ。

- [React 研修 (2024) - Speaker Deck](https://speakerdeck.com/recruitengineers/react-yan-xiu-2024)
- [Reactの内部構造を知っておく (React Tokyo #6 - @calloc134) - Speaker Deck](https://speakerdeck.com/calloc134/reactnonei-bu-gou-zao-wozhi-tuteoku-react-tokyo-number-6-at-calloc134)
- [JSX - 歴史を振り返り、⾯⽩がって、エモくなろう - Speaker Deck](https://speakerdeck.com/pal4de/jsx-li-shi-wozhen-rifan-ri-gatute-emokunarou)
- [Next.jsの考え方](https://zenn.dev/akfm/books/nextjs-basic-principle)

## 命令的から宣言的へ

- React登場前に流行っていたjQueryは命令的に記述する
  - 命令的：状態とUIの関係を命令的に書くということ
    - どういうことかというと、ボタンを押すとCountという表示が＋1されるUIがあるとしたら、ボタンを押す→変数を＋1する→表示しているCountの中身を1にする（元は0の状態という前提）
- 一方、Reactは宣言的に記述する
  - 状態とUIの関係はReactが賢く管理してくれる
  - 実装されたコードから手順書感が消える
  - ロジックとUIの構造が分離され、俯瞰しやすい