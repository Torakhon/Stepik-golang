package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)

	str := ""
	for i, char := range x {
		if i != 0 {
			str += "*"
		}
		str += string(char)
	}

	fmt.Println(str)
}
