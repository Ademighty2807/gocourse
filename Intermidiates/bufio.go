package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// // Bufio dot Reader
	// reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	// // Reading byte slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// line, err := reader.ReadString('\n')

	// if err != nil {
	// 	fmt.Println("Error reading string:", err)
	// 	return
	// }
	// fmt.Println("Read string:", line)

	// Bufio dot Writter

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice

	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writting string

	str := "This is a string.\n"

	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
