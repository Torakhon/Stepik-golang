package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var input string
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	persedTime1, err := time.Parse("2006-01-02 15:04:05", input[:19])
	if err != nil {
		fmt.Println(err)
		return
	}

	if input[11:13] >= "13" {
		persedTime1 = persedTime1.AddDate(0, 0, 1)
	}

	time_ := persedTime1.Format("2006-01-02 15:04:05")
	fmt.Println(time_)

}
