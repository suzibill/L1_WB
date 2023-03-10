package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
}

func (a Human) hello() {
	fmt.Printf("Hello %s", a.Name)
}

func main() {
	fmt.Println("Hello")
	a := Action{Human{12, "Hui"}}
	a.hello()
}
