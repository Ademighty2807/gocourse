package main

import "fmt"

func main() {

	// var ptr *int

	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	// fmt.Println(ptr)
	// fmt.Println(*ptr) // dereferencing a pointer

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	// var a int = 10

	// // ptr = &a

	// fmt.Println(a)
	// fmt.Println(&a)

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}

// A pointer is a variable that stores the memory address of another variable
