package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

// run go program with -race flag
func main() {
	msg = "Hello, world!"
	// var mutex sync.Mutex

	// first race condition
	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	fmt.Println(msg)
}

// func main() {
// 	msg = "Hello, world!"
// 	var mutex sync.Mutex

// 	// first race condition
// 	wg.Add(2)
// 	go updateMessage("Hello, universe!", &mutex)
// 	go updateMessage("Hello, cosmos!", &mutex)
// 	wg.Wait()
// 	fmt.Println(msg)

// 	// second race condition
// 	var v int = 0
// 	for i := 0; i < 10_000; i++ {
// 		wg.Add(1)
// 		go f(&v, &mutex)
// 	}
// 	wg.Wait()
// 	fmt.Println("Finished", v)
// }

// func f(v *int, m *sync.Mutex) {
// 	m.Lock()
// 	*v++
// 	m.Unlock()
// 	wg.Done()
// }
