package store

import "github.com/skoltai/limithandling/domain"

type Store interface {
	AddUser(user domain.User) int
	GetUser(id int) (domain.User, error)
}

type MemoryStore struct {
	Users         *UserCollection
	Subscriptions []domain.Subscription
	Plans         []domain.Plan
	Apps          []domain.App
}

func NewMemoryStore() Store {
	return &MemoryStore{
		Users: NewUserCollection(),
	}
}

func (s *MemoryStore) AddUser(user domain.User) int {
	return s.Users.Create(user)
}

func (s *MemoryStore) GetUser(id int) (domain.User, error) {
	u, err := s.Users.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	return u.User, err
}
