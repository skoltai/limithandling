package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

// Subscription is a storage struct for Subscriptions
type Subscription struct {
	ID     int
	UserID int
	PlanID int
	domain.Subscription
}

type subscriptionCollection struct {
	items  map[int]Subscription
	nextID int
}

func newSubscriptionCollection() *subscriptionCollection {
	return &subscriptionCollection{
		items:  make(map[int]Subscription),
		nextID: 1,
	}
}

func (c *subscriptionCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *subscriptionCollection) create(s Subscription) int {
	id := c.makeID()
	s.ID = id
	c.items[id] = s
	return id
}

func (c *subscriptionCollection) get(id int) (Subscription, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return Subscription{}, errors.New("Subscription not found")
}

func (c *subscriptionCollection) find(f func(Subscription) bool) (Subscription, bool) {
	for _, sub := range c.items {
		if f(sub) {
			return sub, true
		}
	}

	return Subscription{}, false
}
