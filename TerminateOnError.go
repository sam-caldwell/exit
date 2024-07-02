package exit

import (
	"github.com/sam-caldwell/ansi/v2"
)

// TerminateOnError - If err is not nil, terminate with GeneralError.
//
//	Note that any error message will be printed in red using
//	ANSI Color Codes.
//
//go:inline
func TerminateOnError(err error) {
	TerminateOnErrorWithCode(err, GeneralError)
}

// TerminateOnErrorWithCode - If err is not nil, terminate with the given exit code.
//
//	Note that any error message will be printed in red using
//	ANSI Color Codes.
func TerminateOnErrorWithCode(err error, code Code) {
	if err != nil {
		ansi.Red().Println(err.Error()).Reset()
	}
	terminate(code)
}
