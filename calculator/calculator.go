package calculator

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
// Panics if b is 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// ErrDivideByZero is returned when a division by zero is attempted
var ErrDivideByZero = NewError("division by zero")

// Error represents a calculator error
type Error struct {
	message string
}

// NewError creates a new Error
func NewError(message string) error {
	return &Error{message: message}
}

// Error returns the error message
func (e *Error) Error() string {
	return e.message
}
