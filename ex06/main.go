package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1 способ семафор
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Finish 1")
				return
			default:
				fmt.Print("A\n")
			}
		}
	}()

	time.Sleep(10000)
	quit <- true
	fmt.Println("Amogus")

	// 2 способ закрытие канала
	ch := make(chan int)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Finish 2")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- 1
	ch <- 2
	close(ch)
	time.Sleep(10000)

	// 3 способ контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Done")
		case <-ctx.Done():
			fmt.Println("Finish 3")
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(100000)
}
