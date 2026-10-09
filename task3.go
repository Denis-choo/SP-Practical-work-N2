package main

import (
	"fmt"
	"time"
)

func main() {
	semm := make(chan struct{}, 5)
	tick := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 15; i++ {

		go func(id int) {
			semm <- struct{}{}
			defer func() { <-semm }()
			fmt.Println(id)
		}(i)
		<-tick
	}

}
