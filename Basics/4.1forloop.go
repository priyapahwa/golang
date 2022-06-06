package main

import "fmt"

func main() {
	i := 6
	for i <= 12 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 8; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
