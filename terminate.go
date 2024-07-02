package exit

import (
	"os"
)

// terminate - Ensure program execution terminates and all deferred functions run only once.
//
//	This is an abstraction point for future work to overcome some parts of golang that
//	are not optimal.
//
//go:inline
func terminate(code Code) {
	os.Exit(code.ToInt())
}
