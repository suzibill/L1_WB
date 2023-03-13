package main

import "fmt"

func main() {
	s := "привет 😀 hello"
	s = reverseString(s)
	fmt.Println(s)
}

func reverseString(s string) string {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}
	return string(r)
}
