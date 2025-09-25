package main

import "fmt"

func main() {
	str := "ABCDEあいうえお"

	fmt.Println(&str)

	// 文字列に対するインデックスアクセスはバイトが返る
	for i := 0; i < len(str); i++ {
		byte := str[i]
		fmt.Println(i, byte)
	}

	// 文字列に対するfor-rangeはruneが返る
	for i, rune := range str {
		fmt.Println(i, rune)
	}

	fmt.Println(firstChar(str))
}

// 引数の文字列の最初の文字を返す関数
func firstChar(s string) string {
	for _, r := range s {
		return string(r)
	}
	return ""
}
