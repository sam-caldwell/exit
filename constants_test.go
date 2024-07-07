package exit

import (
	"testing"
)

func TestExitCodes(t *testing.T) {
	data := []Code{
		Success,
		GeneralError,
		ConnectionFailed,
		ConnectionTimeout,
		FailedToFreeLock,
		FileNotFound,
		FailedToOpenFile,
		LockCreateFailed,
		InvalidCommand,
		InvalidInput,
		InvalidResult
		MissingArg,
		MissingColor,
		NotFound,
		ParseError,
		UnknownCommand,
	}
	t.Run("test values are in range", func(t *testing.T) {
		for i, d := range data {
			if d < 0 || d > 255 {
				t.Fatalf("failed on row %d with value %v", i, d)
			}
		}
	})
	t.Run("test the Code.ToInt() method", func(t *testing.T) {
		for _, c := range data {
			d := c.ToInt()
			if int(c) != d {
				t.Fatalf("failed on ToInt() c:%d, d:%d", c, d)
			}
		}
	})
}
