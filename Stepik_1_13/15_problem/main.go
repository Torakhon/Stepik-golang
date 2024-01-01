package main

import (
	"fmt"
	"strings"
)

func main() {
	var num, num2 string
	fmt.Scan(&num, &num2)
	str := strings.Replace(num, num2, "", -1)
	fmt.Println(str)
}
