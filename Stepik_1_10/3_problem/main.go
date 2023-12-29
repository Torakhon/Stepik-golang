package main

import "fmt"

func main() {

	var (
		sum int
		m   int
		n   int
	)

	fmt.Scan(&m)
	for i := 1; i <= m; i++ {
		fmt.Scan(&n)
		if 10 <= n && n%8 == 0 && n <= 99 {
			sum += n
		}
	}
	fmt.Println(sum)
}
