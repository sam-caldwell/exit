package main

import (
	"flag"
	"github.com/sam-caldwell/exit"
)

// main - example program for testing/demonstrating exit.Now()
//
// This program will allow an arbitrary exit_code to be passed into the program
// and the program will terminate with that same exit_code.
func main() {

	n := flag.Int("exit_code", 0, "")
	flag.Parse()
	code := exit.Code(*n)
	exit.Now(code)

}
