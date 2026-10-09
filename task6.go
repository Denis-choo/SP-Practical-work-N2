package main

import (
	"fmt"
	"time"
)

func main() {
	sources := []string{"База1", "База2", "База3"}
	responses := make(chan string, len(sources))

	for _, src := range sources {
		go func(s string) {
			if s == "База2" {
				time.Sleep(100 * time.Millisecond)
			} else {
				time.Sleep(300 * time.Millisecond)
			}

			responses <- "Результат из " + s
		}(src)
	}

	fastest := <-responses
	fmt.Println("Первый успешный результат:", fastest)
}
