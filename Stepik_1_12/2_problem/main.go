package main

import "fmt"

func main() {
	var n, n2 int
	var m []int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&n2)
		m = append(m, n2)
	}
	fmt.Println(m[3])
}
