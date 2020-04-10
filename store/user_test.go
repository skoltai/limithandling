package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserCollection(t *testing.T) {
	c := newUserCollection()

	user := domain.User{
		Username: "test1",
	}

	id := c.create(user)

	got, err := c.get(id)
	assert.NoError(t, err)
	assert.Equal(t, user, got.User)

	_, err = c.get(2)
	assert.Error(t, err)
}
