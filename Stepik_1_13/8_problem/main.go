package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	fmt.Scan(&x)
	min := x
	number := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x < min {
			min = x
			number = 0
		}
		if x == min {
			number++
		}
	}
	fmt.Println(number)
}
