package main

import "fmt"

func main() {
	fmt.Println("hello")
	var favBook = "my world"
	fmt.Println(favBook)

	favBook = "new world"
	fmt.Println(favBook)

	var secondBoook string = "2nd Book"
	fmt.Println((secondBoook))

	var count int
	count = 2
	fmt.Println(count)

	// Compound creation -  multiple variables
	var name, age, admin = "abc", 25, false
	fmt.Println(name, age, admin)

	// Block creation
	var (
		sName  = "xyz"
		sAge   = 23
		sAdmin = true
	)
	fmt.Println(sName, sAge, sAdmin)

	favAnimal := "Dog"
	fmt.Println(favAnimal)

	pet1, pet2, pet3 := "cat", "dog", "monkey"
	fmt.Println(pet1, pet2, pet3)

	// Reassign
	pet3 = "tiger"
	//or
	pet2, pet4 := "lion", "rat"

	fmt.Println(pet1, pet2, pet3, pet4)

	// constants

	const seed = 1

}
