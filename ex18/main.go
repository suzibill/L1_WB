package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int32
}

func main() {
	ch := make(chan int)
	var c Counter
	wg := sync.WaitGroup{}

	counter(5, ch, &c, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println(c)
}

func counter(worker int, ch chan int, c *Counter, wg *sync.WaitGroup) {
	wg.Add(worker + 1)
	for i := 0; i < worker; i++ {
		go func(i int) {
			for v := range ch {
				atomic.AddInt32(&c.count, 1)
				fmt.Printf("count: %d worker: %d\n", v, i)
			}
			wg.Done()
		}(i)
	}
	wg.Done()
}
