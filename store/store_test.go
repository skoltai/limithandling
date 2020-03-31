package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	store := NewMemoryStore()

	user := domain.User{Username: "admin", Email: "admin@admin.com"}
	id := store.AddUser(user)
	got, _ := store.GetUser(id)

	assert.Equal(t, user, got)
}
