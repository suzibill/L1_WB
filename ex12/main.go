package main

import "fmt"

func main() {
	a := []string{"aboba", "abeba", "aboba", "abeba", "abeba"}
	set := make(map[string]struct{})

	for _, v := range a {
		set[v] = struct{}{}
	}
	fmt.Println(set)
}
