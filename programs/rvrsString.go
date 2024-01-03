package main

import "fmt"

type NewString string

func main() {
	fmt.Println(sreverse("hello"))
	// var d NewString = "amit"
	d := NewString("amit")
	fmt.Println(d.reverse())
}

func sreverse(s string) string {
	newS := ""
	for i := len(s) - 1; i >= 0; i-- {
		newS = newS + string(s[i])
	}
	return newS
}

//Receiver function
func (s NewString) reverse() string {
	newS := ""
	for i := len(s) - 1; i >= 0; i-- {
		newS = newS + string(s[i])
	}
	return newS
}
