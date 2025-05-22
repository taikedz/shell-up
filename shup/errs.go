package shup

import (
	"os"
	"fmt"
)

// Integer values serve as actual process exit status codes.
const ERROR_SHUP int = 10 // generic SHUP error . Try to not use if a specific error can be more suitable
const ERROR_FILE int = 11
const ERROR_WRITE int = 12

type Failure struct {
	message string
}

func (f Failure) Error() string {
	return f.message
}

func (f Failure) Exit(code int) {
	Fail(code, f.message)
}

func Fail(errno int, message string, items... any) {
	message = fmt.Sprintf(message, items...)
	fmt.Printf("%s\n", message)
	os.Exit(errno)
}
