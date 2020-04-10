package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type User struct {
	domain.User
	ID int
}

type userCollection struct {
	items  map[int]User
	nextID int
}

func newUserCollection() *userCollection {
	return &userCollection{
		items:  make(map[int]User),
		nextID: 1,
	}
}

func (c *userCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *userCollection) create(user domain.User) int {
	id := c.makeID()
	c.items[id] = User{User: user, ID: id}
	return id
}

func (c *userCollection) get(id int) (User, error) {
	if i, ok := c.items[id]; ok {
		return i, nil
	}
	return User{}, errors.New("User not found")
}
