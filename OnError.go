package exit

// OnError - A standard way to safely terminate on error
//
//go:inline
func OnError(err error, code Code, usage string) {

	On(err != nil, code, err, usage)

}
