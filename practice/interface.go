package main

import "fmt"

type Common interface {
	getName() string
}

type Person struct {
	name string
	age  int
}

type Vehicale struct {
	model  string
	number string
}

func main() {

	p := Person{name: "Ag", age: 20}
	v := Vehicale{model: "XCG54", number: "MH10"}
	PrintName(p)
	PrintName(v)
}

func PrintName(g Common) {
	fmt.Println(g.getName())
}

func (p Person) getName() string {
	return p.name
}

func (v Vehicale) getName() string {
	return v.model
}
