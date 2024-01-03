package main

import (
	"fmt"
	"time"
)

func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	idx := make(chan int)
	val := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- f(x1)
			}()
			go func() {
				done <- f(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-idx] = <-val
		}

		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}

func main() {
	fmt.Println("Started!")

	start := time.Now()

	fn := func(x int) int {
		time.Sleep(time.Second * 2)
		return x
	}

	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	n := 5

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("Send", i)
			go func() {
				in1 <- 1
				in2 <- 2
			}()
		}
	}()

	merge2Channels(fn, in1, in2, out, n)

	for i := 0; i < n; i++ {
		value := <-out
		fmt.Println("Result:", value)
	}

	fmt.Println("Elapsed:", time.Since(start))
	fmt.Println("Finish!")
}
