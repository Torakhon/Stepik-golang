package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	fmt.Scan(&x)

	numberStr := strconv.Itoa(x)
	m := ""

	for _, digit := range numberStr {
		digitInt, _ := strconv.Atoi(string(digit))
		squared := digitInt * digitInt
		m += strconv.Itoa(squared)
	}

	fmt.Println(m)
}
