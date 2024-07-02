package exit

import "fmt"

// PanicOn - panic if condition is true, display message
func PanicOn(condition bool, msg any) {
	if condition {
		panic(msg)
	}
}

// PanicOnf - panic if condition is true, display formatted message
//
//go:inline
func PanicOnf(condition bool, format string, msg ...any) {
	PanicOn(condition, fmt.Sprintf(format, msg...))
}
