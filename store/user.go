package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type User struct {
	domain.User
	ID int
}

type UserCollection struct {
	users  map[int]User
	nextID int
}

func NewUserCollection() *UserCollection {
	return &UserCollection{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (c *UserCollection) makeID() int {
	defer func() {
		c.nextID++
	}()
	return c.nextID
}

func (c *UserCollection) Create(user domain.User) int {
	id := c.makeID()
	c.users[id] = User{User: user, ID: id}
	return id
}

func (c *UserCollection) Get(id int) (User, error) {
	if u, ok := c.users[id]; ok {
		return u, nil
	}
	return User{}, errors.New("User not found")
}
