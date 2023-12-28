package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	a = a % 100
	a = a / 10
	fmt.Println(a)
}
