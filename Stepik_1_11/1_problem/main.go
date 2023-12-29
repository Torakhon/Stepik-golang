package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Printf("число %.2f не подходит\n", num)
	} else if num > 10000 {
		fmt.Printf("%e\n", num)
	} else {
		result := math.Pow(num, 2)
		roundedResult := math.Round(result*10000) / 10000
		fmt.Printf("%.4f\n", roundedResult)
	}
}
