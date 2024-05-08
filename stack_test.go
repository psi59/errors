package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"strconv"
	"testing"
)

var testFrame = frame()

type X struct{}

// val returns a Frame pointing to itself.
func (x X) val() Frame {
	return frame()
}

// ptr returns a Frame pointing to itself.
func (x *X) ptr() Frame {
	return frame()
}

func TestFrame_String(t *testing.T) {
	s := frame()
	assert.Equal(t, "github.com/psi59/errors.TestFrame_String(stack_test.go:26)", s.String())
	assert.Equal(t, "github.com/psi59/errors.init(stack_test.go:11)", testFrame.String())
	var x X
	assert.Equal(t, "github.com/psi59/errors.X.val(stack_test.go:17)", x.val().String())
	assert.Equal(t, "github.com/psi59/errors.(*X).ptr(stack_test.go:22)", x.ptr().String())
}

func TestFrame_Format(t *testing.T) {
	var x X
	tests := []struct {
		frame  Frame
		format string
		want   string
	}{
		{
			frame:  x.ptr(),
			format: "%v",
			want:   "github.com/psi59/errors.(*X).ptr(stack_test.go:22)",
		},
		{
			frame:  x.ptr(),
			format: "%+v",
			want:   "\n\tat github.com/psi59/errors.(*X).ptr(stack_test.go:22)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.frame.String(), func(t *testing.T) {
			got := fmt.Sprintf(tt.format, tt.frame)
			assert.Equal(t, tt.want, got)
		})
	}
}

func frame() Frame {
	pc, _, _, _ := runtime.Caller(1)
	return Frame{pc: pc}
}

func TestStackTrace_Format(t *testing.T) {
	var x X
	tests := []struct {
		stack  StackTrace
		format string
		want   string
	}{
		{
			stack: StackTrace{
				x.ptr(),
				x.val(),
			},
			format: "%v",
			want:   "",
		},
		{
			stack: StackTrace{
				x.ptr(),
				x.val(),
			},
			format: "%+v",
			want: "\n\tat github.com/psi59/errors.(*X).ptr(stack_test.go:22)" +
				"\n\tat github.com/psi59/errors.X.val(stack_test.go:17)",
		},
		{
			stack: StackTrace{
				x.ptr(),
				testFrame,
			},
			format: "%+v",
			want: "\n\tat github.com/psi59/errors.(*X).ptr(stack_test.go:22)" +
				"\n\tat github.com/psi59/errors.init(stack_test.go:11)",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			got := fmt.Sprintf(tt.format, tt.stack)
			assert.Equal(t, tt.want, got)
		})
	}
}
