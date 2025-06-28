package name

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	n := New("NAME")
	assert.Equal(t, Name("name"), n)
}

func TestParse(t *testing.T) {
	// success
	n, err := Parse("\tname\n")
	assert.Equal(t, Name("name"), n)
	assert.NoError(t, err)

	// error - invalid Name
	n, err = Parse("\n")
	assert.Empty(t, n)
	assert.EqualError(t, err, `invalid name ""`)
}

func TestBool(t *testing.T) {
	// success - true
	b := Name("name").Bool()
	assert.True(t, b)

	// success - false
	b = Name("").Bool()
	assert.False(t, b)
}

func TestNative(t *testing.T) {
	// success
	s := Name("name").Native()
	assert.Equal(t, "name", s)
}

func TestString(t *testing.T) {
	// success
	s := Name("name").String()
	assert.Equal(t, "name", s)
}
