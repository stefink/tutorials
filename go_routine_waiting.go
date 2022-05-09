package main

import (
	"fmt"
	"sync"
	"time"
)

func helloWorker(id int) {
	fmt.Printf("Hello from %d worker\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d completed his job\n", id)
}

func main() {
	var wg sync.WaitGroup
	for id := 0; id < 5; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			helloWorker(id)
		}(id)
	}
	wg.Wait()
}
