// Package atom implements the Atom interface and functions.
package atom

import (
	"fmt"

	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/atoms/name"
)

// Atom is a single typed program value.
type Atom interface {
	// Bool returns the Atom as a boolean.
	Bool() bool

	// Native returns the Atom as a native value.
	Native() any

	// String returns the Atom as a string.
	String() string
}

// Atomise returns a new Atom from a parsed string.
func Atomise(s string) (Atom, error) {
	if a, err := cell.Parse(s); err == nil {
		return a, nil
	}

	if a, err := name.Parse(s); err == nil {
		return a, nil
	}

	return nil, fmt.Errorf("invalid Atom %q", s)
}
