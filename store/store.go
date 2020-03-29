package store

import "github.com/skoltai/limithandling/domain"

type Store interface {
	GetUsers() []domain.User
	AddUser(user domain.User)
}

type MemoryStore struct {
	Users         []domain.User
	Subscriptions []domain.Subscription
	Plans         []domain.Plan
	Apps          []domain.App
}

func NewMemoryStore() Store {
	return &MemoryStore{
		Users: make([]domain.User, 0),
	}
}

func (s *MemoryStore) AddUser(user domain.User) {
	s.Users = append(s.Users, user)
}

func (s *MemoryStore) GetUsers() []domain.User {
	return s.Users
}
