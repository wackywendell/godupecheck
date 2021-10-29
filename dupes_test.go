package dupechecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduping(t *testing.T) {
	assert := assert.New(t)
	strs := []string{"a", "b", "a", "d", "c", "f", "c", "d"}
	origLength := len(strs)
	dedupeLength := 5

	dedupes1 := DedupeArraywise(strs)
	assert.Len(strs, origLength)
	assert.Len(dedupes1, dedupeLength)
	assert.NotEqual(dedupes1, strs)

	dedupes2 := DedupeMapwise(strs)
	assert.Len(strs, origLength)
	assert.Len(dedupes2, dedupeLength)
	assert.NotEqual(dedupes2, strs)

	assert.Equal(dedupes1, dedupes2)
}
