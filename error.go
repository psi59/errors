// Package errors provides simple error handling primitives and stack trace functionality.
//
// The traditional error handling idiom in Go is roughly akin to:
//
//	if err != nil {
//	        return err
//	}
//
// which when applied recursively up the call stack results in error reports
// without context or debugging information. The errors package allows
// programmers to add context to the failure path in their code in a way
// that does not destroy the original value of the error.
//
// This package exposes the following functions from the standard library's
// errors package as variables for convenience:
//
//   - Is: A function that reports whether any error in the error chain matches a target error.
//   - As: A function that finds the first error in the error chain that matches a given type.
//   - Unwrap: A function that returns the result of calling the Unwrap method on an error, if any.
//
// Additionally, this package provides the following functions:
//
//   - New: Returns a new error with the given error message.
//   - Errorf: Formats according to a format specifier and returns a new error.
//   - WithStack: Annotates the given error with a stack trace.
//   - Wrap: Wraps the given error with a new error that includes the provided message.
//   - Wrapf: Wraps the given error with a new error that includes a formatted message.
//   - WrapWithCause: Wraps the given error with a new error that includes the cause.
//
// The package also provides the `withStack` type, which represents an error with
// an associated stack trace. It implements the `error` interface and provides
// additional methods for error handling and formatting.
//
// The `StackTrace` type represents a stack trace, which is a linked list of `Frame`
// structs. It provides methods for formatting and manipulating the stack trace.
//
// The `Frame` type represents a single frame in a stack trace, containing the
// program counter (pc) of the frame. It provides methods for formatting the frame
// information.
//
// Example usage:
//
//	err := errors.New("some error")
//	err = errors.WithStack(err)
//	err = errors.Wrap(err, "additional context")
//	err = errors.Wrapf(err, "more context with %s", "formatting")
//	err = errors.WrapWithCause(err, cause)
//
//	if errors.Is(err, some_error) {
//	        // Handle the specific error
//	}
//
//	var target_error TargetError
//	if errors.As(err, &target_error) {
//	        // Handle the target error
//	}
//
//	underlying_error := errors.Unwrap(err)
//	if underlying_error != nil {
//	        // Handle the underlying error
//	}
//
//	fmt.Printf("%+v", err) // Print the error with stack trace
package errors

import (
	"errors"
	"fmt"
	"io"
)

var (
	// WithMessage is a variable that is an alias for the Wrap function.
	WithMessage = Wrap
	// WithMessagef is a variable that is an alias for the Wrapf function.
	WithMessagef = Wrapf
)

// New returns a new error with the given error message.
// The error is annotated with a stack trace starting from the point where this function is called.
func New(msg string) error {
	return newWithStack(errors.New(msg), 3)
}

// Errorf formats according to a format specifier and returns a new error.
// The error is annotated with a stack trace starting from the point where this function is called.
func Errorf(format string, args ...any) error {
	return newWithStack(fmt.Errorf(format, args...), 3)
}

// WithStack annotates the given error with a stack trace.
// If the given error is nil, nil is returned.
// The stack trace starts from the point where this function is called.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return newWithStack(err, 3)
}

// newWithStack is an internal helper function that creates a new error with a stack trace.
// If the given error already has a stack trace, it appends the new stack trace to the existing one.
//
// Parameters:
//   - err: The error to annotate with a stack trace.
//   - skip: The number of stack frames to skip when generating the stack trace.
//
// Returns:
//   - A new error annotated with a stack trace.
func newWithStack(err error, skip int) error {
	s := caller(skip)
	var w *withStack
	if errors.As(err, &w) {
		w.stack = appendStackTrace(s, w.stack)
		return w
	}

	return &withStack{
		err:   err,
		stack: s,
	}
}

// Wrap wraps the given error with a new error that includes the provided message.
// If the given error is nil, nil is returned.
// The returned error is annotated with a stack trace starting from the point where this function is called.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	withMsg := fmt.Errorf("%s: %w", msg, err)
	var w *withStack

	stack := caller(2)
	if errors.As(err, &w) {
		stack = appendStackTrace(stack, w.stack)
	}

	return &withStack{
		err:   withMsg,
		stack: stack,
	}
}

// Wrapf wraps the given error with a new error that includes a formatted message.
// If the given error is nil, nil is returned.
// The returned error is annotated with a stack trace starting from the point where this function is called.
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	withMsg := fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
	var w *withStack

	stack := caller(2)
	if errors.As(err, &w) {
		stack = appendStackTrace(stack, w.stack)
	}

	return &withStack{
		err:   withMsg,
		stack: stack,
	}
}

// WrapWithCause wraps the given error with a new error that includes the cause.
// If the given error is nil, nil is returned.
// If the given cause is nil, the original error is wrapped with a stack trace.
// The returned error is annotated with a stack trace starting from the point where this function is called.
func WrapWithCause(err, cause error) error {
	if err == nil {
		return nil
	}
	if cause == nil {
		return newWithStack(err, 3)
	}
	withMsg := fmt.Errorf("%w: %w", err, cause)
	var errWithStack, causeWithStack *withStack

	stack := caller(2)
	if errors.As(err, &errWithStack) {
		stack = appendStackTrace(stack, errWithStack.stack)
	}
	if errors.As(cause, &causeWithStack) {
		stack = appendStackTrace(stack, causeWithStack.stack)
	}

	return &withStack{
		err:   withMsg,
		stack: stack,
	}
}

// withStack is a struct that represents an error with an associated stack trace.
// It implements the error interface and provides additional methods for error handling.
type withStack struct {
	err   error
	stack StackTrace
}

// Is reports whether any error in the error chain matches the target error.
// It delegates the check to the errors.Is function from the standard library.
func (w *withStack) Is(target error) bool {
	return errors.Is(w.err, target)
}

// Unwrap returns the underlying error wrapped by the withStack struct.
// It allows access to the original error for further inspection or error chain traversal.
func (w *withStack) Unwrap() error {
	return w.err
}

// Error returns the error message associated with the withStack struct.
// It implements the error interface.
func (w *withStack) Error() string {
	return w.err.Error()
}

// Format formats the withStack struct according to the fmt.State and verb.
// It supports the following formatting verbs:
//   - 'v': If the '+' flag is set in the fmt.State, it writes the error message
//     followed by the formatted stack trace. If the '+' flag is not set,
//     it falls through to the 's' verb behavior.
//   - 's': Writes the error message.
//   - 'q': Writes the quoted error message.
func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, w.Error())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
