package main

import "fmt"

func main() {
	nums := []int{2, 4, 7}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index:", i)
		}
	}

	keyVal := map[string]string{"one": "Arrays", "two": "Strings"}
	for k, v := range keyVal {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range keyVal {
		fmt.Println("Key:", k)
	}

	for i, c := range "Go Programming" {
		fmt.Println(i, c)
	}
}
