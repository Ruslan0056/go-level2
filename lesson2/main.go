package main

import (
	"fmt"
	"runtime/debug"
)

type ErrorWithTrace struct {
	text  string
	trace string
}

func New(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
	}
}

// The method returns error information
func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s", e.text, e.trace)
}

// The function performs a division, while checking the value of the divisor.
// If the divisor is 0 then the function returns an error.
func div(m, n int) (int, error) {
	if n == 0 {
		return 0, New("Нельзя делить на ноль")
	}
	return m / n, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := div(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
