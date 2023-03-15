package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{7, 9, 30, 4, 50, 6, 70, 8, 90, 10}
	fmt.Println(a1, a2)
	m := make(map[int]int)
	for _, v := range a1 {
		m[v] += 1
	}
	for _, v := range a2 {
		m[v] += 1
	}
	var res []int
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}
	fmt.Println(res)
}
