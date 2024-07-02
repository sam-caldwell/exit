package exit

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestExit_On(t *testing.T) {

	t.Run("test happy path (exit code 0) but lhs != rhs", func(t *testing.T) {
		cmd := "go"
		args := []string{
			"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "b",
			"-code", "0", "-usage", "should_not_appear",
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
			"-code", "0", "-usage", "should_not_appear",
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

	//for n := 1; n <= 255; n++ {
	//	t.Run(fmt.Sprintf("Test exit.Now(%d)", n), func(t *testing.T) {
	//		cmd := "go"
	//		args := []string{
	//			"run", "examples/exit_on.go", "-lhs", "a", "-rhs", "b",
	//			"-code", fmt.Sprintf("%d", n), "-usage", "should_not_appear",
	//		}
	//
	//		exitCode := fmt.Sprintf("%d", n)
	//		out, err := exec.Command("go", "run", "examples/exit_now.go", "-exit_code", exitCode).Output()
	//
	//		if err != nil && err.Error() != "exit status 1" {
	//			t.Fatalf("test failed: %v", err)
	//		}
	//		if !bytes.Equal(out, expectedAnsiResetCodes) {
	//			t.Fatalf("output mismatch\n"+
	//				"out: (%s) %v", string(out), out)
	//		}
	//	})
	//}
}
