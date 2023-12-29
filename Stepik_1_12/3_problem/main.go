package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	nMax := array[0]
	for _, value := range array {
		if value > nMax {
			nMax = value
		}
	}
	fmt.Println(nMax)
}
