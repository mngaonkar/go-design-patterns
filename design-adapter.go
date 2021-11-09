package main

import "fmt"

type AmericanPowerInterface interface {
	OutputVoltage() int
}

type EuropeanPowerInterface interface {
	OutputPower() int
}

type EuropeanPower struct {
}

func (p EuropeanPower) OutputPower() int {
	return 230
}

type AmericanPower struct {
}

func (p AmericanPower) OutputVoltage() int {
	return 120
}

type EuropeanToAmericanAdapter struct {
	european EuropeanPower
}

// conversion logic
func (p EuropeanToAmericanAdapter) convert(volatge int) int {
	fmt.Println("voltage converted from european to american")
	return volatge - 110
}

// call convert logic keeping the interface same
func (p EuropeanToAmericanAdapter) OutputVoltage() int {
	return p.convert(p.european.OutputPower())
}

func main() {
	// adapter has interface similar to source
	var adapter AmericanPowerInterface

	// when using american power, no voltage coversion required
	adapter = AmericanPower{}
	fmt.Printf("voltage is %d\n", adapter.OutputVoltage())

	// when using european power with adapter
	adapter = EuropeanToAmericanAdapter{}
	fmt.Printf("voltage is %d\n", adapter.OutputVoltage())
}
