package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = float64(a) // Type conversion from int to float64

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)

	c := float64(b)
	d := bool(c > 5) // Type conversion from float64 to bool
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)

	e := 3.14
	f := int(e) // Type conversion from float64 to int
	fmt.Println("Value of e:", e)
	fmt.Println("Value of f:", f)

	// Type(value)

	g := "Hello World"
	var h []byte
	h = []byte(g) // Type conversion from string to []byte
	fmt.Println("Value of g:", g)
	fmt.Println("Value of h:", h)
	i := []byte{255, 120, 72} // Type conversion from string to []byte
	j := string(i)            // Type conversion from []byte to string
	fmt.Println("Value of j:", j)
	fmt.Println("Value of i:", i)
}
