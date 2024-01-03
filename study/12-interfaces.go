package main

import "fmt"

type eBot struct {
	name string
}
type sBot struct {
	name string
}

type bot interface {
	greetings(int) (string, int) // decalrtion
	test()
}

func main() {

	eb := eBot{name: "amit"}
	sb := sBot{name: "abc"}

	printGreetings(eb)
	printGreetings(sb)

}

func printGreetings(b bot) {
	fmt.Println(b.greetings(1))
}

// Receiver function
func (e eBot) greetings(age int) (string, int) {
	return e.name + " hello", 0
}

func (eBot) test() {
	println("ok")
}

func (sBot) test() {
	println("ok")
}

func (s sBot) greetings(height int) (string, int) {
	return s.name + " hola", 2
}
