package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsing Integer:", num)
	fmt.Println("Parsing Integer:", num+1)

	num64, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}

	fmt.Println("Parsed Integer:", num64)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Printf("Parsed float: %2f\n", floatval)

	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hex value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)

	fmt.Println("Parsed binary to decimal:", decimal)

	invalidNum := "456abc"
	InvalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing invalid value:", err)
		return
	}
	fmt.Println("Parsed invalid to decimal:", InvalidParse)
}
