package main

import "fmt"

func minimumFromFour() int {
	var (
		m             []int
		n, a, b, c, d int
	)
	fmt.Scan(&a, &b, &c, &d)
	n = a
	m = append(m, a, b, c, d)
	for _, value := range m {
		if value < n {
			n = value
		}
	}
	return n
}

func main() {
	fmt.Println(minimumFromFour())
}
