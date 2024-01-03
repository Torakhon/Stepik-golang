package main

import (
	"fmt"
)

func task(c chan int, n int) {
	c <- n + 1
}

func main() {
	fmt.Println("Started!")

	c := make(chan int)
	go task(c, 99)

	fmt.Println(<-c)

	fmt.Println("Finish!")
}
