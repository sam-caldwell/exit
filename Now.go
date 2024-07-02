package exit

// Now - Safely exit the program after running all deferred functions.
// Use this in place of os.Exit() where possible.
//
// Usage:
//
//	exit.Now(1) // Terminate with exit code 1.
//
//go:inline
func Now(code Code) {
	terminate(code)
}
