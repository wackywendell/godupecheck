package dupechecker

import (
	"fmt"
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

var globalDeduped []string

type deduper func([]string) []string

func BenchmarkDups(b *testing.B) {
	counts := []int{1000, 10_000, 50_000}
	sizes := []int{20, 200}
	funcs := map[string]deduper{"arrays": DedupeArraywise, "mapped": DedupeMapwise}

	for _, s := range sizes {
		for _, c := range counts {
			strs := CreateStrings(c, s)

			for fname, f := range funcs {
				name := fmt.Sprintf("Dedupe-%s-s%d-n%d", fname, s, c)

				b.Run(name, func(pb *testing.B) {
					for i := 0; i < pb.N; i++ {
						globalDeduped = f(strs)
					}
				})
			}
		}
	}
}
