package main

import "fmt"

type Adapter interface {
	Request() string
}

type Adaptee struct {
	name string
}

func (a *Adaptee) SpecificRequest() string {
	return a.name
}

type AdapterImpl struct {
	adaptee *Adaptee
}

func (a *AdapterImpl) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	a := &Adaptee{name: "I am adaptee"}
	adapter := &AdapterImpl{adaptee: a}
	fmt.Println(adapter.Request())
}
