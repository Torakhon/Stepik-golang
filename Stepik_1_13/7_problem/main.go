package main

import "fmt"

func main() {
	var n, number int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		if fmt.Scan(&x); x == 0 {
			number++
		}
	}

	fmt.Println(number)
}
