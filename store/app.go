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

func (c *AppCollection) Create(l App) int {
	id := c.makeID()
	l.ID = id
	c.items[id] = l
	return id
}

func (c *AppCollection) Get(id int) (App, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return App{}, errors.New("App not found")
}
