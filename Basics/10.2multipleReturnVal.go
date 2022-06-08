package main

import "fmt"

func vals() (int, int) {
	return 4, 8
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)
}
