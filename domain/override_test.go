package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByKey(t *testing.T) {
	haystack := []Limit{
		Limit{Key: "concurrency", Value: 1},
		Limit{Key: "buildtime", Value: 10},
		Limit{Key: "builds", Value: 200},
		Limit{Key: "teammembers", Value: 2},
		Limit{Key: "builds", Value: 200},
	}

	assert.Equal(t, &haystack[0], findByKey(haystack, "concurrency"))
	assert.Equal(t, &haystack[2], findByKey(haystack, "builds"))
}

func TestMergeOverrides(t *testing.T) {
	base := []Limit{
		Limit{Key: "concurrency", Value: 1},
		Limit{Key: "buildtime", Value: 10},
		Limit{Key: "builds", Value: 200},
		Limit{Key: "teammembers", Value: 2},
	}
	override := []Limit{
		Limit{Key: "concurrency", Value: 2},
		Limit{Key: "buildtime", Value: 45},
		Limit{Key: "builds", Value: 0},
		Limit{Key: "teammembers", Value: 0},
	}

	assert.Equal(t, override, MergeOverrides(base, override))
	assert.Equal(t, base, MergeOverrides(override, base))
	assert.Equal(t, override, MergeOverrides(override, override))
	assert.Equal(t, override, MergeOverrides([]Limit{}, override))
	assert.Equal(t, base, MergeOverrides(base, []Limit{}))
	assert.Equal(t, []Limit{}, MergeOverrides([]Limit{}, []Limit{}))

	want := []Limit{
		Limit{Key: "concurrency", Value: 2},
		Limit{Key: "buildtime", Value: 45},
		Limit{Key: "builds", Value: 200},
		Limit{Key: "teammembers", Value: 2},
	}

	assert.Equal(t, want, MergeOverrides(override[0:2], base[2:4]))

	want = []Limit{
		Limit{Key: "concurrency", Value: 1},
		Limit{Key: "buildtime", Value: 10},
		Limit{Key: "builds", Value: 0},
		Limit{Key: "teammembers", Value: 2},
	}

	assert.Equal(t, want, MergeOverrides(base, []Limit{Limit{Key: "builds", Value: 0}}))
}
