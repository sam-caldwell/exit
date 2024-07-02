Standardized Exit Codes and Mechanisms
======================================

## Objectives

Develop standardized exit codes and mechanisms to make program termination predictable across
projects, including a safe termination feature which ensures that `defer` functions are
respected.

---
## Usage

`go get github.com/sam-caldwell/exit`

```go
/*
	func main() {
		defer exit.Done() // Ensure Done is called when main exits

		// Your main logic here

		// Program ends here
	}
*/
```


---
## Features

### Exit Codes

A set of constants (e.g. `Success` (0), `GeneralError` (1)) which can be reused,
and which are dependable across projects.

These exit codes are defined in [`constants.go`](./constants.go) with a strong type
called `Code` based on `uint8` for improved type safety.

---
### Safe Termination

Golang does not execute `defer` functions if `os.Exit()` is called. This could leave
sockets, file handles or other resources abandoned in some cases. The program quite
simply does not die cleanly. This package provides a reusable 'safe termination'
option.

---
#### SafeDefer (`exit.Defer()`)

Any `defer` which should be respected on program termination should be defined using
this `exit.Defer()` function.

To use `exit.Defer()`, call--

    ```go
        exit.Defer(func(){ _=f.Close() })
    ```

In this example, we create a deferred function which will consume the return of
`f.Close()` or which may include other more advanced features.

---
#### SafeDefer (`exit.DeferOnSucess()`)

Any `defer` which should be respected only on exit code `0` (`Success`) may call
`exit.DeferOnSuccess()`, which means `os.Exit()` will terminate any other scenario
without guaranteeing deferred functions execute.

---
#### CheckError()

A set of functions for evaluating Golang `error` states and terminating the program
consistently. This feature is designed to work with `exit.Defer()`.

These functions are found in [`CheckError.go`](./CheckError.go), and they include
a simple function that prints error and terminates if error is not nil, a function
that prints a formatted string

---
#### On()

This is a function, designed to work with `exit.Defer()` and terminate its program
if a given condition is true, printing an optional error message and usage string.
