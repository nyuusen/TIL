package main

import (
	"fmt"
)

func main() {
	fmt.Println("exec middle.go")

	fmt.Printf("answer: %d\n", lenString("Hello"))
	fmt.Printf("answer: %d\n", lenString("pikachu"))
}

func lenString(str string) int32 {
	// 関数を完成させてください
	if str == "" {
		return int32(0)
	}

	return 1 + lenString(str[1:])
}

func recursiveAddition(m int32, n int32) int32 {
	// 関数を完成させてください
	if n == 0 {
		return m
	}

	return recursiveAddition(m+1, n-1)
}

func infectedPeople(day int32) int32 {
	// 関数を完成させてください
	if day == 1 {
		return int32(2)
	}
	return infectedPeople(day-1) * 2
}

func multiplyOf7(n int) int {
	fmt.Printf("n: %d\n", n)
	if n == 0 {
		return 0
	}
	return multiplyOf7(n-1) + 7
}
