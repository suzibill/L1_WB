package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	var c big.Int
	_, err := fmt.Scanf("%v%v", a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Mul(a, b))
	fmt.Println(c.Div(a, b))
	fmt.Println(c.Add(a, b))
	fmt.Println(c.Sub(a, b))
}
