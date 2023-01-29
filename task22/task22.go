package main

import (
	"fmt"
	"math"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a, b, значение которых > 2^20.

// a+b
func sum(a, b float64) float64 {
	return a + b
}

// a-b
func diff(a, b float64) float64 {
	return a - b
}

// a*b
func mul(a, b float64) float64 {
	return a * b
}

// a/b
func div(a, b float64) float64 {
	return a / b
}
func main() {
	a := math.Pow(2, 25)
	b := math.Pow(2, 22)
	sum := sum(a, b)
	fmt.Printf("%f + %f = %f\n", a, b, sum)
	diff := diff(a, b)
	fmt.Printf("%f - %f = %f\n", a, b, diff)
	mul := mul(a, b)
	fmt.Printf("%f * %f = %f\n", a, b, mul)
	div := div(a, b)
	fmt.Printf("%f / %f = %f\n", a, b, div)

}
