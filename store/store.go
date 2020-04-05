package store

import "github.com/skoltai/limithandling/domain"

type Store interface {
	AddUser(user domain.User) int
	GetUser(id int) (domain.User, error)
	GetSubscription(id int) (Subscription, error)
	CreateSubscription(sub Subscription) int
	FindSubscription(f func(Subscription) bool) (Subscription, bool)
}

type MemoryStore struct {
	Users          *UserCollection
	Subscriptions  *SubscriptionCollection
	Plans          *PlanCollection
	Apps           *AppCollection
	LimitOverrides *LimitOverrideCollection
}

func NewMemoryStore() Store {
	return &MemoryStore{
		Users:          NewUserCollection(),
		Subscriptions:  NewSubscriptionCollection(),
		Apps:           NewAppCollection(),
		LimitOverrides: NewLimitOverrideCollection(),
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

func (s *MemoryStore) GetSubscription(id int) (Subscription, error) {
	return s.Subscriptions.Get(id)
}

func (s *MemoryStore) CreateSubscription(sub Subscription) int {
	return s.Subscriptions.Create(sub)
}

func (s *MemoryStore) FindSubscription(f func(Subscription) bool) (Subscription, bool) {
	return s.Subscriptions.Find(f)
}
