package main

import "testing"

func Test_updateMessage(t *testing.T) {
	// pass
	msg = "Hello, world!"
	// fail
	// msg := "Hello, world!"

	wg.Add(2)

	// error
	// go updateMessage("x") // race condition
	go updateMessage("Goodbye, cruel world!")
	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("Incorrect value in msg")
	}
}
