package main

import "fmt"

func main() {

	fmt.Println("one")
	panicFunCalled(1, 0)
	fmt.Println("end main")

}

func panicFunCalled(a int, b int) {
	defer tryRecover()
	fmt.Println("Two")
	// panic("Panic state called")
	d := a / b
	fmt.Println("Panicked", d)
}

func tryRecover() {
	err := recover()
	if err != nil {
		fmt.Println("Recovered:", err)
	}
}
