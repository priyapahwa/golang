package main

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))

	var cel Celcius
	var far Fahrenheit
	fmt.Println(cel == 0)
	fmt.Println(far >= 0)
	fmt.Println(cel == Celcius(far))

	tempC := FToC(212.0)
	fmt.Println(tempC.String())
	fmt.Printf("%v\n", tempC)
	fmt.Printf("%s\n", tempC)
	fmt.Println(tempC)
	fmt.Printf("%g\n", tempC)   //does not call string
	fmt.Println(float64(tempC)) //does not call string
}

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func (tempC Celcius) String() string {
	return fmt.Sprintf("%g°C", tempC)
}
