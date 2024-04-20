package main

import "main/internal/cmd"

// main is the entry point of the Go program.
//
// It calls the Run method of the cmd package's Execute interface to run the Game function.
// There are no parameters or return types for this function.
func main() {
	cmd.Run().Game()
}
