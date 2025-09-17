## React Hooks(v18)　by tanaka-san

- useId
  - HTMLタグのid属性に付与する値を生成する
  - JSXのループ時のkeyに使用するものではない
- useTransition
  - レンダリング処理を裏側で実行するための関数(startTransition)を提供する
  - 具体的には、startTransition内に重い状態更新処理等を記述することにより、画面がブロックされることを回避する
  - [useTransitionについて理解する](https://zenn.dev/hakoten/articles/e7ea977e00b4f8)
- useDeferredValue
  - レンダリングを遅延させる
  - useDeferredValueが更新されたら、裏側でレンダリングを試行する
  - 頻度が高いレンダリングを遅延させ、パフォーマンス向上に有用

[Hooks API Reference – React](https://legacy.reactjs.org/docs/hooks-reference.html)

## Flutter by koike-san

- path
  - パスを結合したりする時に使用する
  - nodeのpathモジュールと同じような感じ？
- sqflite
  - ローカルストレージにあるファイルをRDBライクに操作できる
  - 端末からアプリを削除すると、データも削除される？
