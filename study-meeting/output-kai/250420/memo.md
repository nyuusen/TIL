# CSS周り by tanaka-san

- CSS
  - ブラウザがCSSを読み込むだけ
- CSS in JS
  - JSで定義したCSSを解析
  - 動的なCSSは実装はしやすい反面、JSの実行が二重となるのでパフォーマンス上の課題となる
  - Styled-Componentは、CSS in JSライブラリの1つ
- ゼロランタイムCSS
  - ビルド時に静的なCSSファイルを生成する
  - 型定義ができる
  - 動的実装が簡易的
  - CSS modules/Linaria/Vanilla-extract