package store

import "github.com/skoltai/limithandling/domain"

// AppRepository specifies the possible interactions with Apps
type AppRepository interface {
	Create(app App) int
	Get(id int) (App, error)
	Update(app App) bool
	All() []App
	LimitOverrides(id int) []domain.Limit
}

// SimpleAppRepository implements a simple, in-memory AppRepository
type SimpleAppRepository struct {
	store *MemoryStore
}

// NewSimpleAppRepository is a constructor for SimpleAppRepository
func NewSimpleAppRepository(store *MemoryStore) AppRepository {
	return &SimpleAppRepository{store: store}
}

// Create creates an App record and returns with the ID
func (r *SimpleAppRepository) Create(app App) int {
	return r.store.Apps.create(app)
}

// Get retrieves an App by ID or returns an empty App and an error if it can't be found
func (r *SimpleAppRepository) Get(id int) (App, error) {
	return r.store.Apps.get(id)
}

// Update updates an App and returns wether the update was successful
func (r *SimpleAppRepository) Update(app App) bool {
	return r.store.Apps.update(app)
}

// All returns all stored Apps for testing purposes
func (r *SimpleAppRepository) All() []App {
	apps := make([]App, 0)
	for _, a := range r.store.Apps.items {
		apps = append(apps, a)
	}
	return apps
}

// LimitOverrides retrieves LimitOverrides by app ID
func (r *SimpleAppRepository) LimitOverrides(id int) []domain.Limit {
	limits := r.store.LimitOverrides.filter(func(l LimitOverride) bool {
		return l.AppID == id
	})
	res := make([]domain.Limit, 0)
	for _, l := range limits {
		res = append(res, l.Limit)
	}

	return res
}
