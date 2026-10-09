package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := make(chan int)
	reads := make(chan chan int)

	go func() {
		state := 0
		for {
			select {
			case val := <-writes:
				state += val
			case reply := <-reads:
				reply <- state
			}
		}
	}()
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(val int) {
			defer wg.Done()
			writes <- val
		}(i)
	}

	wg.Wait()

	reply := make(chan int)
	reads <- reply
	final := <-reply

	fmt.Printf("Итоговое состояние: %d\n", final)
}
