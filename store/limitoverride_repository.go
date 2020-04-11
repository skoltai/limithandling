package store

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
