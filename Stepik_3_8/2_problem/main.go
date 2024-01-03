package main

import (
	"fmt"
)

func task2(c chan string, text string) {
	for i := 0; i < 5; i++ {
		c <- text + " "
	}
}

func main() {
	fmt.Println("Started!")

	c := make(chan string)
	go task2(c, "a")

	for i := 0; i < 5; i++ {
		fmt.Print(<-c)
	}

	fmt.Println("\nFinish!")
}
