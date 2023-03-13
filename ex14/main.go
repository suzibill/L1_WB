package main

import "fmt"

func main() {
	checkType(1)
	checkType(true)
	checkType("aboba")
	checkType(1.2)
	checkType(make(chan int))
	checkType(make(chan float64))
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}
