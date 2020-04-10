package store

import (
	"github.com/skoltai/limithandling/domain"
)

type MemoryStore struct {
	Users          *userCollection
	Subscriptions  *subscriptionCollection
	Plans          *planCollection
	Apps           *appCollection
	LimitOverrides *limitOverrideCollection
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users:          newUserCollection(),
		Subscriptions:  newSubscriptionCollection(),
		Apps:           newAppCollection(),
		LimitOverrides: newLimitOverrideCollection(),
	}
}

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

type SubscriptionRepository interface {
	Get(id int) (Subscription, error)
	Create(sub Subscription) int
	Find(f func(Subscription) bool) (Subscription, bool)
}

type SimpleSubscriptionRepository struct {
	store *MemoryStore
}

func NewSimpleSubscriptionRepository(store *MemoryStore) SubscriptionRepository {
	return &SimpleSubscriptionRepository{store: store}
}

func (r *SimpleSubscriptionRepository) Get(id int) (Subscription, error) {
	return r.store.Subscriptions.get(id)
}

func (r *SimpleSubscriptionRepository) Create(sub Subscription) int {
	return r.store.Subscriptions.create(sub)
}

func (r *SimpleSubscriptionRepository) Find(f func(Subscription) bool) (Subscription, bool) {
	return r.store.Subscriptions.find(f)
}

type AppRepository interface {
	Create(app App) int
	Get(id int) (App, error)
	Update(app App) bool
}

type SimpleAppRepository struct {
	store *MemoryStore
}

func NewSimpleAppRepository(store *MemoryStore) AppRepository {
	return &SimpleAppRepository{store: store}
}

func (r *SimpleAppRepository) Create(app App) int {
	return r.store.Apps.create(app)
}

func (r *SimpleAppRepository) Get(id int) (App, error) {
	return r.store.Apps.get(id)
}

func (r *SimpleAppRepository) Update(app App) bool {
	return r.store.Apps.update(app)
}

func (s *MemoryStore) GetApps() map[int]App {
	return s.Apps.items
}

type LimitOverrideRepository interface {
	Filter(f func(l LimitOverride) bool) []LimitOverride
	Create(l LimitOverride) int
	Update(l LimitOverride) bool
}

type SimpleLimitOverrideRepository struct {
	store *MemoryStore
}

func NewSimpleLimitOverrideRepository(store *MemoryStore) LimitOverrideRepository {
	return &SimpleLimitOverrideRepository{store: store}
}
func (r *SimpleLimitOverrideRepository) Filter(f func(l LimitOverride) bool) []LimitOverride {
	return r.store.LimitOverrides.filter(f)
}

func (r *SimpleLimitOverrideRepository) Create(l LimitOverride) int {
	return r.store.LimitOverrides.create(l)
}

func (r *SimpleLimitOverrideRepository) Update(l LimitOverride) bool {
	return r.store.LimitOverrides.update(l)
}

type PlanRepository interface {
	Get(id int) (Plan, error)
}

type SimplePlanRepository struct {
	store *MemoryStore
}

func NewSimplePlanRepository(store *MemoryStore) PlanRepository {
	return &SimplePlanRepository{store: store}
}

func (r *SimplePlanRepository) Get(id int) (Plan, error) {
	return r.store.Plans.get(id)
}
