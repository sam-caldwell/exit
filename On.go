package exit

import (
	"fmt"
	"github.com/sam-caldwell/ansi/v2"
)

// On - A standard way to quickly terminate a program if a given condition is true.
//
//	If usage is present, format it as part of the output.  Otherwise, usage is empty ("").
//	If err != "" print a formatted error message. Terminate the program with safe deferred
//	functions using this package. See exit.Defer() and README.md.
func On(condition bool, code Code, err error, usage string) {

	const (
		errMessage = "Error: %s%s"
		errUsage   = "\n\nUsage:\n%s\n"
	)

	if condition {
		var usageMsg string

		if code != 0 {
			ansi.Red()
		}

		if usage != "" {
			usageMsg = fmt.Sprintf(errUsage, usage)
		}

		if err == nil || err.Error() == "" {
			ansi.Printf("%v%s\n", err, usageMsg)
		} else {
			ansi.Printf(errMessage, err, usageMsg)
		}

		ansi.Reset()

		terminate(code)

	}
}
