package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type App struct {
	ID             int
	OwnerID        int
	SubscriptionID int
	domain.App
}

type AppCollection struct {
	items  map[int]App
	nextID int
}

func NewAppCollection() *AppCollection {
	return &AppCollection{
		items:  make(map[int]App),
		nextID: 1,
	}
}

func (c *AppCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *AppCollection) Create(a App) int {
	id := c.makeID()
	a.ID = id
	c.items[id] = a
	return id
}

func (c *AppCollection) Get(id int) (App, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return App{}, errors.New("App not found")
}

func (c *AppCollection) Update(a App) bool {
	if _, ok := c.items[a.ID]; !ok {
		return false
	}

	c.items[a.ID] = a
	return true
}
