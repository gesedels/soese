package queue

import (
	"fmt"
	"strings"

	"github.com/gesedels/soese/soese/atoms/atom"
)

// Queue is a first-in-first-out queue of Atoms.
type Queue struct {
	Atoms []atom.Atom
}

// New returns a new Queue from an Atom sequence.
func New(as ...atom.Atom) *Queue {
	return &Queue{as}
}

// Dequeue removes and returns the first Atom in the Queue.
func (q *Queue) Dequeue() (atom.Atom, error) {
	if len(q.Atoms) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	a := q.Atoms[0]
	q.Atoms = q.Atoms[1:]
	return a, nil
}

// DequeueSlice removes and returns the first n Atoms in the Queue.
func (q *Queue) DequeueSlice(n int) ([]atom.Atom, error) {
	if len(q.Atoms) < n {
		return nil, fmt.Errorf("queue is insufficient")
	}

	as := q.Atoms[:n]
	q.Atoms = q.Atoms[n:]
	return as, nil
}

// Empty returns true if the Queue is empty.
func (q *Queue) Empty() bool {
	return len(q.Atoms) == 0
}

// Enqueue appends an Atom to the end of the Queue.
func (q *Queue) Enqueue(a atom.Atom) {
	q.Atoms = append(q.Atoms, a)
}

// EnqueueSlice appends an Atom slice to the end of the Queue.
func (q *Queue) EnqueueSlice(as []atom.Atom) {
	q.Atoms = append(q.Atoms, as...)
}

// Len returns the number of Atoms in the Queue.
func (q *Queue) Len() int {
	return len(q.Atoms)
}

// String returns a string representation of the Queue.
func (q *Queue) String() string {
	var ss []string
	for _, a := range q.Atoms {
		ss = append(ss, a.String())
	}

	return strings.Join(ss, " ")
}
