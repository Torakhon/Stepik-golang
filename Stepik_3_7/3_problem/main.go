package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	dat := strings.Split(s, ",")
	dat1, _ := time.Parse("02.01.2006 15:04:05", dat[0][:19])
	dat2, _ := time.Parse("02.01.2006 15:04:05", dat[1][:19])
	var res time.Duration
	if dat1.After(dat2) {
		res = dat1.Sub(dat2)
	} else {
		res = dat2.Sub(dat1)
	}
	fmt.Println(res)

}
