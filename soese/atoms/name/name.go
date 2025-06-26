// Package name implements the Name Atom type and methods.
package name

import (
	"fmt"
	"strings"
)

// Name is a single value reference string.
type Name string

// New returns a new Name from a string.
func New(text string) Name {
	return Name(strings.ToLower(text))
}

// Parse returns a new Name from a parsed string.
func Parse(text string) (Name, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("invalid Name %q", text)
	}

	return New(text), nil
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
