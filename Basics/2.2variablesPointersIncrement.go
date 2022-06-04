package main

import "fmt"

var p = f()

func main() {

	fmt.Println(f() == f())

	i := 1
	fmt.Println(i, &i)
	incr(&i)
	fmt.Println(i, &i)
	fmt.Println(incr(&i))
	fmt.Println(i, &i)

}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
