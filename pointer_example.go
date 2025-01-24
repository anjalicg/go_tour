package main

import "fmt"

func main_pointer() {
	i, j := 1, 12
	fmt.Println(i + j)
	p := &i
	q := &j
	fmt.Println(p, q)
	fmt.Printf("Types for p and q: %T, %T\n", p, q)

	fmt.Println(*p)
	fmt.Println(*p * 9.0)
}
