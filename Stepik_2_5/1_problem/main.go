package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(n)
	if unicode.IsUpper(rs[0]) && rs[len(rs)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
