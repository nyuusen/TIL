package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("exec middle.go")

	// fmt.Printf("return value: %d\n", squareSummation(2))
	// fmt.Printf("return value: %d\n", squareSummation(3))
	// fmt.Printf("return value: %d\n", squareSummation(367))

	// fmt.Printf("return value: %s\n", mergeString("abc", "def"))

	// fmt.Printf("return value: %s\n", swapPosition("abcd"))
	// fmt.Printf("return value: %s\n", swapPosition("abcde"))

	// fmt.Printf("return value: %d\n", product(3, 5))
	// fmt.Printf("return value: %d\n", product(5, -1))
	// fmt.Printf("return value: %d\n", product(-9, 8))
	// fmt.Printf("return value: %d\n", product(-10, -5))
	// fmt.Printf("return value: %d\n", product(0, 2))

	// fmt.Printf("return value: %s\n", commonPrefix("abcdefg", "abcxyz"))

	// fmt.Printf("return value: %d\n", numberOfDots(1))
	// fmt.Printf("return value: %d\n", numberOfDots(2))
	// fmt.Printf("return value: %d\n", numberOfDots(3))
	// fmt.Printf("return value: %d\n", numberOfDots(4))

	// fmt.Printf("return value: %d\n", totalSquareArea(1))
	// fmt.Printf("return value: %d\n", totalSquareArea(2))
	// fmt.Printf("return value: %d\n", totalSquareArea(3))

	// fmt.Printf("return value: %s\n", sheeps(3))

	// fmt.Printf("return value: %s\n", reverseString("abc"))
	// fmt.Printf("return value: %s\n", reverseString("recursion"))

	// fmt.Printf("return value: %d\n", countDivisibleByK(3, 2))
	// fmt.Printf("return value: %d\n", countDivisibleByK(30, 5))
	// fmt.Printf("return value: %d\n", countDivisibleByK(24, 2))

	// fmt.Printf("return value: %d\n", maximumPeople(12, 18))
	// fmt.Printf("return value: %d\n", maximumPeople(30, 242))

	// fmt.Printf("return value: %d\n", threeGCD(12, 18, 24))
	// fmt.Printf("return value: %d\n", threeGCD(30, 243, 91))

	// fmt.Printf("return value: %s\n", stringCompression("aabbb"))
	// fmt.Printf("return value: %s\n", stringCompression("ab"))

	fmt.Printf("return value: %d\n", maxBread(10, 2, 3))
}

func maxBread(money int32, price int32, sticker int32) int32 {
	// 購入できる最大のパンの個数
	if money == 0 || money < price {
		return 0
	}
	breads := money / price
	return breads + exchangeBreadsByStickers(breads, sticker)
}

func exchangeBreadsByStickers(own int32, threshold int32) int32 {
	if own < threshold {
		return 0
	}
	breads := own / threshold
	return breads + exchangeBreadsByStickers((own%threshold+breads), threshold)
}

func stringCompression(s string) string {
	// 文字列の圧縮
	// aaabbb -> a3b3
	if len(s) == 0 {
		return ""
	}
	target := s[0]
	count := 1
	for count < len(s) && s[count] == target {
		count++
	}
	result := string(s[0])
	if count > 1 {
		result = fmt.Sprintf(result+"%d", count)
	}
	return result + stringCompression(s[count:])
}

func threeGCD(x int32, y int32, z int32) int32 {
	// 3つの最大公約数
	fmt.Printf("x: %d, y:%d, z:%d\n", x, y, z)
	if y == 0 && z == 0 {
		return x
	}
	// まずはx,yの最大公約数を求める
	if x != 0 && y != 0 && z != 0 {
		if x > y {
			return threeGCD(y, x, z)
		}
		return threeGCD(x, y%x, z)
	}
	// x,yの最大公約数とzの最大公約数を求める
	if y == 0 {
		return threeGCD(x, z, y)
	}
	if x > y {
		return threeGCD(y, x, z)
	}
	return threeGCD(x, y%x, z)
}

func maximumPeople(x int32, y int32) int32 {
	// 最大公約数
	if y == 0 {
		return x
	}
	if x > y {
		return maximumPeople(y, x)
	}
	return maximumPeople(x, y%x)
}

func countDivisibleByK(n int32, k int32) int32 {
	// kで割り続ける
	// 28,2 -> 28/2=14 14/2=7 の2回
	rem := n % k
	if rem != 0 {
		return 0
	}
	return 1 + countDivisibleByK(n/k, k)
}

func reverseString(s string) string {
	// 文字列の逆表示
	sLength := len(s)
	if sLength == 0 {
		return ""
	}
	last := string(s[sLength-1])
	newStr := s[:sLength-1]
	return last + reverseString(newStr)
}

func sheeps(count int32) string {
	// 羊を数える
	if count == 0 {
		return ""
	}
	return sheeps(count-1) + strconv.Itoa(int(count)) + " sheep ~ "
}

func totalSquareArea(x int32) int32 {
	// 正方形の合計面積
	if x == 0 {
		return 0
	}
	return int32(math.Pow(float64(x), 3)) + totalSquareArea(x-1)
}

func numberOfDots(x int32) int32 {
	// パスカル
	// 1番目は1, 2番目は3、3番目は6、4番目は10、5番目は15
	if x == 0 {
		return 0
	}
	return x + numberOfDots(x-1)
}

func commonPrefix(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	if s1[0] != s2[0] {
		return ""
	}
	return string(s1[0]) + commonPrefix(s1[1:], s2[1:])
}

func product(x int32, y int32) int32 {
	// 5, 3 -> 15
	// 5, -1 -> -5
	// -10, -5 -> -50
	fmt.Printf("x: %d, y: %d\n", x, y)
	if y == 0 {
		return 0
	}
	if y < 0 {
		return -x + product(x, y+1)
	}
	return x + product(x, y-1)
}

func swapPosition(s string) string {
	// abcd -> badc
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	return s[1:2] + s[0:1] + swapPosition(s[2:])
}

func mergeString(s1 string, s2 string) string {
	// abc, def -> adbecf
	if len(s1) == 0 && len(s2) == 0 {
		return ""
	}
	return s1[:1] + s2[:1] + mergeString(s1[1:], s2[1:])
}

func squareSummation(n int32) int32 {
	// 3の2乗の階乗＝2の2乗＋3の2乗
	// S(n) = S(n-1) + S
	if n == 0 {
		return 0
	}
	return int32(math.Pow(float64(n), 2)) + squareSummation(n-1)
}

func factorial(n int32) int64 {
	if n == 0 {
		return int64(1)
	}
	return factorial(n-1) * int64(n)
}

func lenString(str string) int32 {
	if str == "" {
		return int32(0)
	}

	return 1 + lenString(str[1:])
}

func recursiveAddition(m int32, n int32) int32 {
	if n == 0 {
		return m
	}

	return recursiveAddition(m+1, n-1)
}

func infectedPeople(day int32) int32 {
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
