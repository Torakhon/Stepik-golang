package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	for i := 0; i < 3; i++ {
		sum += a % 10
		a = a / 10
	}
	fmt.Println(sum)
}
