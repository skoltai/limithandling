package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserCollection(t *testing.T) {
	c := NewUserCollection()

	user := domain.User{
		Username: "test1",
	}

	id := c.Create(user)

	got, err := c.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, user, got.User)

	_, err = c.Get(2)
	assert.Error(t, err)
}
