package main

import "fmt"

func main() {

	var (
		i int
		m int
		n int
	)

	fmt.Scan(&i, &m)
	for i <= m {
		n += i
		i++
	}
	fmt.Println(n)
}
