package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// START OMIT

func handle() {
	panic("oh no!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	go handle()

	// For blocking, press ctrl+c to stop
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

// STOP OMIT

// START HELPER OMIT
func protect(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("I am your protector: %v\n", r)
		}
	}()

	f()
}

// STOP HELPER OMIT
