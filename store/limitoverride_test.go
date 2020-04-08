package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestLimitOverrideFilter(t *testing.T) {
	c := NewLimitOverrideCollection()
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
		l.ID = c.Create(*l)
	}

	got := c.Filter(func(l LimitOverride) bool {
		return l.AppID == 2
	})

	assert.Equal(t, *limits[2], got[0])

	got = c.Filter(func(l LimitOverride) bool {
		return l.AppID == 3
	})

	assert.ElementsMatch(t, []LimitOverride{*limits[3], *limits[4], *limits[5]}, got)
}

func TestLimitOverrideUpdate(t *testing.T) {
	c := NewLimitOverrideCollection()
	assert.False(t, c.Update(LimitOverride{ID: 0}))
	assert.False(t, c.Update(LimitOverride{ID: 1}))

	lo := LimitOverride{AppID: 1, Limit: domain.Limit{Key: "concurrency", Value: 10}}
	lo.ID = c.Create(lo)

	got, _ := c.Get(lo.ID)
	assert.Equal(t, lo, got)

	want := LimitOverride{ID: lo.ID, AppID: 1, Limit: domain.Limit{Key: "test", Value: 1}}
	ok := c.Update(want)
	assert.True(t, ok)
	got, _ = c.Get(want.ID)
	assert.Equal(t, want, got)
}
