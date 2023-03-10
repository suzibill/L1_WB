package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for data := range ch {
			fmt.Print(data)
		}
	}()

	go func() {
		for {
			ch <- 1
		}
	}()

	time.Sleep(1 * time.Second)
}
