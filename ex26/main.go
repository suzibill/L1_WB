package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueSymbols("aboba"))
	fmt.Println(uniqueSymbols("qwertyuiopasdfghjklzxcvbnm"))
	fmt.Println(uniqueSymbols("aaaaaas"))
}

func uniqueSymbols(s string) bool {
	s = strings.ToLower(s)
	set := make(map[rune]struct{})
	for _, v := range s {
		if _, exists := set[v]; exists {
			return false
		}
		set[v] = struct{}{}
	}
	return true
}
