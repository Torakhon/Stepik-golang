package main

import "fmt"

func main() {
	var num, n, n2 int
	fmt.Scan(&num)
	for {
		if num < 10 {
			n += num
			break
		} else {
			n += num % 10
			num = num / 10
		}
	}
	for {
		if n < 10 {
			n2 += n
			break
		} else {
			n2 += n % 10
			n = n / 10
		}
	}
	fmt.Println(n2)

}
