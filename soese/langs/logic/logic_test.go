package logic

import (
	"testing"

	"github.com/gesedels/soese/soese/atoms/atom"
	"github.com/gesedels/soese/soese/atoms/cell"
	"github.com/gesedels/soese/soese/atoms/name"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// success
	as, err := Parse("1 foo")
	assert.Equal(t, []atom.Atom{cell.Cell(1), name.Name("foo")}, as)
	assert.NoError(t, err)
}

func TestParseQueue(t *testing.T) {
	// success
	q, err := ParseQueue("1 foo")
	assert.Equal(t, []atom.Atom{cell.Cell(1), name.Name("foo")}, q.Atoms)
	assert.NoError(t, err)
}
