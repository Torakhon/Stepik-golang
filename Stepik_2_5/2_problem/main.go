package main

import "fmt"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {

	var str string
	fmt.Scan(&str)

	strRev := reverse(str)
	if str == strRev {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
