package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	// Pointers are comparable;
	// Two pointers are equal iff they point to same variables or both are nil
	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil)
}
