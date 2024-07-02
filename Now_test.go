package exit

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestExit_Now(t *testing.T) {
	expectedAnsiResetCodes := []byte{27, 91, 48, 109}

	t.Run("test happy path (exit code 0)", func(t *testing.T) {
		actual, err := exec.Command("go", "run",
			"examples/exit_now.go", "-exit_code", "0").Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%s'\n"+
				"expected: '%v'\n", actual, expectedAnsiResetCodes)
		}
	})

	for i := 1; i <= 255; i++ {
		n := Code(i)
		t.Run(fmt.Sprintf("Test exit.Now(%d)", n), func(t *testing.T) {
			exitCode := fmt.Sprintf("%d", n)
			out, err := exec.Command("go", "run", "examples/exit_now.go", "-exit_code", exitCode).Output()

			if err != nil && err.Error() != "exit status 1" {
				t.Fatalf("test failed: %v", err)
			}
			if !bytes.Equal(out, expectedAnsiResetCodes) {
				t.Fatalf("output mismatch\n"+
					"out: (%s) %v", string(out), out)
			}
		})
	}
}
