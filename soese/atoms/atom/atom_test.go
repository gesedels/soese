package atom

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// success - Cell
	atom, err := Atomise("1.234")
	assert.Equal(t, cell.Cell(1.234), atom)
	assert.NoError(t, err)

	// success - Name
	atom, err = Atomise("name")
	assert.Equal(t, name.Name("name"), atom)
	assert.NoError(t, err)

	// error - invalid Atom
	atom, err = Atomise("")
	assert.Nil(t, atom)
	assert.EqualError(t, err, `invalid Atom ""`)
}
