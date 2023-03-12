package main

import "fmt"

type void struct{}

func main() {
	a := []string{"aboba", "abeba", "aboba", "abeba", "abeba"}
	set := make(map[string]void)
	var member void

	for _, v := range a {
		set[v] = member
	}
	fmt.Println(set)
}
