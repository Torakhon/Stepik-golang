package main

import "fmt"

func vote(x int, y int, z int) int {
	var (
		m    []int
		a, b int
	)
	m = append(m, x, y, z)
	for _, value := range m {
		if value == 0 {
			a++
		} else if value == 1 {
			b++
		}
	}
	if a > b {
		x = 0
	} else if a < b {
		x = 1
	}
	return x
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(vote(a, b, c))
}
