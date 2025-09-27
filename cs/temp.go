package main

import "fmt"

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
