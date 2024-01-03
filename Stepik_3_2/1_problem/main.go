package main

import (
	"fmt"
)

func convert(x int64) uint16 {
	var a uint16
	a = uint16(x)
	return a
}

func main() {
	var num = int64(1000)
	var num2 = convert(num)
	fmt.Printf("%T %v\n", num2, num2)
}
