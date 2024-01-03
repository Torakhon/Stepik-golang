package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var sum int

	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			os.Stderr.WriteString("Error")
			continue
		}
		sum += num
	}

	sumStr := strconv.Itoa(sum)
	os.Stdout.WriteString(sumStr)
}
