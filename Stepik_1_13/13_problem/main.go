package main

import (
	"fmt"
	"log"
)

func fib(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}
	return a
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		if n == fib(i) {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
