package main

import "fmt"

func replase(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	replase(&a, &b)
}
