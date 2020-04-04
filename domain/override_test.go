package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByKey(t *testing.T) {
	haystack := []Limit{
		{Key: "concurrency", Value: 1},
		{Key: "buildtime", Value: 10},
		{Key: "builds", Value: 200},
		{Key: "teammembers", Value: 2},
		{Key: "builds", Value: 200},
	}

	assert.Equal(t, &haystack[0], findByKey(haystack, "concurrency"))
	assert.Equal(t, &haystack[2], findByKey(haystack, "builds"))
}

func TestMergeOverrides(t *testing.T) {
	base := []Limit{
		{Key: "concurrency", Value: 1},
		{Key: "buildtime", Value: 10},
		{Key: "builds", Value: 200},
		{Key: "teammembers", Value: 2},
	}
	override := []Limit{
		{Key: "concurrency", Value: 2},
		{Key: "buildtime", Value: 45},
		{Key: "builds", Value: 0},
		{Key: "teammembers", Value: 0},
	}

	assert.Equal(t, override, MergeOverrides(base, override))
	assert.Equal(t, base, MergeOverrides(override, base))
	assert.Equal(t, override, MergeOverrides(override, override))
	assert.Equal(t, override, MergeOverrides([]Limit{}, override))
	assert.Equal(t, base, MergeOverrides(base, []Limit{}))
	assert.Equal(t, []Limit{}, MergeOverrides([]Limit{}, []Limit{}))

	want := []Limit{
		{Key: "concurrency", Value: 2},
		{Key: "buildtime", Value: 45},
		{Key: "builds", Value: 200},
		{Key: "teammembers", Value: 2},
	}

	assert.Equal(t, want, MergeOverrides(override[0:2], base[2:4]))

	want = []Limit{
		{Key: "concurrency", Value: 1},
		{Key: "buildtime", Value: 10},
		{Key: "builds", Value: 0},
		{Key: "teammembers", Value: 2},
	}

	assert.Equal(t, want, MergeOverrides(base, []Limit{{Key: "builds", Value: 0}}))
}
