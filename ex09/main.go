package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	chanIn := make(chan int)
	chanOut := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	go func() {
		for v := range chanIn {
			chanOut <- v * 2
		}
	}()
	for _, v := range a {
		chanIn <- v
		_, err := fmt.Fprint(os.Stdout, <-chanOut)
		if err != nil {
			log.Fatal(err)
		}
	}

}
