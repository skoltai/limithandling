package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	store := NewMemoryStore()
	r := NewSimpleUserRepository(store)

	user := domain.User{Username: "admin", Email: "admin@admin.com"}
	id := r.Create(user)
	got, _ := r.Get(id)

	assert.Equal(t, user, got.User)
}
