package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	cell := New(1.234)
	assert.Equal(t, Cell(1.234), cell)
}

func TestParse(t *testing.T) {
	// success
	cell, err := Parse("1.234")
	assert.Equal(t, Cell(1.234), cell)
	assert.NoError(t, err)

	// error - invalid Cell
	cell, err = Parse("")
	assert.Zero(t, cell)
	assert.EqualError(t, err, `invalid Cell ""`)
}

func TestBool(t *testing.T) {
	// success - true
	okay := Cell(1.234).Bool()
	assert.True(t, okay)

	// success - false
	okay = Cell(0).Bool()
	assert.False(t, okay)
}

func TestNative(t *testing.T) {
	// success
	flot := Cell(1.234).Native()
	assert.Equal(t, float64(1.234), flot)
}

func TestString(t *testing.T) {
	// success - no decimals
	text := Cell(1).String()
	assert.Equal(t, "1", text)

	// success - with decimals
	text = Cell(1.234).String()
	assert.Equal(t, "1.234", text)
}
