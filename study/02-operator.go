package main

func main() {
	// Arithmetic
	//  + - * / %
	n1 := 4
	n2 := 3

	println("Add:", n1+n2)
	println("diff:", n1-n2)
	println("mul:", n1*n2)
	println("divd:", n1/n2)
	println("remndr:", n1%n2) //?

	// Relational op
	// < >  = !

	println(n1 > n2)
	println(n1 < n2)
	println(n1 >= n2)
	println(n1 != n2)
	println(n1 == n2)

	//logical op
	// && ||

	name := "abc"
	age := 25

	println("invite::", name == "abc" && age > 23)
	println("invite::", name == "abcd" || age > 26)
	println("invite::", !(name == "abc") && age > 23)

}
