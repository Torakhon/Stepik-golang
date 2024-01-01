package main

import "fmt"

func func_text(text string) {
	fmt.Println(text)
}

func main() {
	var str string
	fmt.Scan(&str)
	func_text(str)
}
