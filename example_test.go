package errors_test

import (
	"fmt"
	"github.com/psi59/errors"
)

func ExampleNew() {
	err := errors.New("example error")
	fmt.Printf("%+v", err)
	// Output:
	// example error
	//	at github.com/psi59/errors_test.ExampleNew(example_test.go:9)
}

func ExampleErrorf() {
	err := errors.Errorf("example error %d", 123)
	fmt.Printf("%+v", err)
	// Output:
	// example error 123
	//	at github.com/psi59/errors_test.ExampleErrorf(example_test.go:17)
}

func ExampleWithStack() {
	err := errors.New("example error")
	err = errors.WithStack(err)
	fmt.Printf("%+v", err)
	// Output:
	//example error
	//	at github.com/psi59/errors_test.ExampleWithStack(example_test.go:22)
	//	at github.com/psi59/errors_test.ExampleWithStack(example_test.go:21)
}

func ExampleWrap() {
	err := errors.New("example error")
	err = errors.Wrap(err, "wrapped")
	fmt.Printf("%+v", err)
	// Output:
	// wrapped: example error
	//	at github.com/psi59/errors_test.ExampleWrap(example_test.go:36)
	//	at github.com/psi59/errors_test.ExampleWrap(example_test.go:35)
}

func ExampleWrapf() {
	err := errors.New("example error")
	err = errors.Wrapf(err, "wrapped %d", 123)
	fmt.Printf("%+v", err)
	// Output:
	// wrapped 123: example error
	//	at github.com/psi59/errors_test.ExampleWrapf(example_test.go:46)
	//	at github.com/psi59/errors_test.ExampleWrapf(example_test.go:45)
}

func ExampleWrapWithCause() {
	cause := errors.New("cause error")
	err := errors.New("example error")
	err = errors.WrapWithCause(err, cause)
	fmt.Printf("%+v", err)
	// Output:
	// example error: cause error
	//	at github.com/psi59/errors_test.ExampleWrapWithCause(example_test.go:57)
	//	at github.com/psi59/errors_test.ExampleWrapWithCause(example_test.go:56)
	//	at github.com/psi59/errors_test.ExampleWrapWithCause(example_test.go:55)
}
