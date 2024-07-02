package exit

// Code - Exit code values (0-255) returned by a program
type Code uint8

// ToInt - Convert numeric ExitCode value to integer
func (e *Code) ToInt() int {
	return int(*e)
}
