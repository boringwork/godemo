package testing

import (
	"testing"

	"gotest.tools/assert"
)

func TestEverything(t *testing.T) {
	// booleans
	assert.Assert(t, false)
	// assert.Assert(t, !missing)

	// // primitives
	// assert.Equal(t, count, 1)
	// assert.Equal(t, msg, "the message")
	// assert.Assert(t, total != 10) // NotEqual

	// // errors
	// assert.NilError(t, closer.Close())
	// assert.Error(t, err, "the exact error message")
	// assert.ErrorContains(t, err, "includes this")
	// assert.ErrorType(t, err, os.IsNotExist)

	// // complex types
	// // assert.DeepEqual(t, result, myStruct{Name: "title"})
	// assert.Assert(t, is.Len(items, 3))
	// assert.Assert(t, len(sequence) != 0) // NotEmpty
	// assert.Assert(t, is.Contains(mapping, "key"))

	// pointers and interface
	// assert.Assert(t, is.Nil(ref))
	// assert.Assert(t, ref != nil) // NotNil
}
