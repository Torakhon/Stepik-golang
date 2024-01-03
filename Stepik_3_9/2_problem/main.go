package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Started!")

	work := func() {
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")

	fmt.Println("Finish!")
}
