package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	_, el := divide(x, n)
	if el != nil {
		fmt.Println("ошибка ")
	} else {
		fmt.Println(x / n)
	}
}
