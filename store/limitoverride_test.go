package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestLimitOverrideFilter(t *testing.T) {
	c := NewLimitOverrideCollection()
	limit := domain.Limit{Key: "builds", Value: 10}
	c.Create(LimitOverride{AppID: 1, Limit: domain.Limit{Key: "concurrency", Value: 10}})
	c.Create(LimitOverride{AppID: 1, Limit: domain.Limit{Key: "buildtime", Value: 10}})
	c.Create(LimitOverride{AppID: 2, Limit: limit})
	c.Create(LimitOverride{AppID: 3, Limit: domain.Limit{Key: "concurrency", Value: 10}})
	c.Create(LimitOverride{AppID: 3, Limit: domain.Limit{Key: "buildtime", Value: 10}})
	c.Create(LimitOverride{AppID: 3, Limit: domain.Limit{Key: "builds", Value: 10}})
	c.Create(LimitOverride{AppID: 4, Limit: domain.Limit{Key: "teammembers", Value: 10}})

	got := c.Filter(func(l LimitOverride) bool {
		return l.AppID == 2
	})

	assert.Equal(t, limit, got[0].Limit)

	got = c.Filter(func(l LimitOverride) bool {
		return l.AppID == 3
	})

	assert.Len(t, got, 3)
}
