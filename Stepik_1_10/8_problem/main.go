package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b)
	for a != 0 {
		n *= 10
		n += a % 10
		a /= 10
	}
	for n != 0 {
		c := n % 10
		n /= 10
		d := b
		for d != 0 {
			if c == d%10 {
				fmt.Printf("%d ", c)
			}
			d /= 10
		}
	}
}
