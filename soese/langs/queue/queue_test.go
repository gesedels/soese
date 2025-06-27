package queue

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/atom"
	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/stretchr/testify/assert"
)

var foo, bar = name.Name("foo"), name.Name("bar")

func TestNew(t *testing.T) {
	// success
	q := New(foo, bar)
	assert.Equal(t, []atom.Atom{foo, bar}, q.Atoms)
}

func TestDequeue(t *testing.T) {
	// setup
	q := New(foo)

	// success
	a, err := q.Dequeue()
	assert.Equal(t, foo, a)
	assert.Empty(t, q.Atoms)
	assert.NoError(t, err)

	// error - Queue is empty
	a, err = q.Dequeue()
	assert.Nil(t, a)
	assert.EqualError(t, err, "Queue is empty")
}

func TestDequeueSlice(t *testing.T) {
	// setup
	q := New(foo, bar)

	// success
	as, err := q.DequeueSlice(2)
	assert.Equal(t, []atom.Atom{foo, bar}, as)
	assert.Empty(t, q.Atoms)
	assert.NoError(t, err)

	// error - Queue is empty
	as, err = q.DequeueSlice(1)
	assert.Empty(t, as)
	assert.EqualError(t, err, "Queue is insufficient")
}

func TestEmpty(t *testing.T) {
	// success - true
	b := New().Empty()
	assert.True(t, b)

	// success - false
	b = New(foo, bar).Empty()
	assert.False(t, b)
}

func TestEnqueue(t *testing.T) {
	// setup
	q := New()

	// success
	q.Enqueue(foo)
	assert.Equal(t, []atom.Atom{foo}, q.Atoms)
}

func TestEnqueueSlice(t *testing.T) {
	// setup
	q := New()

	// success
	q.EnqueueSlice([]atom.Atom{foo, bar})
	assert.Equal(t, []atom.Atom{foo, bar}, q.Atoms)
}

func TestLen(t *testing.T) {
	// success
	i := New(foo, bar).Len()
	assert.Equal(t, 2, i)
}

func TestString(t *testing.T) {
	// success
	s := New(foo, bar).String()
	assert.Equal(t, "foo bar", s)
}
