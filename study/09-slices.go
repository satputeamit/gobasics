package main

import "fmt"

type Test struct {
	name string
	age  int
}

func main() {
	myslice := []Test{}
	fmt.Println(myslice)
	mt := Test{
		name: "abc",
		age:  12,
	}

	mt2 := Test{
		name: "lop",
		age:  55,
	}
	myslice = append(myslice, mt)
	myslice = append(myslice, mt2)
	fmt.Println(myslice[0].age)

	//

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr[:5])
}
