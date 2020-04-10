package store

import (
	"github.com/skoltai/limithandling/domain"
)

type TestHelpers interface {
	GetApps() map[int]App
}

type Store interface {
	AddUser(user domain.User) int
	GetUser(id int) (User, error)
	GetSubscription(id int) (Subscription, error)
	CreateSubscription(sub Subscription) int
	FindSubscription(f func(Subscription) bool) (Subscription, bool)
	CreateApp(app App) int
	GetApp(id int) (App, error)
	UpdateApp(app App) bool
	FilterLimitOverrides(f func(l LimitOverride) bool) []LimitOverride
	CreateLimitOverride(l LimitOverride) int
	UpdateLimitOverride(l LimitOverride) bool
	GetPlan(id int) (Plan, error)
	TestHelpers
}

type MemoryStore struct {
	Users          *userCollection
	Subscriptions  *subscriptionCollection
	Plans          *planCollection
	Apps           *appCollection
	LimitOverrides *limitOverrideCollection
}

func NewMemoryStore() Store {
	return &MemoryStore{
		Users:          newUserCollection(),
		Subscriptions:  newSubscriptionCollection(),
		Apps:           newAppCollection(),
		LimitOverrides: newLimitOverrideCollection(),
	}
}

func (s *MemoryStore) AddUser(user domain.User) int {
	return s.Users.create(user)
}

func (s *MemoryStore) GetUser(id int) (User, error) {
	return s.Users.get(id)
}

func (s *MemoryStore) GetSubscription(id int) (Subscription, error) {
	return s.Subscriptions.get(id)
}

func (s *MemoryStore) CreateSubscription(sub Subscription) int {
	return s.Subscriptions.create(sub)
}

func (s *MemoryStore) FindSubscription(f func(Subscription) bool) (Subscription, bool) {
	return s.Subscriptions.find(f)
}

func (s *MemoryStore) CreateApp(app App) int {
	return s.Apps.create(app)
}

func (s *MemoryStore) GetApp(id int) (App, error) {
	return s.Apps.get(id)
}

func (s *MemoryStore) UpdateApp(app App) bool {
	return s.Apps.update(app)
}

func (s *MemoryStore) GetApps() map[int]App {
	return s.Apps.items
}

func (s *MemoryStore) FilterLimitOverrides(f func(l LimitOverride) bool) []LimitOverride {
	return s.LimitOverrides.filter(f)
}

func (s *MemoryStore) CreateLimitOverride(l LimitOverride) int {
	return s.LimitOverrides.create(l)
}

func (s *MemoryStore) UpdateLimitOverride(l LimitOverride) bool {
	return s.LimitOverrides.update(l)
}

func (s *MemoryStore) GetPlan(id int) (Plan, error) {
	return s.Plans.get(id)
}
