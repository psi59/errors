package errors

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	got := New("EOF")
	assert.EqualError(t, got, "EOF")
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		format string
		args   []any
		want   string
	}{
		{
			format: "failed to execute query: %s, %v",
			args:   []any{"SELECT * FROM err WHERE id=?", "test"},
			want:   "failed to execute query: SELECT * FROM err WHERE id=?, test",
		},
		{
			format: "EOF: %v",
			args:   []any{"test"},
			want:   "EOF: test",
		},
	}
	for _, tt := range tests {
		got := Errorf(tt.format, tt.args...)
		t.Run(got.Error(), func(t *testing.T) {
			assert.Equal(t, got.Error(), tt.want)
		})
	}
}

func TestWithStack(t *testing.T) {
	t.Run("nil err", func(t *testing.T) {
		got := WithStack(nil)
		assert.Nil(t, got)
	})

	t.Run("OK", func(t *testing.T) {
		got := WithStack(io.EOF)
		assert.ErrorIs(t, got, io.EOF)
		assert.EqualError(t, got, io.EOF.Error())
	})
}

func TestWrap(t *testing.T) {
	t.Run("nil err", func(t *testing.T) {
		got := Wrap(nil, "failed to execute query")
		assert.Nil(t, got)
	})

	t.Run("OK", func(t *testing.T) {
		got := Wrap(io.EOF, "failed to execute query")
		assert.ErrorIs(t, got, io.EOF)
		assert.EqualError(t, got, fmt.Sprintf("failed to execute query: %v", io.EOF))
	})
}

func TestWrapf(t *testing.T) {
	t.Run("nil err", func(t *testing.T) {
		got := Wrapf(nil, "failed to execute query: %v", "test")
		assert.Nil(t, got)
	})

	t.Run("OK", func(t *testing.T) {
		got := Wrapf(io.EOF, "failed to execute query: %v", "test")
		assert.ErrorIs(t, got, io.EOF)
		assert.EqualError(t, got, fmt.Sprintf("failed to execute query: %v: %v", "test", io.EOF))
	})
}

func TestWrapWithCause(t *testing.T) {
	err := Wrap(io.EOF, "failed to execute query")
	cause := Errorf("cause")
	t.Run("nil err", func(t *testing.T) {
		got := WrapWithCause(nil, cause)
		assert.Nil(t, got)
	})

	t.Run("nil cause", func(t *testing.T) {
		got := WrapWithCause(err, nil)
		assert.ErrorIs(t, got, err)
		assert.EqualError(t, got, err.Error())
	})

	t.Run("OK", func(t *testing.T) {
		got := WrapWithCause(err, cause)
		assert.ErrorIs(t, got, err)
		assert.ErrorIs(t, got, cause)
		assert.EqualError(t, got, fmt.Sprintf("%v: %v", err, cause))
	})
}

func Test_withStack_Format(t *testing.T) {
	var x X
	tests := []struct {
		err    error
		stack  StackTrace
		format string
		want   string
	}{
		{
			err: fmt.Errorf("test"),
			stack: StackTrace{
				x.val(),
				x.ptr(),
			},
			format: "%v",
			want:   "test",
		},
		{
			err: fmt.Errorf("test"),
			stack: StackTrace{
				testFrame,
				x.val(),
			},
			format: "%s",
			want:   "test",
		},
		{
			err: fmt.Errorf("test"),
			stack: StackTrace{
				x.val(),
				x.ptr(),
			},
			format: "%+v",
			want: "test" +
				"\n\tat github.com/psi59/errors.X.val(stack_test.go:17)" +
				"\n\tat github.com/psi59/errors.(*X).ptr(stack_test.go:22)",
		},
		{
			err: fmt.Errorf("test"),
			stack: StackTrace{
				testFrame,
				x.val(),
			},
			format: "%+v",
			want: "test" +
				"\n\tat github.com/psi59/errors.init(stack_test.go:11)" +
				"\n\tat github.com/psi59/errors.X.val(stack_test.go:17)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.err.Error(), func(t *testing.T) {
			w := &withStack{
				err:   tt.err,
				stack: tt.stack,
			}
			got := fmt.Sprintf(tt.format, w)

			assert.Equal(t, tt.want, got)
		})
	}
}
