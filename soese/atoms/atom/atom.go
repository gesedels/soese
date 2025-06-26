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

// Atomise returns a parsed Atom from a string.
func Atomise(text string) (Atom, error) {
	if cell, err := cell.Parse(text); err == nil {
		return cell, nil
	}

	if name, err := name.Parse(text); err == nil {
		return name, nil
	}

	return nil, fmt.Errorf("invalid Atom %q", text)
}
