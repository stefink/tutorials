package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	message := make(chan string)
	go func() {
		wg.Add(5)
		for i := 0; i < 5; i++ {
			go func() {
				defer wg.Done()
				message <- "hello"
			}()
		}
		wg.Wait()
		close(message)
	}()
	for msg := range message {
		fmt.Println(msg)
	}
}
