package main

import "fmt"

func main() {
	mFibbo(10)
}

func fibbo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibbo(n-1) + fibbo(n-2)

}

// rglr

func nFibbo(n int) {
	a, b := 0, 1
	for i := 0; i <= n; i++ {
		if i <= 1 {
			fmt.Println(i)
		}
		a, b = b, a+b
		fmt.Println(b)
	}
}

func mFibbo(n int) {
	l := []int{0, 1}
	if n <= 1 {
		println(l[n])
	}

	for len(l) < n+1 {
		l = append(l, l[(len(l)-1)]+l[(len(l)-2)])
	}

	fmt.Println(l)
}
