package main

import (
	"fmt"
)

func main() {
	var num, numrev string
	fmt.Scan(&num)
	for _, v := range num {
		numrev = string(v) + numrev
	}
	fmt.Println(numrev)
}
