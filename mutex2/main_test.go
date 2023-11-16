package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	fmt.Println("Before Pipe")
	r, w, _ := os.Pipe()
	fmt.Println("After Pipe")
	os.Stdout = w

	fmt.Println("Before call to main()")
	main()
	fmt.Println("After call to main()")

	_ = w.Close()

	fmt.Println("Before Read All")
	result, _ := io.ReadAll(r)
	fmt.Println("After Read All")
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Error("Incorrect Balance Returned!")
	}
}
