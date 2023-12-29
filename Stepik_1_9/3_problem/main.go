package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for x > 9 {
		x = x / 10
	}
	fmt.Print(x)
}
