# ハイドレーション

ハイドレーションの必要性がわからなかったので調べた。

## ハイドレーションとは？

> React によってサーバーサイドレンダリングされたページは単純にクライアントに返却するだけでは JavaScript のイベントを受け取れるインタラクティブな状態にはなりません。onClick プロパティなどで渡されたイベントリスナーを DOM に登録する必要があります。このイベントリスナーを DOM に登録する処理を Hydrate と呼びます。

## なぜ必要？

- SSRされたHTMLだけでは、イベントリスナー（クリックなど）がバインドされないため
  - 素のJSで、document.querySelector('button').addEventListener()のように書けばイベントは効く
  - が、Reactのような仮想DOMを使うフレームワークでは不十分

## 流れ

- SSR（サーバー側でHTML生成し、ブラウザに返却する）
- ブラウザは、HTMLとともにJavaScriptファイルを読み込み実行（このJavaScriptファイルの中に状態管理やイベントバインドに関するロジックが含まれている）
- JavaScript(React)のハイドレーション関連の処理が実行され、SSRされたDOMとReactの仮想DOMを突き合わせて、イベント・状態等の機能を復元する

## React Server Componentsというアプローチ

- RSCには、Server ComponentsとClient Componentsがある
- Next.js AppRouterではRSCがデフォルト
- Server Components
  - ビルド時にレンダリングされたものをクライアントに送るので、ハイドレーションが不要
  - つまり、クライアントはサーバーが出力したHTMLを表示するだけ
  - Server ComponentsとSSRとの違い
    - SSRは、ページ全体をサーバーサイドでレンダリングし、完全なHTMLを返却する
    - SCは、特定のコンポーネントをサーバーサイドでレンダリングし、他の部分はクライアントサイドでレンダリングする
- Client Components
  - useStateのようなインタラクティブなAPIを使用するにはClient Componentsを使用する
  - `use client`ディレクティブを追加することでClient Componentsとして扱える
- [サーバコンポーネント – React](https://ja.react.dev/reference/rsc/server-components)
