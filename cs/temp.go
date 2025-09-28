package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(weekly7DaysSales(260))
	fmt.Println(weekly7DaysSales(255))
}

func weekly7DaysSales(ticketPrice int32) int32 {
	// float64にキャストして計算する
	diff := (250.0 - float64(ticketPrice)) / 10.0
	// int32にキャストして返す
	return int32(150000 + (diff * 7000))
}

func vacationRental(people int32, day int32) int32 {
	// 関数を完成させてください
	var stayPrice int32
	if day <= 3 {
		stayPrice = 80
	} else if day < 10 {
		stayPrice = 60
	} else {
		stayPrice = 50
	}

	stayPriceSum := stayPrice * day * people
	cleanPrice := float32(stayPriceSum) * 0.12
	sum := (float32(stayPriceSum) + cleanPrice) * 1.08

	// sumを小数点以下切り捨てしてint32にキャストして返す
	return int32(math.Floor(float64(sum)))
}

func howMuchIsYourDebt(year int32) int32 {
	// 関数を完成させてください
	// 年利20％で10000円を借りた場合、year年後にいくら返済する必要があるかを計算して返す
	debt := float64(10000) * math.Pow(1.2, float64(year))
	return int32(math.Floor(debt))
}

func isRationalNumber(number int32) bool {
	// 関数を完成させてください
	res := math.Sqrt(float64(number))
	return res == float64(int32(res))
}

func toLowerCase(stringInput string) string {
	// 関数を完成させてください
	return strings.ToLower(stringInput)
}

func insertUnderscoreAt(s string, i int32) string {
	// 関数を完成させてください
	if len(s)-1 < int(i) {
		return s
	}
	return s[:i] + "_" + s[i:]
}

func lastFourHint(stringInput string) string {
	// 関数を完成させてください
	length := len(stringInput)
	if length < 6 {
		return "There is no Hint"
	}

	return "Hint is:" + stringInput[length-4:]
}

func isValidEmail(email string) bool{
    // 関数を完成させてください
		if strings.HasPrefix(email, "@") {
			return false
		}

		if strings.Contains(email, " ") {
			return false
		}

		if strings.Count(email, "@") != 1 {
			return false
		}

		if strings.Index(email, "@") > strings.LastIndex(email, ".") {
			return false
		}

		return true
}