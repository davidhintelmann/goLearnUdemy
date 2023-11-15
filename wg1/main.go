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
// requires:
// go env -w CGO_ENABLED=1
// and install MinGW-w64 toolchain
// https://code.visualstudio.com/docs/cpp/config-mingw
func main() {
	msg = "Hello, world!"

	// race condition
	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	fmt.Println(msg)
}
