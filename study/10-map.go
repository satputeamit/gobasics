package main

import "fmt"

func main() {
	//dict - key value pair, unorder
	cart := make(map[string]int)

	cart["mobile"] = 3
	cart["laptop"] = 2

	data, found := cart["laptop"]

	//delete

	delete(cart, "laptop")

	fmt.Println(cart, data, found)
}
