package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}
	wg.Add(5)
	r := 0
	for _, v := range a {
		go func(x int, r *int) {
			*r += x * x
			wg.Done()
		}(v, &r)
	}
	wg.Wait()
	fmt.Println(r)
}
