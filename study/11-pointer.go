package main

import "fmt"

func main() {
	a := 1

	var b *int
	b = &a
	*b += 1
	fmt.Println(a, *b)
}
