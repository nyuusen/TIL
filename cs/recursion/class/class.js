// オブジェクトの問題は、記述したテストケースが実行されることで出力が正しいかを判断します。
// 通常のコーディング問題のように入力が自動で与えられるわけではありません。
// まずは、問題で指定されているDogクラスを作成してみましょう。

// なお、右上の各ボタンは、下記の通りです。
// 実行：記述したコードを実行します。デバッグに使用します。
// テスト：提出時と同じ環境で実行されます。提出前の最終チェックに使用します。
// 提出：問題文中のテストケースが実行され、出力が正しいか確認されます。

// ① Dogクラスを作成してみましょう。
class Dog {
    // 問題文の指定した条件に従って、コンストラクタやメソッドを記述します。
    
    constructor(name, size, age) {
      this.name = name
      this.size = size
      this.age = age
    }

    bark() {
      if (this.size >= 50) {
        return "Wooof! Woof!"
      }
      if (this.size >= 20) {
        return "Ruff! Ruff!"
      }
      return "Yip! Yip!"
    }

    calcHumanAge() {
      return 12 + (this.age - 1) * 7
    }    
}

// ② 問題文にある一つ目のテストケースを記述しましょう
let goldenRetriever = new Dog("Golden Retriever", 60, 10);
// コンソール出力することでテストケースが正しい値であるかが判定されます。
console.log(goldenRetriever.bark());
console.log(goldenRetriever.calcHumanAge());

// ③ 上記を参考に他のテストケースを作成しましょう。
// 問題文中のテストケースを全て満たすことで合格します。

// ④ テストボタンを押して正しく出力されているかを確認しましょう。
// メソッド名、文字列のスペースなどは注意深く確認しましょう。

// ⑤ テストが確認できたら、提出ボタンを押して完了です。
