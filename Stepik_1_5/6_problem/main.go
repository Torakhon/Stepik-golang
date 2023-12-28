package main

import "fmt"

func main() {

	var d int
	fmt.Scan(&d)
	h := d / 30
	m := 2 * (d % 30)
	fmt.Println("It is", h, "hours", m, "minutes.")
}
