package main

func countDown(n int) {
	for i := 1; i <= n; i++ {
		println(i)
	}
}

func add(a int, b int) int {
	return a + b
}

func main() {
	countDown(100)
	println(add(4, 5))
}
