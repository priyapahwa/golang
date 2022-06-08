package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 2
	m["key2"] = 4
	fmt.Println("Required Map:", m)

	val1 := m["key1"]
	fmt.Println("Value 1:", val1)

	fmt.Println("Length:", len(m))

	delete(m, "key2")
	fmt.Println("Modified Map:", m)

	_, prs1 := m["key1"]
	_, prs2 := m["key2"]
	fmt.Println("prs:", prs1, prs2)

	n := map[string]int{"foo": 4, "bar": 2}
	fmt.Println("New Map:", n)
}
