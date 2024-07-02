package exit

import (
	"fmt"
)

// On - A standard way to quickly terminate a program if a given condition is true.
//
//	On code == 0, we simply terminate if the condition is true
//	On code != 0 (error state), we will--
//	  display a formatted error (if err != nil && err.Error() != "")
//	  display a formatted usage message (if usage != "")
func On(condition bool, code Code, err error, usage string) {

	const (
		errMessage = "Error: %s%s"
		errUsage   = "\n\nUsage:\n%s\n"
	)

	if condition {
		var usageMsg string

		if code != 0 {

			if usage != "" {
				usageMsg = fmt.Sprintf(errUsage, usage)
			}

			if err == nil || err.Error() == "" {
				fmt.Printf("%v%s\n", err, usageMsg)
			} else {
				fmt.Printf(errMessage, err, usageMsg)
			}

		}
		terminate(code)
	}
}
