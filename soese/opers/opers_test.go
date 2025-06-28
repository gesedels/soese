package opers

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/langs/queue"
	"github.com/gesedels/soese/soese/langs/stack"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// setup
	s := stack.New()
	Opers["test"] = func(q *queue.Queue, s *stack.Stack) error {
		s.Push(1)
		return nil
	}

	// success
	err := Run(nil, s, "test")
	assert.Equal(t, []cell.Cell{1}, s.Cells)
	assert.NoError(t, err)

	// error - does not exist
	err = Run(nil, s, "")
	assert.EqualError(t, err, `operator "" does not exist`)
}
