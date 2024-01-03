package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started!")

	calculator := func(arguments <-chan int, done <-chan struct{}) <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)

			var sum int

			for {
				select {
				case value := <-arguments:
					sum += value
				case <-done:
					result <- sum
					return
				}
			}
		}()

		return result
	}

	arguments := make(chan int)
	done := make(chan struct{})

	go func() {

		for i := 1; i < 10; i++ {
			arguments <- i
		}

		done <- struct{}{}
	}()
	fmt.Println("Result:", <-calculator(arguments, done))

	go func() {
		for i := 1; i < 5; i++ {
			arguments <- i
		}
		done <- struct{}{}
	}()
	fmt.Println("Result:", <-calculator(arguments, done))

	fmt.Println("Finish!")
}
