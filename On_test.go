package exit

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/ansi/v2"
	"os/exec"
	"testing"
)

func TestExit_On(t *testing.T) {

	t.Run("test happy path (exit code 0) but lhs != rhs", func(t *testing.T) {
		cmd := "go"
		args := []string{
			"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "b",
			"-code", "0", "-usage", "should_not_appear", "-show_error",
		}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expectedAnsiResetCodes := []byte{27, 91, 48, 109, 111, 107, 10}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expectedAnsiResetCodes)
		}
	})

	t.Run("test happy path (exit code 0) but lhs == rhs", func(t *testing.T) {
		cmd := "go"
		args := []string{
			"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "a",
			"-code", "0", "-usage", "should_not_appear", "-show_error",
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

	t.Run("test with non-zero exit code (with usage)", func(t *testing.T) {
		for n := 1; n <= 255; n++ {
			t.Run(fmt.Sprintf("Test exit.On(%d)", n), func(t *testing.T) {
				exitCode := fmt.Sprintf("%d", n)
				cmd := "go"
				args := []string{
					"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "a",
					"-code", exitCode, "-usage", "should_not_appear", "-show_error",
				}
				out, err := exec.Command(cmd, args...).Output()

				if err != nil && err.Error() != "exit status 1" {
					t.Fatalf("test failed: %v", err)
				}

				expectedAnsiResetCodes := append(append([]byte{
					27, 91, 48, 109, 27, 91, 51, 49, 109, /*  0-8: header (not visible) */
					69, 114, 114, 111, 114, 58, 32, /*        9-15: Error: */
					108, 104, 115, 32, 40, 97, 41, /*         16-22: lhs (a) */
					32, 33, 61, 32, /*                        23-26: != */
					114, 104, 115, 32, 40, 97, 41, /*         27-33: rhs(a) */
					32, 114, 101, 116, 117, 114, 110, 115, /* 34-41: returns*/
					32, 99, 111, 100, 101, 32, /*             42-47: code */
					40},
					[]byte(exitCode)...),
					[]byte{41, 32, /*               48-52: (<n>) */
						117, 115, 97, 103, 101, 58, 32, /* usage: */
						115, 104, 111, 117, 108, 100, 95, 110, 111, 116, 95, 97, 112, 112, 101, 97, 114, 10, 10,
						85, 115, 97, 103, 101, 58, 10, 115, 104, 111, 117, 108, 100, 95, 110, 111, 116, 95, 97,
						112, 112, 101, 97, 114, 10, 27, 91, 48, 109}...)
				if len(out) != len(expectedAnsiResetCodes) {
					t.Fatalf("output length mismatch\n"+
						"out:      s: %s\n"+
						"          b: %03d\n"+
						"expected: s: %s\n"+
						"          b: %03d",
						string(out), out, string(expectedAnsiResetCodes), expectedAnsiResetCodes)
					return
				}
				if !bytes.Equal(out, expectedAnsiResetCodes) {
					for i := 0; i < len(out); i++ {
						ansi.Printf("%03d ", i)
					}
					fmt.Println()
					for i := 0; i < len(out); i++ {
						ansi.Printf("%03d ", out[i])
					}
					fmt.Println()
					for i := 0; i < len(expectedAnsiResetCodes); i++ {
						if out[i] != expectedAnsiResetCodes[i] {
							ansi.Red()
						}
						ansi.Printf("%03d ", expectedAnsiResetCodes[i])
						ansi.Reset()
					}
					fmt.Println()
					t.Fatal("output mismatch")
				}
			})
			t.Run(fmt.Sprintf("Test exit.On(%d) (no -usage)", n), func(t *testing.T) {
				exitCode := fmt.Sprintf("%d", n)
				cmd := "go"
				args := []string{
					"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "a",
					"-code", exitCode, "-show_error",
				}
				out, err := exec.Command(cmd, args...).Output()

				if err != nil && err.Error() != "exit status 1" {
					t.Fatalf("test failed: %v", err)
				}

				expectedAnsiResetCodes := append(append([]byte{
					27, 91, 48, 109, 27, 91, 51, 49, 109, /*  0-8: header (not visible) */
					69, 114, 114, 111, 114, 58, 32, /*        9-15: Error: */
					108, 104, 115, 32, 40, 97, 41, /*         16-22: lhs (a) */
					32, 33, 61, 32, /*                        23-26: != */
					114, 104, 115, 32, 40, 97, 41, /*         27-33: rhs(a) */
					32, 114, 101, 116, 117, 114, 110, 115, /* 34-41: returns*/
					32, 99, 111, 100, 101, 32, /*             42-47: code */
					40},
					[]byte(exitCode)...),
					[]byte{41, 32, /*               48-52: (<n>) */
						117, 115, 97, 103, 101, 58, 32, /* usage: */
						27, 91, 48, 109}...)
				if len(out) != len(expectedAnsiResetCodes) {
					t.Fatalf("output length mismatch\n"+
						"out:      s: %s\n"+
						"          b: %03d\n"+
						"expected: s: %s\n"+
						"          b: %03d",
						string(out), out, string(expectedAnsiResetCodes), expectedAnsiResetCodes)
					return
				}
				if !bytes.Equal(out, expectedAnsiResetCodes) {
					for i := 0; i < len(out); i++ {
						ansi.Printf("%03d ", i)
					}
					fmt.Println()
					for i := 0; i < len(out); i++ {
						ansi.Printf("%03d ", out[i])
					}
					fmt.Println()
					for i := 0; i < len(expectedAnsiResetCodes); i++ {
						if out[i] != expectedAnsiResetCodes[i] {
							ansi.Red()
						}
						ansi.Printf("%03d ", expectedAnsiResetCodes[i])
						ansi.Reset()
					}
					fmt.Println()
					t.Fatal("output mismatch")
				}
			})
		}
	})
}
