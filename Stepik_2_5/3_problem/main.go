package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, x2 string
	fmt.Scan(&x, &x2)
	n := strings.Index(x, x2)
	fmt.Println(n)
}
