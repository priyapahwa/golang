package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend vibes")
	default:
		fmt.Println("Weekday blues")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before Noon")
	default:
		fmt.Println("After Noon")
	}

	predictType := func(x interface{}) {
		switch y := x.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("Unknown %T\n", y)
		}
	}

	predictType(true)
	predictType(5)
	predictType("Clause")
}
