package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

	// fmt.Printf("return value: %d\n", maxBread(10, 2, 3))

	// fmt.Printf("return value: %d\n", intSquareRoot(12))
	// fmt.Printf("return value: %d\n", intSquareRoot(35))

	// fmt.Printf("return value: %t\n", recursiveIsPrime(1))
	// fmt.Printf("return value: %t\n", recursiveIsPrime(2))
	// fmt.Printf("return value: %t\n", recursiveIsPrime(25))
	// fmt.Printf("return value: %t\n", recursiveIsPrime(29))

	// fmt.Printf("return value: %d\n", countSquare(28, 32))

	// fmt.Printf("return value: %d\n", splitAndAdd(234))

	// fmt.Printf("return value: %d\n", multipleOfTwoTotal(3))
	// fmt.Printf("return value: %d\n", multipleOfTwoTotal(2))

	// fmt.Printf("return value: %d\n", fibonacciNumber(5))
	// fmt.Printf("return value: %d\n", fibonacciNumber(8))
	// fmt.Printf("return value: %d\n", fibonacciNumber(10))

	// fmt.Printf("return value: %d\n", numberOfWay(10))

	// fmt.Printf("return value: %d\n", towerOfHanoi(3))
	// fmt.Printf("return value: %d\n", towerOfHanoi(5))

	// fmt.Printf("return value: %d\n", divideBy3Count(10))

	fmt.Printf("return value: %d\n", howLongToReachFundGoal(5421, 10421, 5))
	fmt.Printf("return value: %d\n", howLongToReachFundGoal(600, 10400, 7))
	fmt.Printf("return value: %d\n", howLongToReachFundGoal(32555,5200000,12))
	fmt.Printf("return value: %d\n", howLongToReachFundGoal(650,35000,65))

}

func howLongToReachFundGoal(capitalMoney int32, goalMoney int32, interest int32) int32 {
	// 何年投資すれば土地が買えるか
	// goalMoney:土地の現在価格、interest:投資対象の年利、capitalMoney:投資額
	// 土地の価格は経過年数が偶数の場合は2%上昇,奇数の場合は3%上昇
	// 投資額と土地の価格をそれぞれ計算する必要がある
	// 投資額が土地の価格を上回れば良い
	var year int32 = 1
	return howLongToReachFundGoalHelper(float64(capitalMoney), float64(goalMoney), interest, year)
}

func howLongToReachFundGoalHelper(capitalMoney float64, goalMoney float64, interest int32, year int32) int32 {
	if year >= 80 {
		return 80
	}

	if capitalMoney >= goalMoney {
		return year
	}
	// 年利5％: 元金*1.05
	cm := float64(capitalMoney) * (float64(interest)/100 + 1)
	gm := calculateGoalMoney(goalMoney, year)
	year += 1
	return howLongToReachFundGoalHelper(cm, gm, interest, year)

}

func calculateGoalMoney(price float64, year int32) float64 {
	if year%2 == 0 {
		return float64(price) * 1.02
	}
	return float64(price) * 1.03
}

func divisor(number int32) string {
	var numbers []string
	var n int32 = 1
	for n <= number {
		dh := divisorHelper(number, n)
		if dh != 0 {
			numbers = append(numbers, strconv.Itoa(int(n)))
		}
		n++
	}
	return strings.Join(numbers, "-")
}

func divisorHelper(original, number int32) int32 {
	if number == 1 {
		return 1
	}
	if original%number != 0 {
		return 0
	}
	return number
}

func divideBy3Count(n int32) int32 {
	if n < 3 {
		return 0
	}
	return 1 + divideBy3Count(n/3)
}

func towerOfHanoi(discs int32) int32 {
	// ハノイの塔
	// f(3)=f(2-1)+1+f(2-1)
	if discs == 1 {
		return 1
	}
	return towerOfHanoi(discs-1) + 1 + towerOfHanoi(discs-1)
}

func numberOfWay(x int32) int32 {
	// xになるまでの組み合わせ
	// numberOfWay(6)=numberOfWay(5)+numberOfWay(4)
	// numberOfWay(n)=numberOfWay(n-1)+numberOfWay(n-2)
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x == 2 {
		return 2
	}
	return numberOfWay(x-1) + numberOfWay(x-2)
}

func fibonacciNumber(n int32) int32 {
	// フィボナッチ数列
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciNumber(n-1) + fibonacciNumber(n-2)
}

func multipleOfTwoTotal(n int32) int32 {
	// 2の倍数の総和の総和
	// nが3のとき20,nが2のとき8
	// multipleOfTwoTotal(3) = multipleOfTwoTotalHelper(3) + multipleOfTwoTotal(2)
	if n == 0 {
		return 0
	}
	return multipleOfTwoTotalHelper(n) + multipleOfTwoTotal(n-1)
}

func multipleOfTwoTotalHelper(n int32) int32 {
	// 2の倍数の総和
	// nが3のとき12,nが2のとき6
	// multipleOfTwoTotalHelper(3) = 2 * n +  multipleOfTwoTotalHelper(2)
	if n == 0 {
		return 0
	}
	return (n * 2) + multipleOfTwoTotalHelper(n-1)
}

func splitAndAdd(digits int32) int32 {
	// 桁ごとに抽出して足し算
	// 123 -> 1+2+3
	// 0以下の場合は0を返す(ベースケース)
	if digits <= 0 {
		return 0
	}
	// 10未満(1桁)の場合はそのまま返す
	if digits < 10 {
		return digits
	}
	// 10で割った数の余り(123だったら3)と割った数(12)を再帰
	return digits%10 + splitAndAdd(digits/10)
}

func countSquare(x int32, y int32) int32 {
	// x,yを辺とする長方形から、できるだけ大きい正方形がいくつ作れるか
	// まずは最大公約数を求める
	side := gcd(x, y)
	// 長方形面積から正方形面積を求める
	return (x * y) / (side * side)
}

func gcd(x int32, y int32) int32 {
	if x > y {
		return gcd(y, x)
	}
	if x == 0 {
		return y
	}
	return gcd(x, y%x)
}

func recursiveIsPrime(n int32) bool {
	// 素数判定
	return recursiveIsPrimeHelper(n, n-1)
}

func recursiveIsPrimeHelper(n int32, n2 int32) bool {
	// 素数判定
	// 素数である条件：1以外ではない数字で割り切れる
	if n2 <= 0 {
		return false
	}
	if n2 == 1 {
		return true
	}
	if n%n2 == 0 {
		return false
	}
	return recursiveIsPrimeHelper(n, n2-1)
}

func intSquareRoot(n int32) int32 {
	// 整数上の平方根
	// 自然数nの平方根の整数部分を計算する
	return intSquareRootHelper(n, n)
}

func intSquareRootHelper(target int32, n int32) int32 {
	result := n * n
	if target >= result {
		return n
	}
	return intSquareRootHelper(target, n-1)
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
