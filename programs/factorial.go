package main

import (
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	num, _ := strconv.Atoi(input)
	println(facto(num))

}

func facto(n int) int {
	if n <= 1 {
		return 1
	}
	return n * facto(n-1)
}
