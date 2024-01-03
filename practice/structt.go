package main

import "fmt"

type School struct {
	name    string
	address string
}

type Person struct {
	name   string
	age    int
	school School
}

func main() {

	s := School{
		name:    "fffff",
		address: "ppppppp",
	}
	p := Person{
		name:   "abc",
		age:    25,
		school: s,
	}

	fmt.Println(p.school.address)

}
