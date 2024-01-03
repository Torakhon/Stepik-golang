package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func adding(str1, str2 string) int64 {
	cleanStr1 := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str1)
	cleanStr2 := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str2)

	num1, _ := strconv.ParseInt(cleanStr1, 10, 64)
	num2, _ := strconv.ParseInt(cleanStr2, 10, 64)

	return num1 + num2
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}
