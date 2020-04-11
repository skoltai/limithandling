package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

// App is a storage struct for Apps
type App struct {
	ID             int
	OwnerID        int
	SubscriptionID int
	domain.App
}

type appCollection struct {
	items  map[int]App
	nextID int
}

func newAppCollection() *appCollection {
	return &appCollection{
		items:  make(map[int]App),
		nextID: 1,
	}
}

func (c *appCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *appCollection) create(a App) int {
	id := c.makeID()
	a.ID = id
	c.items[id] = a
	return id
}

func (c *appCollection) get(id int) (App, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return App{}, errors.New("App not found")
}

func (c *appCollection) update(a App) bool {
	if _, ok := c.items[a.ID]; !ok {
		return false
	}

	c.items[a.ID] = a
	return true
}
