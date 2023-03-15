package main

import "fmt"

func main() {
	a := []string{"aboba", "abeba", "aboba", "abeba", "abeba"}
	set := make(map[string]struct{})
	var member struct{}

	for _, v := range a {
		set[v] = member
	}
	fmt.Println(set)
}
