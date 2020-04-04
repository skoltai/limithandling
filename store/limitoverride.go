package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type LimitOverride struct {
	ID    int
	AppID int
	domain.Limit
}

type LimitOverrideCollection struct {
	items  map[int]LimitOverride
	nextID int
}

func NewLimitOverrideCollection() *LimitOverrideCollection {
	return &LimitOverrideCollection{
		items:  make(map[int]LimitOverride),
		nextID: 1,
	}
}

func (c *LimitOverrideCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *LimitOverrideCollection) Create(l LimitOverride) int {
	id := c.makeID()
	l.ID = id
	c.items[id] = l
	return id
}

func (c *LimitOverrideCollection) Get(id int) (LimitOverride, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return LimitOverride{}, errors.New("LimitOverride not found")
}

func (c *LimitOverrideCollection) Filter(f func(l LimitOverride) bool) []LimitOverride {
	res := make([]LimitOverride, 0)
	for _, l := range c.items {
		if f(l) {
			res = append(res, l)
		}
	}

	return res
}
