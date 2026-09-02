package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	// numbers[4] = 20

	// fmt.Println(numbers)

	// numbers[0] = 9

	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

	// fmt.Println("Fuits array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	// fmt.Println("Original Array:", originalArray)

	// fmt.Println("Copied Array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index,", ":", numbers[i])
	// }

	// for i, v := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// }

	// underscoe is used to ignore the value of the second return value

	a, _ := someFunction()
	fmt.Println("Value of a:", a)

	_, b := someFunction()
	fmt.Println("Value of b:", b)

	// ignoring the first return value and using the second return value
	_, c := someFunction()
	fmt.Println("Value of c:", c)

	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Are the arrays equal?", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:", matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray [3]int

	copiedArray = originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)

	// fmt.Println("Copied Array:", copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}
