package main

import "fmt"

func main() {
	var a [2]string //size specified
	a[0] = "anjali"
	a[1] = "cg"
	fmt.Println(a)

	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)

	// slices :dynamically sized,flexible view into elements of an array

	var s []int //size empty
	s = nums[1:3]
	fmt.Println(s)

	//A slice literal is like an array literal without the length.
	// [3]bool{true, true, false} is array
	// []bool{true, true, false} is slice

}
