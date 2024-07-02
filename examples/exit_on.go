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

	lhs := flag.String("lhs", "", "")
	rhs := flag.String("rhs", "", "")
	n := flag.Int("code", 0, "")
	usage := flag.String("usage", "", "")
	flag.Parse()
	code := exit.Code(*n)
	err := fmt.Errorf("lhs (%v) != rhs (%v) returns code (%v) usage: %s", *lhs, *rhs, code, *usage)
	exit.On(*lhs == *rhs, code, err, *usage)
	fmt.Println("ok")
}