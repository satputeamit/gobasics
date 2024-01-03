package main

import "fmt"

func main() {
	// similar to interface
	type Animal struct {
		class  string
		age    int
		gender bool
	}

	teddy := Animal{
		class:  "dog",
		age:    5,
		gender: true,
	}

	fmt.Println(teddy.class)
	teddy.age = 7
	fmt.Println(teddy)

	cat := Animal{"cat", 1, false}
	fmt.Println(cat)

	// another way

	test := struct {
		class  string
		age    int
		gender bool
	}{
		class:  "default",
		age:    0,
		gender: true,
	}

	fmt.Println(test)

}
