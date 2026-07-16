package main

import "fmt"

func main() {
	num := 423
	fmt.Printf("%05d\n", num) // 00423

	message := "Hello"

	// message := "Hello dfgfhgsdfger" // |Hello dfgfhgsdfger| |Hello dfgfhgsdfger|
	fmt.Printf("|%10s|\n", message)  // |     Hello|
	fmt.Printf("|%-10s|\n", message) // |Hello     |

	// Go supports string interpolation using backticks

	message1 := "Hello \n World!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)

	fmt.Println(message2)

	// `\d` "\\d"

	sqlQuery := `SELECT * FROM users WHERE age > 30`
	fmt.Println(sqlQuery) // SELECT * FROM users WHERE age > 30
	// backtick help use to keep string the way it's without being affected by go characters
}
