package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	a := 1
	var b int = 2
	c, _ := strconv.Atoi(os.Args[1])
	fmt.Println(a + b + c)

	lst := []int{1}
	lst = append(lst, 2)
	lst = append(lst, 3)
	fmt.Println(lst, lst[1])

	var lst1 []string = []string{}
	fmt.Println(lst1)

	var lst2 []string
	lst2 = append(lst2, "A")
	lst2 = append(lst2, "B")
	fmt.Println(lst2)

}
