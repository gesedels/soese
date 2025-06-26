package name

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	name := New("NAME")
	assert.Equal(t, Name("name"), name)
}

func TestParse(t *testing.T) {
	// success
	name, err := Parse("\tname\n")
	assert.Equal(t, Name("name"), name)
	assert.NoError(t, err)

	// error - invalid Name
	name, err = Parse("\n")
	assert.Empty(t, name)
	assert.EqualError(t, err, `invalid Name ""`)
}

func TestBool(t *testing.T) {
	// success - true
	okay := Name("name").Bool()
	assert.True(t, okay)

	// success - false
	okay = Name("").Bool()
	assert.False(t, okay)
}

func TestNative(t *testing.T) {
	// success
	text := Name("name").Native()
	assert.Equal(t, "name", text)
}

func TestString(t *testing.T) {
	// success
	text := Name("name").String()
	assert.Equal(t, "name", text)
}
