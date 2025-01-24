package main

import (
	"fmt"
	"time" // the package name is the same as the last element of the import path
)

func add(x, y int) int { // type is after name and the return type is also specified
	// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return x + y
}

//Return multiple output example
func swap(x, y string) (string, string) {
	return y, x
}

// Named return function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main1() {
	defer fmt.Println("This is a deferred statement")
	fmt.Println("Current time is :", time.Now())
	fmt.Println("Add 3 and 2 gives:", add(3, 2))
	a, b := swap("hello", "world")
	fmt.Println("Swap `hello world`:", a, b)
	fmt.Println(split(17))

	// Type conversion
	var f float64 = 3.94
	fmt.Println("In float", f)
	var x uint64 = uint64(f)
	fmt.Println("In uint64", x)

	// Loops - Only one looping construct - the "for"
	// Has 3 components: init statement, condition statement, post statement
	// Variables declared in init statement are visible only in the scope of the "for" statement.
	// init and post statements are optional  --> same as while loop
	// if condition is also removed, it loops forever --> same as while True
	for i := 1; i <= 10; i++ {
		defer fmt.Println("Entering the loop for i =", i)
	}
}
