package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestLimitOverrideFilter(t *testing.T) {
	c := newLimitOverrideCollection()
	limits := []*LimitOverride{
		{AppID: 1, Limit: domain.Limit{Key: "concurrency", Value: 10}},
		{AppID: 1, Limit: domain.Limit{Key: "buildtime", Value: 10}},
		{AppID: 2, Limit: domain.Limit{Key: "builds", Value: 10}},
		{AppID: 3, Limit: domain.Limit{Key: "concurrency", Value: 10}},
		{AppID: 3, Limit: domain.Limit{Key: "buildtime", Value: 10}},
		{AppID: 3, Limit: domain.Limit{Key: "builds", Value: 10}},
		{AppID: 4, Limit: domain.Limit{Key: "teammembers", Value: 10}},
	}
	for _, l := range limits {
		l.ID = c.create(*l)
	}

	got := c.filter(func(l LimitOverride) bool {
		return l.AppID == 2
	})

	assert.Equal(t, *limits[2], got[0])

	got = c.filter(func(l LimitOverride) bool {
		return l.AppID == 3
	})

	assert.ElementsMatch(t, []LimitOverride{*limits[3], *limits[4], *limits[5]}, got)
}

func TestLimitOverrideUpdate(t *testing.T) {
	c := newLimitOverrideCollection()
	assert.False(t, c.update(LimitOverride{ID: 0}))
	assert.False(t, c.update(LimitOverride{ID: 1}))

	lo := LimitOverride{AppID: 1, Limit: domain.Limit{Key: "concurrency", Value: 10}}
	lo.ID = c.create(lo)

	got, _ := c.get(lo.ID)
	assert.Equal(t, lo, got)

	want := LimitOverride{ID: lo.ID, AppID: 1, Limit: domain.Limit{Key: "test", Value: 1}}
	ok := c.update(want)
	assert.True(t, ok)
	got, _ = c.get(want.ID)
	assert.Equal(t, want, got)
}
