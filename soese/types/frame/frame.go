// Package frame implements the Frame type and methods.
package frame

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gesedels/soese/soese/types/queue"
	"github.com/gesedels/soese/soese/types/stack"
)

// Frame is a top-level program container and controller.
type Frame struct {
	Reader io.Reader
	Writer io.Writer
	Queue  *queue.Queue
	Stack  *stack.Stack
}

// New returns a new Frame from a Reader, Writer, Queue and Stack.
func New(r io.Reader, w io.Writer, q *queue.Queue, s *stack.Stack) *Frame {
	return &Frame{r, w, q, s}
}

// NewEmpty returns a new empty Frame from a Reader and Writer.
func NewEmpty(r io.Reader, w io.Writer) *Frame {
	return New(r, w, queue.New(), stack.New())
}

// Read returns a newline-ending string from the Frame's Reader.
func (f *Frame) Read() string {
	r := bufio.NewReader(f.Reader)
	s, _ := r.ReadString('\n')
	return s
}

// Write writes a formatted string to the Frame's Writer.
func (f *Frame) Write(s string) {
	fmt.Fprint(f.Writer, s)
}
