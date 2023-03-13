package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "привет 😀 hello"
	s = reverseWords(s)
	fmt.Println(s)
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	fmt.Println(ss)
	var res string
	for i := len(ss) - 1; i >= 0; i-- {
		res += ss[i]
		if i != 0 {
			res += " "
		}
	}
	return res
}
