package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var val int
	fmt.Print("Enter the number to calculate its factorial: ")
	fmt.Scanln(&val)

	fmt.Println("Factorial of", val, ":", fact(val))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Printf("The term %v of fibonacci sequence: %v \n", val, fib(val))
}
