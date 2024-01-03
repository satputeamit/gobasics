package main

import "fmt"

func main() {
	nums := [3]int{1, 2, 3}
	nums[1] = 5
	fmt.Println(nums)
	//
	var sales [4]float32
	sales[0] = 10.2
	fmt.Println(sales)

	sales2 := []int{1, 2, 3}
	sales2 = append(sales2, 5)
	sales2 = append(sales2, 6)
	fmt.Println(sales2)

	fmt.Println(len(sales2))

}
