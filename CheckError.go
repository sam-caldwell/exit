package exit

import (
	"github.com/sam-caldwell/ansi/v2"
)

// CheckError - Check if error is nil and if not print an error message and terminate with exit code (1).
//
//	If err != nil, we print the error and terminate.
//	But if err == nil, we will do nothing.
//
//go:inline
func CheckError(err error) {
	CheckErrorf(err, "Error: %v")
}

// CheckErrorf - Check if error is nil and if not print a formatted error string and terminate with exit code (1).
//
//	If err != nil, we print the error with a format string and terminate.
//	But if err == nil, we will do nothing.
func CheckErrorf(err error, fmtString string) {
	if err != nil {
		ansi.Red().Printf(fmtString, err.Error()).Reset()
		terminate(GeneralError)
	}
}

// CheckErrorWithCode - if err != nil, print error message and exit with the given code.
func CheckErrorWithCode(err error, exitCode Code) {
	if err != nil {
		ansi.Red().Printf("Error: %s", err.Error()).Reset()
		terminate(exitCode)
	}
}
