package main

import "fmt"

func point(x1 *int, x2 *int) {
	result := *x1 * *x2
	fmt.Println(result)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	point(&a, &b)
}
