package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 100
	storage := make(map[int]int)

	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(storage)
}
