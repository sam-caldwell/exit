package exit

const (
	/*
	 * Notes on Exit Codes:
	 * 1. Exit codes in this file are for POSIX systems (e.g. Linux, MacOS, BSD).  They will have issues on Windows
	 *    since ExitCode 1 on Windows is success, and 0 is an error state.  Other non-zero values MAY be errors.
	 * 2. The ExitCode must be 0...255 (1 byte).
	 * 3. We do not use iota here to define exit code values because any change could cause values to change, which
	 *    would break any external commands executing a program that uses this tooling.
	 */

	// Success - (Linux/macOS/BSD) - No errors. Successful termination.
	Success Code = 0

	// GeneralError - (Linux/macOS/BSD) - General undefined error.
	GeneralError Code = 1

	// ConnectionFailed - (Linux/macOS/BSD) - A connection failed.
	ConnectionFailed Code = 10
	// ConnectionTimeout - (Linux/macOS/BSD) - A connection timed out.
	ConnectionTimeout Code = 11

	// FailedToFreeLock - (Linux/macOS/BSD) - A lock failed to free.
	FailedToFreeLock Code = 20

	// FileNotFound - (Linux/macOS/BSD) - A file was not found.
	FileNotFound Code = 21

	// FailedToOpenFile - (Linux/macOS/BSD) - A file failed to be opened.
	FailedToOpenFile Code = 22

	// LockCreateFailed - (Linux/macOS/BSD) - A required lock could not be created.
	LockCreateFailed Code = 30

	// InvalidCommand - (Linux/macOS/BSD) - An invalid command was encountered.
	InvalidCommand Code = 40

	// InvalidInput - (Linux/macOS/BSD) - Invalid input encountered.
	InvalidInput Code = 3

	// MissingArg - (Linux/macOS/BSD) - The program is missing an argument.
	MissingArg Code = 10

	// MissingColor - (Linux/macOS/BSD) - The program is missing a color input.
	MissingColor Code = 13

	// NotFound (Linux/macOS/BSD) - A required entity is not found.
	NotFound = 20

	// UnknownCommand - The given command is unknown.
	UnknownCommand = 11
)
