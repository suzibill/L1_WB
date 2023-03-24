package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Aboba")
	Sleep1(1 * time.Second)
	fmt.Println("Aboba")
	Sleep2(2 * time.Second)
	fmt.Println("Aboba")
}

func Sleep1(d time.Duration) {
	<-time.After(d)
}

func Sleep2(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) > d {
			break
		}
	}
}
