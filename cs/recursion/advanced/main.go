package main

import "fmt"

func main() {
	compareLength := func(s1, s2 string) bool { return len(s1) >= len(s2) }
	compareAsciiTotal := func(s1, s2 string) bool { return sumAscii(s1) >= sumAscii(s2) }

	fmt.Println(maxByCriteria(compareLength, []string{"apple", "yumberry", "grape", "banana", "mandarin"}))
	fmt.Println(maxByCriteria(compareLength, []string{"zoomzoom", "choochoo", "beepbeep", "ahhhahhh"}))
	fmt.Println(maxByCriteria(compareAsciiTotal, []string{"apple", "yumberry", "grape", "banana", "mandarin"}))
	fmt.Println(maxByCriteria(compareAsciiTotal, []string{"zoom", "choochoo", "beepbeep", "ahhhahhh"}))
}

func maxByCriteria(f func(s1, s2 string) bool, ary []string) string {
	target := ary[0]
	for _, v := range ary {
		if res := f(v, target); res {
			target = v
		}
	}
	return target
}

func sumAscii(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}
