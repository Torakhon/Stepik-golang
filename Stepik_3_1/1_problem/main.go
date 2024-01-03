package main

import (
	"fmt"
)

func work(x int) int {
	if x > 3 {
		return x + 1
	} else {
		return x - 1
	}
}

func main() {
	m := make(map[int]int)
	var x int
	for i := 0; i < 10; i++ {
		fmt.Scan(&x)
		if el, ok := m[x]; ok {
			fmt.Print(el, " ")
		} else {
			m[x] = work(x)
			fmt.Print(m[x], " ")
		}

	}
}
