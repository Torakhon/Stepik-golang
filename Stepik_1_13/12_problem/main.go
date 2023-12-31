package main

import "fmt"

func main() {
	var num int
	num2 := 1
	fmt.Scan(&num)
	for {
		if num2 <= num {
			fmt.Printf("%d ", num2)
			num2 *= 2
		} else {
			break
		}
	}
}
