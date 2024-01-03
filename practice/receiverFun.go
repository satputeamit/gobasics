package main

import "fmt"

type NewStr string

func main() {

	st := NewStr("abc")
	fmt.Println(st.getLength())

	var st1 NewStr = "gfdddd"
	fmt.Println(st1.getLength())

}

func (s NewStr) getLength() int {
	return len(s)
}
