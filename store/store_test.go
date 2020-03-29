package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	store := NewMemoryStore()

	user := domain.User{Username: "admin", Email: "admin@admin.com"}
	store.AddUser(user)

	assert.Equal(t, store.GetUsers()[0], user)
}
