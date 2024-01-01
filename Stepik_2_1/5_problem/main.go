package main

import "fmt"

func sumInt(args ...int) (int, int) {
	sum := 0
	for _, x := range args {
		sum += x
	}
	return len(args), sum
}

func main() {
	fmt.Println(sumInt(4))
	fmt.Println(sumInt(1, 2, 3, 4))
}
