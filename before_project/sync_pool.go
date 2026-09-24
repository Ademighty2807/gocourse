package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person,")
			return &person{}
		},
	}

	// Get an object
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18

	fmt.Println("Got Person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to pool")
}
