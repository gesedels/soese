package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	c := New(1.234)
	assert.Equal(t, Cell(1.234), c)
}

func TestParse(t *testing.T) {
	// success
	c, err := Parse("1.234")
	assert.Equal(t, Cell(1.234), c)
	assert.NoError(t, err)

	// error - invalid Cell
	c, err = Parse("")
	assert.Zero(t, c)
	assert.EqualError(t, err, `invalid cell ""`)
}

func TestBool(t *testing.T) {
	// success - true
	b := Cell(1.234).Bool()
	assert.True(t, b)

	// success - false
	b = Cell(0).Bool()
	assert.False(t, b)
}

func TestNative(t *testing.T) {
	// success
	f := Cell(1.234).Native()
	assert.Equal(t, float64(1.234), f)
}

func TestString(t *testing.T) {
	// success - no decimals
	s := Cell(1).String()
	assert.Equal(t, "1", s)

	// success - with decimals
	s = Cell(1.234).String()
	assert.Equal(t, "1.234", s)
}
