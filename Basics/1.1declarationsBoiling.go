// Print boiling point of water
package main

import "fmt"

const bpCelsius = 100.0

func main() {
	var c = bpCelsius
	var f = (c * 9 / 5) + 32
	fmt.Printf("Boiling Point of Water = %g°C or %g°F \n", c, f)
}
