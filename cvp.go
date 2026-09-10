package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(time.Now())
		fmt.Println(string(letter))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	go printNumbers()
	go printNumbers()

	time.Sleep(2 * time.Second)

}
