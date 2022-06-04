package main

import "fmt"

const freezingF, boilingF = 32.0, 212.0

func main() {
	fmt.Printf("Conversion for boiling point of water: %g°F = %g°C \n", freezingF, fToC(freezingF))
	fmt.Printf("Conversion for freezing point of water: %g°F = %g°C \n", freezingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
