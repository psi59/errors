package errors

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

var workingDir string

func init() {
	s, _ := os.Getwd()
	workingDir = s
}

// StackTrace represents a stack trace, which is a collection of Frames.
// It provides methods for formatting and appending stack traces.
type StackTrace []Frame

// Format formats the StackTrace according to the fmt.State and verb.
// It iterates over each Frame in the StackTrace and calls its Format method.
func (s StackTrace) Format(state fmt.State, verb rune) {
	if verb == 'v' {
		if state.Flag('+') {
			for _, f := range s {
				f.Format(state, verb)
			}
		}
	}
}

func appendStackTrace(s, ss StackTrace) StackTrace {
	appended := make([]Frame, len(s)+len(ss))
	for i, frame := range s {
		appended[i] = frame
	}
	for i, frame := range ss {
		appended[i+len(s)] = frame
	}

	return appended
}

// Frame represents a single frame in a stack trace.
type Frame struct {
	pc uintptr
}

// String returns a formatted string representation of the Frame.
// It includes the function name, file path (relative to the working directory), and line number.
func (f Frame) String() string {
	file := "unknown"
	funcName := "unknown"
	var line int
	fn := runtime.FuncForPC(f.pc)
	if fn != nil {
		funcName = fn.Name()
		file, line = fn.FileLine(f.pc)
	}

	return fmt.Sprintf("%s(%s:%d)", funcName, strings.TrimPrefix(file, workingDir+"/"), line)
}

// Format formats the Frame according to the fmt.State and verb.
// If the verb is 'v' and the '+' flag is set in the fmt.State, it writes a newline,
// a tab character, and the formatted Frame. Otherwise, it writes the formatted Frame.
func (f Frame) Format(state fmt.State, verb rune) {
	if verb == 'v' {
		if state.Flag('+') {
			io.WriteString(state, "\n\tat ")
		}

		io.WriteString(state, f.String())
	}
}

// caller returns a new StackTrace starting from the specified number of frames to skip.
func caller(skip int) StackTrace {
	pc, _, _, _ := runtime.Caller(skip)
	return []Frame{{pc: pc}}
}
