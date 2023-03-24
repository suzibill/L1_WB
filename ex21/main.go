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

type adapter struct {
	adaptee *Adaptee
}

func (a *adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	a := &Adaptee{name: "I am adaptee"}
	OneAdapter := &adapter{adaptee: a}
	fmt.Println(OneAdapter.Request())
}
