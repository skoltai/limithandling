package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

// LimitOverrideRepository specifies the possible interactions with LimitOverride records
type LimitOverrideRepository interface {
	Create(l LimitOverride) int
	Update(l LimitOverride) bool
	Upsert(appID int, limit domain.Limit) error
}

// SimpleLimitOverrideRepository implements a simple, in-memory LimitOverrideRepository
type SimpleLimitOverrideRepository struct {
	store *MemoryStore
}

// NewSimpleLimitOverrideRepository is a constructor for SimpleLimitOverrideRepository
func NewSimpleLimitOverrideRepository(store *MemoryStore) LimitOverrideRepository {
	return &SimpleLimitOverrideRepository{store: store}
}

// Create creates a LimitOverride record and returns with the ID
func (r *SimpleLimitOverrideRepository) Create(l LimitOverride) int {
	return r.store.LimitOverrides.create(l)
}

// Update updates a LimitOverride and returns wether the update was successful
func (r *SimpleLimitOverrideRepository) Update(l LimitOverride) bool {
	return r.store.LimitOverrides.update(l)
}

// Upsert updates or creates a LimitOverride for an App
func (r *SimpleLimitOverrideRepository) Upsert(appID int, limit domain.Limit) error {
	limits := r.store.LimitOverrides.filter(func(l LimitOverride) bool {
		return l.AppID == appID && l.Limit.Key == limit.Key
	})

	switch len(limits) {
	case 0:
		r.Create(LimitOverride{AppID: appID, Limit: limit})
	case 1:
		r.Update(LimitOverride{ID: limits[0].ID, AppID: appID, Limit: limit})
	default:
		// This should never happen
		return errors.New("duplicated limit override")
	}

	return nil
}
