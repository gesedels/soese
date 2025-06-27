package atom

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// success - Cell
	a, err := Atomise("1.234")
	assert.Equal(t, cell.Cell(1.234), a)
	assert.NoError(t, err)

	// success - Name
	a, err = Atomise("name")
	assert.Equal(t, name.Name("name"), a)
	assert.NoError(t, err)

	// error - invalid Atom
	a, err = Atomise("")
	assert.Nil(t, a)
	assert.EqualError(t, err, `invalid Atom ""`)
}
