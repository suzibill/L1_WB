package main

import "fmt"

func main() {
	a := "кринж"
	b := "база"

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
