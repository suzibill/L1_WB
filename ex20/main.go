package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "привет 😀 café"
	s = reverseWords(s)
	fmt.Println(s)
}

func reverseWords(s string) string {
	ss := strings.Fields(s)
	fmt.Println(ss)
	var sb strings.Builder
	for i := len(ss) - 1; i >= 0; i-- {
		sb.WriteString(ss[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
