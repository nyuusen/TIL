package beginner

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	howMuchIsYourDebt(2)
	howMuchIsYourDebt(3)
	howMuchIsYourDebt(5)
	howMuchIsYourDebt(10)
}

func weekly7DaysSales(ticketPrice int32) int32 {
	// float64にキャストして計算する
	diff := (250.0 - float64(ticketPrice)) / 10.0
	// int32にキャストして返す
	return int32(150000 + (diff * 7000))
}

func vacationRental(people int32, day int32) int32 {

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

	// 年利20％で10000円を借りた場合、year年後にいくら返済する必要があるかを計算して返す
	principal := 10000
	rate := 1.2
	interest := math.Pow(rate, float64(year))
	sum := float64(principal) * interest

	fmt.Printf("year: %d, principal: %d, rate: %f, interest: %f sum: %f \n", year, principal, rate, interest, sum)

	return int32(math.Floor(sum))
}

func isRationalNumber(number int32) bool {

	res := math.Sqrt(float64(number))
	return res == float64(int32(res))
}

func toLowerCase(stringInput string) string {

	return strings.ToLower(stringInput)
}

func insertUnderscoreAt(s string, i int32) string {

	if len(s)-1 < int(i) {
		return s
	}
	return s[:i] + "_" + s[i:]
}

func lastFourHint(stringInput string) string {

	length := len(stringInput)
	if length < 6 {
		return "There is no Hint"
	}

	return "Hint is:" + stringInput[length-4:]
}

func isValidEmail(email string) bool {

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

func middleSubstring(s string) string {
	runes := []rune(s)
	n := len(runes)

	if n <= 2 {
		if n == 0 {
			return ""
		}
		return string(runes[0])
	}

	var start, length int
	if n%2 == 0 {
		// 偶数
		length = n / 2
		start = (n - length + 1) / 2
	} else {
		// 奇数
		length = n / 2
		start = (length + 1) / 2
	}

	return string(runes[start : start+length])
}

func calculateLocation(latitude float64, longitude float64) string {

	var ret1 string
	var ret2 string
	switch {
	case latitude == 0:
		ret1 = "equator"
	case latitude > 0:
		ret1 = "north"
	case latitude < 0:
		ret1 = "south"
	}

	switch {
	case longitude == 0:
		ret2 = "prime meridian"
	case longitude > 0:
		ret2 = "east"
	case longitude < 0:
		ret2 = "west"
	}

	return ret1 + "/" + ret2
}
