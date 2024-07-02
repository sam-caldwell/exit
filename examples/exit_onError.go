package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/exit/v2"
)

// main - example program for testing/demonstrating exit.On()
//
// This program will allow the user to test the On() functionality of the exit package.
func main() {
	showError := flag.Bool("show_error", false, "")
	n := flag.Int("code", 0, "")
	usage := flag.String("usage", "", "")
	flag.Parse()
	code := exit.Code(*n)
	var err error
	if *showError {
		err = fmt.Errorf("returns code (%v) usage: %s", code, *usage)
	}
	exit.OnError(err, code, *usage)
	fmt.Println("ok")
}
