package main

import "fmt"

func main() {
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := [10]int{7, 9, 30, 4, 50, 6, 70, 8, 90, 10}
	fmt.Println(a1, a2)
    var resp []int
	index := make(map[int]struct{})
	for _, v := range a1 {
		m[v] = struct{}
	}
	for _, v := range a2 {
		if _, exists := index[v]; exists {
			resp = append(resp, v)
		}
	}
	fmt.Println(res)
}
