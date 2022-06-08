package main

import "fmt"

func subtraction(a int, b int) int {
	return a - b
}

func addition(a, b, c int) int {
	return a + b + c
}

func main() {
	diff := subtraction(2, 7)
	fmt.Println("2-7:", diff)

	sum := addition(2, 4, 6)
	fmt.Println("2+4+6:", sum)
}
