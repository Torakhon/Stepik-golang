package main

import "fmt"

func main() {
	var m []int
	var n, n2 int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n2)
		m = append(m, n2)
	}
	for i := 0; i < n; i += 2 {
		fmt.Printf("%v ", m[i])
	}

}
