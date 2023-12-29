package main

import "fmt"

func main() {
	var x, p, y int

	fmt.Scan(&x, &p, &y)

	years := 0

	for x < y {
		increase := x * p / 100
		x += increase
		years++
	}

	fmt.Println(years)
}
