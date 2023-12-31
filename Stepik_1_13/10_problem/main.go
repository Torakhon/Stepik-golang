package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	for i := num2; i >= num1; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("NO")
}
