package main

import "fmt"

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func main_vertex() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	fmt.Println("v.X:", v.X)
	pv := &v
	pv.X = 12 // dot operator with the pointer variable also
	fmt.Println("v:", v)
	v1 := Vertex{4, 5} // A struct literal denotes a newly allocated struct value
	// by listing the values of its fields.
	v2 := Vertex{X: 20}
	fmt.Println(v1)
	fmt.Println(v2) // Automatically puts 0 for Y as it is not assigned
	pv1 := &Vertex{100, 90}
	fmt.Println(pv1)
	fmt.Printf("Type pv1 =%T and for pv=%T\n", pv1, pv)
	fmt.Printf("Value pv1 =%v and for pv=%v\n", pv1, pv)
}
