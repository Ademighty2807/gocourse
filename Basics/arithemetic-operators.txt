package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3

	var result int

	result = a + b
	fmt.Printf("Addition: %d\n", result)

	result = a - b
	fmt.Printf("Subtraction: %d\n", result)

	result = a * b
	fmt.Printf("Multiplication: %d\n", result)

	result = a / b
	fmt.Printf("Division : %d\n", result)

	result = a % b
	fmt.Printf("Modulus: %d\n", result)

	const pi = 22 / 7.0
	fmt.Printf("Value of pi: %f\n", pi)

	// Overflow with signed integers

	var maxInt int64 = 9223372036854775807 // Maximum value for int64

	fmt.Printf("Max int64: %d\n", maxInt)

	maxInt = maxInt + 1
	fmt.Printf("Max int64 + 1: %d\n", maxInt)

	// Overflow with unsigned integers

	var uMaxInt uint64 = 18446744073709551615 // Maximum value for uint64

	fmt.Printf("Max uint64: %d\n", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Printf("Max uint64 + 1: %d\n", uMaxInt)

	// Underflow with floating-point numbers
	var smallFloat float64 = 1.0e-323 // Smallest positive float64 value
	fmt.Printf("Smallest positive float64:", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64 // Divide by the maximum float64 value to cause underflow
	fmt.Printf("Smallest positive float64: ", smallFloat)
}
