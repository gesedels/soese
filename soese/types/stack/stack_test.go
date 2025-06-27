package stack

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	s := New(1, 2, 3)
	assert.Equal(t, []cell.Cell{1, 2, 3}, s.Cells)
}

func TestEmpty(t *testing.T) {
	// success - true
	b := New().Empty()
	assert.True(t, b)

	// success - false
	b = New(1, 2, 3).Empty()
	assert.False(t, b)
}

func TestLen(t *testing.T) {
	// success
	i := New(1, 2, 3).Len()
	assert.Equal(t, 3, i)
}

func TestPop(t *testing.T) {
	// setup
	s := New(1)

	// success
	c, err := s.Pop()
	assert.Equal(t, cell.Cell(1), c)
	assert.Empty(t, s.Cells)
	assert.NoError(t, err)

	// error - Stack is empty
	c, err = s.Pop()
	assert.Zero(t, c)
	assert.EqualError(t, err, "Stack is empty")
}

func TestPopSlice(t *testing.T) {
	// setup
	s := New(1, 2, 3)

	// success
	cells, err := s.PopSlice(3)
	assert.Equal(t, []cell.Cell{3, 2, 1}, cells)
	assert.Empty(t, s.Cells)
	assert.NoError(t, err)

	// error - Stack has fewer Cells
	cells, err = s.PopSlice(1)
	assert.Empty(t, cells)
	assert.EqualError(t, err, "Stack is insufficient")
}

func TestPush(t *testing.T) {
	// setup
	s := New()

	// success
	s.Push(1)
	assert.Equal(t, []cell.Cell{1}, s.Cells)
}

func TestPushSlice(t *testing.T) {
	// setup
	s := New()

	// success
	s.PushSlice([]cell.Cell{1, 2, 3})
	assert.Equal(t, []cell.Cell{1, 2, 3}, s.Cells)
}

func TestString(t *testing.T) {
	// success
	s := New(1, 2, 3).String()
	assert.Equal(t, "1 2 3", s)
}
