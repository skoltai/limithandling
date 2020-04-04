package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type Subscription struct {
	ID     int
	UserID int
	PlanID int
	domain.Subscription
}

type SubscriptionCollection struct {
	items  map[int]Subscription
	nextID int
}

func NewSubscriptionCollection() *SubscriptionCollection {
	return &SubscriptionCollection{
		items:  make(map[int]Subscription),
		nextID: 1,
	}
}

func (c *SubscriptionCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *SubscriptionCollection) Create(s Subscription) int {
	id := c.makeID()
	s.ID = id
	c.items[id] = s
	return id
}

func (c *SubscriptionCollection) Get(id int) (Subscription, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return Subscription{}, errors.New("Subscription not found")
}
