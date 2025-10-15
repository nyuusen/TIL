package main

import (
	"fmt"
)

func main() {
	var a []int
	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	b := a
	b = append(b, 5)
	c := a
	c = append(c, 6)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}
