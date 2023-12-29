package main

import "fmt"

func main() {
	var max, n, x int
	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		switch {
		case x > max:
			max, n = x, 1
		case x == max:
			n++
		}
	}
	fmt.Println(n)
}
