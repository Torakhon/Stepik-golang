package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	x, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var str string
	for i := 0; i < utf8.RuneCountInString(x); i++ {
		n := strings.Count(x, string(x[i]))
		if n == 1 {
			str = strings.Join([]string{str, string(x[i])}, "")
		}
	}
	fmt.Println(str)
}
