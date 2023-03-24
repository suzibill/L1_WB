package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = deleteElement(s, 2)
	fmt.Println(s)
}

func deleteElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
