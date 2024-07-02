Standardized Exit Codes and Mechanisms
======================================

## Objectives

Develop standardized exit codes and mechanisms to make program termination predictable across
projects.

---

## Usage

`go get github.com/sam-caldwell/exit`

---

## Features

### Exit Codes

A set of constants (e.g. `Success` (0), `GeneralError` (1)) which can be reused,
and which are dependable across projects.

These exit codes are defined in [`constants.go`](./constants.go) with a strong type
called `Code` based on `uint8` for improved type safety.

---

### Helper Functions

| Function or Group                             | Description                                                               | 
|-----------------------------------------------|---------------------------------------------------------------------------| 
| [`On()`](./On.go)                             | A function to exit on a given condition being `true`.                     |
| [`OnError()`](./OnError.go)                   | A function to exit on error condition.                                    |
| [`PanicOn()`](./PanicOn.go)                   | A set of functions to panic on a given condition                          |
| [`TerminateOnError()`](./TerminateOnError.go) | A function to terminate program execution if an error exists              |
| [`Now()`](./Now.go)                           | A function which terminates execution immediately and safely.             