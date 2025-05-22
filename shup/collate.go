package shup

import (
	"strings"
	"fmt"

	"github.com/taikedz/shell-up/shup"
)

const MAX_RECUR int = 100 // no sane shell script should be importing to this depth...?

const MACRO_INCLUDE string = "#%include "
const BLANKS string = " \t"


type RecursionError struct {
	stack []string
}

func (re *RecursionError) Add(file string, line int) {
	re.stack = append(re.stack, fmt.Sprintf("%s:%d", file, line))
}

func (re RecursionError) Error() string {
	return strings.Join("\n", re.stack)
}

func Collate(stacksize int, filepath string, output *Writer, registry *shup.FileRegistry) *RecursionError {
	if stacksize >= MAX_RECUR {
		err := RecursionError{}
		return &err
	}

	lines, err := loadFile(filepath)
	if err != nil {
		errs.Fail(errs.ERROR_FILE, "Could not load %s : %v", filepath, err)
	}

	for line_no, line := range(lines) {
		if found, target := getIncludeTarget(line); found {
			is_new, abspath, err := registry.Register(target)
			if err != nil {
				errs.Fail(errs.ERROR_SHUP, "Could not register '%s' (%s:%d)", target, filepath, line_no+1)
			}
			if is_new {
				recur_err := Collate(stacksize+1, abspath, output, registry)
				if recur_err != nil {
					recur_err.Add(filepath, line_no)
					return recur_err
				}
			}
		} else {
			// If nothing resolves, just write the line
			// Add any other macro handlers above here.
			if err := output.Write(line); err != nil {
				errs.Fail(errs.ERROR_WRITE, "Could not write : %s", err)
			}
		}
	}

	return nil
}

func loadFile(filepath string) []string, error {
	fh, err := os.Open(filepath)
	if err != nil {
		erss.Fail(errs.ERROR_FILE, "Could not open: '%s'", filepath)
	}
	defer fh.Close()

	return fh.ReadLines() // TODO
}

func getIncludeTarget(line string) (bool, string) {
	line = strings.Trim(line, BLANKS)
	if strings.IndexOf(MACRO_INCLUDE, line) == 0 {
		return true, line[len(MACRO_INCLUDE):]
	}
	return false, ""
}