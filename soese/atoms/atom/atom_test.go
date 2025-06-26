package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// error - invalid Atom
	atom, err := Atomise("")
	assert.Nil(t, atom)
	assert.EqualError(t, err, `invalid Atom ""`)
}
