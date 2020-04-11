package store

import "github.com/skoltai/limithandling/domain"

// UserRepository specifies the possible interactions with User records
type UserRepository interface {
	Create(user domain.User) int
	Get(id int) (User, error)
}

// SimpleUserRepository implements a simple, in-memory UserRepository
type SimpleUserRepository struct {
	store *MemoryStore
}

// NewSimpleUserRepository is a constructor for SimpleUserRepository
func NewSimpleUserRepository(store *MemoryStore) UserRepository {
	return &SimpleUserRepository{store: store}
}

// Create creates a User record and returns with the ID
func (r *SimpleUserRepository) Create(user domain.User) int {
	return r.store.Users.create(user)
}

// Get retrieves a User by ID or returns an empty User and an error if it can't be found
func (r *SimpleUserRepository) Get(id int) (User, error) {
	return r.store.Users.get(id)
}
