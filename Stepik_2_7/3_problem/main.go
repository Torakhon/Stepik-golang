package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	n := '0'

	for _, char := range x {
		if char > n {
			n = char
		}
	}

	fmt.Println(string(n))
}
