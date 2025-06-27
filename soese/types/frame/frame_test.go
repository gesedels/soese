package frame

import (
	"bytes"
	"testing"

	"github.com/gesedels/soese/soese/atoms/atom"
	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/gesedels/soese/soese/types/queue"
	"github.com/gesedels/soese/soese/types/stack"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	f := New(nil, nil, queue.New(name.Name("foo")), stack.New(1))
	assert.Nil(t, f.Reader)
	assert.Nil(t, f.Writer)
	assert.Equal(t, []atom.Atom{name.Name("foo")}, f.Queue.Atoms)
	assert.Equal(t, []cell.Cell{1}, f.Stack.Cells)
}

func TestNewEmpty(t *testing.T) {
	// success
	f := NewEmpty(nil, nil)
	assert.Nil(t, f.Reader)
	assert.Nil(t, f.Writer)
	assert.Empty(t, f.Queue)
	assert.Empty(t, f.Stack)
}

func TestRead(t *testing.T) {
	// setup
	b := bytes.NewBufferString("foo\n")
	f := NewEmpty(b, nil)

	// success
	s := f.Read()
	assert.Equal(t, "foo\n", s)
}

func TestWrite(t *testing.T) {
	// setup
	b := bytes.NewBuffer(nil)
	f := NewEmpty(nil, b)

	// success
	f.Write("foo\n")
	assert.Equal(t, "foo\n", b.String())
}
