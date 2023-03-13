package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func main() {
	ch := make(chan int)
	var c Counter
	wg := sync.WaitGroup{}
	wg.Add(1)
	go counter(5, ch, &c, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println(c)
}

func counter(worker int, ch chan int, c *Counter, wg *sync.WaitGroup) {
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func(i int) {
			for v := range ch {
				c.mu.Lock()
				c.count++
				c.mu.Unlock()
				fmt.Printf("count: %d worker: %d\n", v, i)
			}
			wg.Done()
		}(i)
	}
	wg.Done()
}
