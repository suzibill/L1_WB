package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}
	wg.Add(5)
	for _, v := range a {
		go func(x int) {
			fmt.Println(x * x)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
