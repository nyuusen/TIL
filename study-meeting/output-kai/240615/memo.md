## React v18 のHooks

- useId
  - HTMLタグのid属性を作成する
- useTransition
  - ない(useStateを使用する等)とレンダリング完了までブロッキングされる
    - レンダリングが遅い画面があるとUXが悪い
  - あるとスムーズに画面操作ができる
    - 裏側でレンダリングを走らせてくれるので、画面操作は可能
- useDeferredValue
  - 入力を待ってレンダリングを遅延させる

[Hooks API Reference – React](https://legacy.reactjs.org/docs/hooks-reference.html)

## Flutter

- path
  - パスを結合したりする時に使用する
  - nodeのpathモジュールと同じような感じかな？
- sqflite