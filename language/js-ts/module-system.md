# モジュールシステム

- ES ModulesとCommonJSの2種類がある
- Node.jsではデフォルトがCommonJSであり、Next.jsではデフォルトがES Modulesとなっている
- 設定箇所
  1. tsconfig.jsonのcompilerOption.module
  2. package.jsonのtypeフィールド(moduleかcommonjsを設定)
- 上記設定の違い
  - tsconfig.jsonでの設定
    - TSがトランスパイル時に生成するJSコードのモジュール形式をしている(トランスパイル後のJSコードの指定されたモジュールシステムになる)
    - TSコンパイラが、import/export構文をどのモジュール形式に変換するかを決定する
  - package.jsonでの設定
    - Node.jsが実行時のJSファイルをどのモジュールシステム形式として扱うかを指定する
    - .jsファイルをESMとして扱うか、CJSとして扱うか
  - 双方の設定を合わせることで、トランスパイル時・実行時のモジュール形式を統一させることができる
- ファイル拡張子
  - js: プロジェクトの設定次第
  - mjs: 常にESMとして扱われる
  - cjs: 常にCJSとして扱われる