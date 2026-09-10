package cvp

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println(time.Now())
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func printLetters() {
// 	for _, letter := range "ABCDE" {
// 		fmt.Println(time.Now())
// 		fmt.Println(letter)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting\n", id)
	fmt.Println(time.Now())
	for range 100_000_000 {
	}
	fmt.Printf("Task %d is finished\n", id)
	fmt.Println(time.Now())
}

func main() {

	// go printNumbers()
	// go printNumbers()

	// time.Sleep(5 * time.Second)

	numThreads := 12

	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}
	wg.Wait()

}
