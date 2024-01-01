package main

import (
	"fmt"
	"unicode"
)

func IsValidPassword(password string) bool {
	rs := []rune(password)
	if len(rs) < 5 {
		return false
	}

	for _, r := range rs {
		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Digit, r) {
			return false
		}
	}

	return true
}

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	if IsValidPassword(text) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
