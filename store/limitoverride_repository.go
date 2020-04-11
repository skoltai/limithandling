package store

// LimitOverrideRepository specifies the possible interactions with LimitOverride records
type LimitOverrideRepository interface {
	Filter(f func(l LimitOverride) bool) []LimitOverride
	Create(l LimitOverride) int
	Update(l LimitOverride) bool
}

// SimpleLimitOverrideRepository implements a simple, in-memory LimitOverrideRepository
type SimpleLimitOverrideRepository struct {
	store *MemoryStore
}

// NewSimpleLimitOverrideRepository is a constructor for SimpleLimitOverrideRepository
func NewSimpleLimitOverrideRepository(store *MemoryStore) LimitOverrideRepository {
	return &SimpleLimitOverrideRepository{store: store}
}

// Filter iterates over LimitOverride records and returns all elements that pass the test implemented by the provided function
func (r *SimpleLimitOverrideRepository) Filter(f func(l LimitOverride) bool) []LimitOverride {
	return r.store.LimitOverrides.filter(f)
}

// Create creates a LimitOverride record and returns with the ID
func (r *SimpleLimitOverrideRepository) Create(l LimitOverride) int {
	return r.store.LimitOverrides.create(l)
}

// Update updates a LimitOverride and returns wether the update was successful
func (r *SimpleLimitOverrideRepository) Update(l LimitOverride) bool {
	return r.store.LimitOverrides.update(l)
}
