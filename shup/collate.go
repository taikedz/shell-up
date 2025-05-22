package shup

import (
	"strings"
	"fmt"
	"io"
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
	return strings.Join(re.stack, "\n")
}

func Collate(stacksize int, filepath string, output io.Writer, registry *FileRegistry) *RecursionError {
	if stacksize >= MAX_RECUR {
		err := RecursionError{}
		return &err
	}

	lines, err := ReadLines(filepath)
	if err != nil {
		Fail(ERROR_FILE, "Could not load %s : %v", filepath, err)
	}

	for line_no, line := range(lines) {
		if found, target := getIncludeTarget(line); found {
			is_new, abspath, err := registry.Register(target) // FIXME - resolve target relative to filepath !
			if err != nil {
				Fail(ERROR_SHUP, "Could not register '%s' (%s:%d)", target, filepath, line_no+1)
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
			line_data := []byte(line)
			line_data = append(line_data, '\n')
			if _, err := output.Write(line_data); err != nil {
				Fail(ERROR_WRITE, "Could not write : %s", err)
			}
		}
	}

	return nil
}

func getIncludeTarget(line string) (bool, string) {
	line = strings.Trim(line, BLANKS)
	if strings.Index(line, MACRO_INCLUDE) == 0 {
		return true, line[len(MACRO_INCLUDE):]
	}
	return false, ""
}