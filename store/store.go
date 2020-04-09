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

func (s *MemoryStore) GetUser(id int) (User, error) {
	return s.Users.Get(id)
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

func (s *MemoryStore) CreateApp(app App) int {
	return s.Apps.Create(app)
}

func (s *MemoryStore) GetApp(id int) (App, error) {
	return s.Apps.Get(id)
}

func (s *MemoryStore) UpdateApp(app App) bool {
	return s.Apps.Update(app)
}

func (s *MemoryStore) GetApps() map[int]App {
	return s.Apps.items
}

func (s *MemoryStore) FilterLimitOverrides(f func(l LimitOverride) bool) []LimitOverride {
	return s.LimitOverrides.Filter(f)
}

func (s *MemoryStore) CreateLimitOverride(l LimitOverride) int {
	return s.LimitOverrides.Create(l)
}

func (s *MemoryStore) UpdateLimitOverride(l LimitOverride) bool {
	return s.LimitOverrides.Update(l)
}

func (s *MemoryStore) GetPlan(id int) (Plan, error) {
	return s.Plans.Get(id)
}
