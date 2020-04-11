package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

// LimitOverride is a storage struct for Limit overrides
type LimitOverride struct {
	ID    int
	AppID int
	domain.Limit
}

type limitOverrideCollection struct {
	items  map[int]LimitOverride
	nextID int
}

func newLimitOverrideCollection() *limitOverrideCollection {
	return &limitOverrideCollection{
		items:  make(map[int]LimitOverride),
		nextID: 1,
	}
}

func (c *limitOverrideCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *limitOverrideCollection) create(l LimitOverride) int {
	id := c.makeID()
	l.ID = id
	c.items[id] = l
	return id
}

func (c *limitOverrideCollection) get(id int) (LimitOverride, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return LimitOverride{}, errors.New("LimitOverride not found")
}

func (c *limitOverrideCollection) filter(f func(l LimitOverride) bool) []LimitOverride {
	res := make([]LimitOverride, 0)
	for _, l := range c.items {
		if f(l) {
			res = append(res, l)
		}
	}

	return res
}

func (c *limitOverrideCollection) update(l LimitOverride) bool {
	if _, ok := c.items[l.ID]; !ok {
		return false
	}

	c.items[l.ID] = l
	return true
}
