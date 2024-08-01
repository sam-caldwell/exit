package exit_test

import (
	"testing"

	"github.com/sam-caldwell/exit"
)

func TestExit_Code(t *testing.T) {
	t.Run("Verify 8-bit value range", func(t *testing.T) {
		var c exit.Code
		for i := -512; i < 512; i++ {
			c = exit.Code(i)
			// nolint:staticcheck
			if c < 0 || c > 255 {
				t.Fatalf("c is out of bounds")
			}
		}
	})
	t.Run("Verify Code.ToInt()", func(t *testing.T) {
		for i := 0; i < 255; i++ {
			c := exit.Code(i)
			if c.ToInt() != i {
				t.Fatalf("mismatched value")
			}
		}
	})
}
