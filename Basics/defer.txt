package main

import "fmt"

func main() {
	// defer statement is used to delay the execution of a function until the surrounding function returns.

	// sorrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	process(10)

}

func process(i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("Deferred statement executed")
	fmt.Println("Normal statement executed")
	defer fmt.Println("Deferred statement executed 2")
	defer fmt.Println("Deferred statement executed 3")
	fmt.Println("Normal statement executed2")

	fmt.Println("Value of i:", i)
}

// Pratical Use cases
// 1. Resource cleanup: Defer is commonly used to ensure that resources are properly released, such as closing files or network connections, even if an error occurs.
// 2. Logging, Tracing and debugging: Defer can be used to log function entry and exit points, making it easier to trace the flow of execution and identify issues.
// 3. Unlocking mutexes: In concurrent programming, defer can be used to ensure that mutexes are unlocked, preventing deadlocks and ensuring proper synchronization.
// 4. Error handling: Defer can be used to handle errors gracefully by deferring error handling functions that will be executed when the surrounding function returns, allowing for proper cleanup and error reporting.

// Best practices for using defer in Go:
// 1. Keep defer actions short and simple.
// 2. Use defer at the beginning of a function to ensure it's executed before the function returns.
// 3. Avoid using defer in loops, as it can lead to performance issues.
// 4. Be careful with the order of defer statements, as they are executed in LIFO (Last In, First Out) order.
// 5 Avoid complex control logic in deferred functions, as it can make the code harder to understand and maintain.
