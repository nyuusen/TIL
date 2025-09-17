# 関数型にインターフェイスを実装する

## はじめに

Goの特徴として、構造体にメソッドを定義する以外に関数型にもインターフェイスを実装できる。

net/httpパッケージを例に調べてみたのでメモ。

## 関数型にインターフェイスを実装とは？をnet/httpパッケージを例に理解する

前提として、今回題材にするnet/httpパッケージのhttp.Handlerインターフェイスについて理解しておく。

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Handlerは、ServeHTTPというメソッドを定義しているインターフェイスである。\
中身としては、Handlerという名の通り、HTTPリクエストを受けて、HTTPレスポンスを返す処理が記述されるものである。

本題に入るが、まずは違いをわかりやすくするために、通常の(よくある)方法である「構造体とメソッドを使ってインターフェイスを実装する」の実装例を書いてみる。

```go
package main

import (
    "fmt"
    "net/http"
)

// 構造体定義
type MyHandler struct{}

// 構造体メソッドでインターフェイスを実装
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    handler := MyHandler{}
    http.Handle("/hello", handler)
    http.ListenAndServe(":8080", nil)
}
```

MyHandlerという構造体に対し、Handlerインターフェイスを満たすServeHTTPというメソッドの実装を行なっている。\
その後、main関数内でMyHandlerインスタンスを生成し、http.Handleに渡すことで、パスに対してハンドラーを登録している。

次に「関数型を使ったインターフェイスを実装する」の実装例を書いてみる。

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

helloHandlerというfunc(http.ResponseWriter, *http.Request)シグネチャの関数を実装する。　\
...あれ、ハンドラーを定義するHandlerインターフェイスとハンドラーを登録するHandleはどこにいった...？

今回登場したhttp.HandlerFuncは、以下のように定義されている。

```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}
```

HandlerFuncという関数型に対し、ServeHTTPメソッドが実装されている。\
つまり、以下のhelloHandlerの実装時点＝HandlerFunc型の実装時点で自動的にServeHTTPというメソッドの実装が行われている。

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}
```

HandlerFuncを使用してハンドラーを実装した場合は、HandleFuncでハンドラーを登録する。\
両者の違いとしては、Handleの第2引数にはHandlerインターフェイスを実装しているハンドラーを登録するものであり、\
HandlerFuncの第2引数には、単純な関数をハンドラーとして登録するものである。

```go
func Handle(pattern string, handler Handler)

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

## 整理すると

**ハンドラーの登録**

Handle or HandleFunc

**ハンドラーの実装**

Handler or HandlerFunc

特段理由がなければ、短く書ける後者の方を選択するケースが多いと思われる。
HandlerFuncを使うとServeHTTPを実装する(空の)構造体をわざわざ作らなくて良いのがポイント。

## 最後に

あれどっちがどっちだっけってよくなってたけど、今回まとめたから流石に覚えた（多分)

## 参考

[http package - net/http - Go Packages](https://pkg.go.dev/net/http#Handler)
[Goのhttp.Handlerやhttp.HandlerFuncをちゃんと理解する - oinume journal](https://journal.lampetty.net/entry/understanding-http-handler-in-go)
