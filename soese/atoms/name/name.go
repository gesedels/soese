// Package name implements the Name Atom type and methods.
package name

import (
	"fmt"
	"strings"
)

// Name is a single reference string Atom.
type Name string

// New returns a new Name from a string.
func New(s string) Name {
	return Name(strings.ToLower(s))
}

// Parse returns a new Name from a parsed string.
func Parse(s string) (Name, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("invalid name %q", s)
	}

	return New(s), nil
}

// Bool returns the Name as a boolean.
func (n Name) Bool() bool {
	return string(n) != ""
}

// Native returns the Name as a native value.
func (n Name) Native() any {
	return string(n)
}

// String returns the Name as a string.
func (n Name) String() string {
	return string(n)
}
