package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}
	}()
	for _, v := range a {
		ch1 <- v
		fmt.Println(<-ch2)
	}

}
