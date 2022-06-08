package main

import "fmt"

//arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(2, 4)
	sum(1, 3, 7)

	nums := []int{2, 3, 5, 6}
	sum(nums...)
}
