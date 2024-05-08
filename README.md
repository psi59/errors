# Errors

This package is inspired by the `github.com/pkg/errors` package and provides simple error handling primitives and stack trace functionality in Go. It allows you to add context to errors, wrap errors with additional information, and retrieve stack traces for better debugging and error reporting.

## Introduction

The traditional error handling idiom in Go is roughly akin to:

```go
if err != nil {
    return err
}
```

However, this approach can result in errors being reported without sufficient context or debugging information. The custom error package addresses this issue by providing functions and types that allow you to add context to errors and retrieve stack traces.

## Features

- Create new errors with `New` and `Errorf` functions.
- Wrap errors with additional context using `Wrap`, `Wrapf`, and `WrapWithCause` functions.
- Format errors with stack traces using `%+v` verb in `fmt` package.
- Compatible with `errors`, `github.com/pkg/errors` package.

## Installation

To use this package in your Go project, you can install it using `go get`:

```shell
go get github.com/psi59/errors
```

## Usage

Here are some examples of how to use the custom error package:

### Creating New Errors

```go
err := errors.New("something went wrong")
```

### Wrapping Errors

```go
err := someFunction()
if err != nil {
    return errors.Wrap(err, "failed to do something")
}
```

### Retrieving Stack Traces

```go
err := someFunction()
if err != nil {
    fmt.Printf("%+v\n", err)
}
```

### Formatting Errors with Stack Traces

```go
err := someFunction()
if err != nil {
    fmt.Printf("%+v\n", err)
}
```

Output:
```
failed to do something: original error
	at github.com/psi59/main.someFunction(something.go:57)
	at github.com/psi59/main.someFunction(something.go:56)
	at github.com/psi59/main.someFunction(something.go:55)
```

## Benchmark

```
goos: darwin
goarch: arm64
pkg: github.com/psi59/errors
BenchmarkErrors
BenchmarkErrors/pkg/errors-stack-10
BenchmarkErrors/pkg/errors-stack-10-10         	 3084790	       382.7 ns/op	     312 B/op	       6 allocs/op
BenchmarkErrors/errors-stack-10
BenchmarkErrors/errors-stack-10-10             	42213016	        27.86 ns/op	      16 B/op	       1 allocs/op
BenchmarkErrors/pkg/errors-stack-100
BenchmarkErrors/pkg/errors-stack-100-10        	 1219909	       979.6 ns/op	     312 B/op	       6 allocs/op
BenchmarkErrors/errors-stack-100
BenchmarkErrors/errors-stack-100-10            	 1922048	       622.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkErrors/pkg/errors-stack-1000
BenchmarkErrors/pkg/errors-stack-1000-10       	  187238	      6383 ns/op	     312 B/op	       6 allocs/op
BenchmarkErrors/errors-stack-1000
BenchmarkErrors/errors-stack-1000-10           	  200432	      6013 ns/op	      16 B/op	       1 allocs/op
BenchmarkStackFormatting
BenchmarkStackFormatting/%s-stack-10
BenchmarkStackFormatting/%s-stack-10-10        	19674388	        60.12 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stack-10
BenchmarkStackFormatting/%v-stack-10-10        	19284439	        61.70 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stack-10
BenchmarkStackFormatting/%+v-stack-10-10       	 4179945	       287.1 ns/op	     160 B/op	       4 allocs/op
BenchmarkStackFormatting/%s-stack-30
BenchmarkStackFormatting/%s-stack-30-10        	19776687	        60.28 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stack-30
BenchmarkStackFormatting/%v-stack-30-10        	19264303	        61.69 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stack-30
BenchmarkStackFormatting/%+v-stack-30-10       	 4077853	       289.3 ns/op	     160 B/op	       4 allocs/op
BenchmarkStackFormatting/%s-stack-60
BenchmarkStackFormatting/%s-stack-60-10        	19691418	        60.39 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stack-60
BenchmarkStackFormatting/%v-stack-60-10        	19310701	        62.05 ns/op	       8 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stack-60
BenchmarkStackFormatting/%+v-stack-60-10       	 4149933	       290.0 ns/op	     160 B/op	       4 allocs/op
BenchmarkStackFormatting/%s-stacktrace-10
BenchmarkStackFormatting/%s-stacktrace-10-10   	17911460	        66.99 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stacktrace-10
BenchmarkStackFormatting/%v-stacktrace-10-10   	17212515	        68.87 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-10
BenchmarkStackFormatting/%+v-stacktrace-10-10  	 3886822	       309.6 ns/op	     184 B/op	       5 allocs/op
BenchmarkStackFormatting/%s-stacktrace-30
BenchmarkStackFormatting/%s-stacktrace-30-10   	17860774	        66.80 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stacktrace-30
BenchmarkStackFormatting/%v-stacktrace-30-10   	17184109	        68.65 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-30
BenchmarkStackFormatting/%+v-stacktrace-30-10  	 3900709	       309.5 ns/op	     184 B/op	       5 allocs/op
BenchmarkStackFormatting/%s-stacktrace-60
BenchmarkStackFormatting/%s-stacktrace-60-10   	17731980	        66.57 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%v-stacktrace-60
BenchmarkStackFormatting/%v-stacktrace-60-10   	17391933	        69.05 ns/op	      24 B/op	       1 allocs/op
BenchmarkStackFormatting/%+v-stacktrace-60
BenchmarkStackFormatting/%+v-stacktrace-60-10  	 3918705	       317.7 ns/op	     184 B/op	       5 allocs/op
```

## License

This package is licensed under the [MIT License](LICENSE).
