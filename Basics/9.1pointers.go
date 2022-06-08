package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroVal(i)
	fmt.Println("Zero Value:", i)

	zeroPtr(&i)
	fmt.Println("Zero Pointer:", i)

	fmt.Println("Pointer:", &i)

}
