package main

import "fmt"

func main() {
	var i int

	fmt.Println("Enter an integer to find whether it is even or odd: ")
	fmt.Scanln(&i)

	if i%2 == 0 {
		fmt.Println(i, "is an even integer")
	} else {
		fmt.Println(i, "is an odd integer")
	}
}
