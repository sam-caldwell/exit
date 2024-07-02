package exit

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestExit_OnError(t *testing.T) {

	t.Run("test happy path (exit code 0)", func(t *testing.T) {
		cmd := "go"
		args := []string{
			"run", "examples/exit_onError.go", "-show_error", "-code", "0",
		}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expectedAnsiResetCodes := []byte{27, 91, 48, 109, 27, 91, 48, 109}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expectedAnsiResetCodes)
		}
	})
	t.Run("test happy path (exit code 1)", func(t *testing.T) {
		code := 1
		cmd := "go"
		args := []string{
			"run", "examples/exit_onError.go", "-show_error", "-code", fmt.Sprintf("%d", code),
		}
		actual, err := exec.Command(cmd, args...).Output()
		if code == 0 && err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expectedAnsiResetCodes := []byte{
			27, 91, 48, 109, 27, 91, 51, 49, 109, 69, 114, 114, 111, 114, 58, 32, 114, 101, 116, 117, 114, 110,
			115, 32, 99, 111, 100, 101, 32, 40, 49, 41, 32, 117, 115, 97, 103, 101, 58, 32, 27, 91, 48, 109}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: \n"+
				"          s: '%s'\n"+
				"          b: '%v'\n"+
				"expected: '%v'\n",
				string(actual), actual, expectedAnsiResetCodes)
		}
	})
	t.Run("test happy path (exit code 1)(with usage)", func(t *testing.T) {
		code := 1
		cmd := "go"
		args := []string{
			"run", "examples/exit_onError.go", "-show_error", "-code", fmt.Sprintf("%d", code),
			"-usage", "test_usage",
		}
		actual, err := exec.Command(cmd, args...).Output()
		if code == 0 && err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expectedAnsiResetCodes := []byte{
			27, 91, 48, 109, 27, 91, 51, 49, 109, 69, 114, 114, 111, 114, 58, 32, 114, 101, 116, 117,
			114, 110, 115, 32, 99, 111, 100, 101, 32, 40, 49, 41, 32, 117, 115, 97, 103, 101, 58, 32,
			116, 101, 115, 116, 95, 117, 115, 97, 103, 101, 10, 10, 85, 115, 97, 103, 101, 58, 10,
			116, 101, 115, 116, 95, 117, 115, 97, 103, 101, 10, 27, 91, 48, 109,
		}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: \n"+
				"          s: '%s'\n"+
				"          b: '%v'\n"+
				"expected: '%v'\n",
				string(actual), actual, expectedAnsiResetCodes)
		}
	})
}