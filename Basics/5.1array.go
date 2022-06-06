package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty:", a)

	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("Get:", a[4])
	fmt.Println("Length:", len(a))
}
