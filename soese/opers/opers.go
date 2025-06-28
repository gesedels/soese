// Package opers implements operation types and definitions.
package opers

import (
	"fmt"

	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/gesedels/soese/soese/langs/queue"
	"github.com/gesedels/soese/soese/langs/stack"
)

// Oper is a built-in function that operates on a Queue and Stack.
type Oper func(*queue.Queue, *stack.Stack) error

// Opers is a map of all existing Oper functions.
var Opers = map[name.Name]Oper{
	// maths
	// "+": Add2,
}

// Run applies a named Oper to a Queue and Stack.
func Run(q *queue.Queue, s *stack.Stack, n name.Name) error {
	oper, ok := Opers[n]
	if !ok {
		return fmt.Errorf("operator %q does not exist", n)
	}

	return oper(q, s)
}
