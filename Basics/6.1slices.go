package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Lenght", len(s))
	fmt.Println("Set:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("New Set:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy Set:", c)

	l1 := s[2:5]
	fmt.Println("Slice 1:", l1)

	l2 := s[:5]
	fmt.Println("Slice 2:", l2)

	l3 := s[2:]
	fmt.Println("Slice 3:", l3)

	t := []string{"g", "h", "i"}
	fmt.Println("Set 2:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Set:", twoD)
}
