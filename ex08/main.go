package main

import "fmt"

func main() {
	var n int64 = 12
	fmt.Println(n)
	n = reverseBit(n, 0, 0)
	n = reverseBit(n, 1, 0)
	n = reverseBit(n, 2, 0)
	fmt.Println(n)
	n = reverseBit(n, 0, 1)
	n = reverseBit(n, 1, 0)
	n = reverseBit(n, 2, 1)
	fmt.Println(n)
	n = reverseBit(n, 1, 1)
	fmt.Println(n)
}

func reverseBit(n int64, bit uint, status int) int64 {
	if status == 1 {
		n |= 1 << (bit % 64)
	} else if status == 0 {
		n &^= 1 << (bit % 64)
	}
	return n
}
