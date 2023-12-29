package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	num_a := num / 100
	num_b := num / 10 % 10
	num_c := num % 10

	if num_a != num_b && num_b != num_c && num_c != num_a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
