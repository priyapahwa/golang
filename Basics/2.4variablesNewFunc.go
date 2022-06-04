// The expression new(T) creates an unnamed variable of type T
package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	q := new(int)
	fmt.Println(p == q)
}
