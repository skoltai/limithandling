package store

import "github.com/skoltai/limithandling/domain"

type UserRepository interface {
	Create(user domain.User) int
	Get(id int) (User, error)
}

type SimpleUserRepository struct {
	store *MemoryStore
}

func NewSimpleUserRepository(store *MemoryStore) UserRepository {
	return &SimpleUserRepository{store: store}
}

func (r *SimpleUserRepository) Create(user domain.User) int {
	return r.store.Users.create(user)
}

func (r *SimpleUserRepository) Get(id int) (User, error) {
	return r.store.Users.get(id)
}
