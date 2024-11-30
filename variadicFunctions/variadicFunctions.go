package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func getMultiple(nums ...int) int {
	multiple := 1

	for _, nums := range nums {
		multiple *= nums
	}
	return multiple
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(nums...)

	var multiply int = getMultiple(nums...)
	fmt.Println("getMultiple(nums) =", multiply)
}
