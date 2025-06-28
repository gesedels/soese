// Package logic implements program parsing and evaluation functions.
package logic

import (
	"strings"

	"github.com/gesedels/soese/soese/atoms/atom"
	"github.com/gesedels/soese/soese/langs/queue"
)

// Parse returns an Atom slice from a parsed string.
func Parse(s string) ([]atom.Atom, error) {
	var as []atom.Atom
	for s := range strings.FieldsSeq(s) {
		a, err := atom.Atomise(s)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// ParseQueue returns a populated Queue from a parsed string.
func ParseQueue(s string) (*queue.Queue, error) {
	as, err := Parse(s)
	if err != nil {
		return nil, err
	}

	return queue.New(as...), nil
}
