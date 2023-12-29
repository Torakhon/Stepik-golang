package main

import "fmt"

func main() {

	var (
		num   int
		sum_1 int
		sum_2 int
	)
	fmt.Scan(&num)

	for i := 0; i < 6; i++ {
		if i < 3 {
			sum_1 += num % 10
			num /= 10
		} else {
			sum_2 += num % 10
			num /= 10
		}
	}

	if sum_1 == sum_2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
