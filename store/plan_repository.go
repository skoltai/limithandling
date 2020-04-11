package store

// PlanRepository specifies the possible interactions with Plan records
type PlanRepository interface {
	Get(id int) (Plan, error)
}

// SimplePlanRepository implements a simple, in-memory PlanRepository
type SimplePlanRepository struct {
	store *MemoryStore
}

// NewSimplePlanRepository is a constructor for SimplePlanRepository
func NewSimplePlanRepository(store *MemoryStore) PlanRepository {
	return &SimplePlanRepository{store: store}
}

// Get retrieves a Plan by ID or returns an empty Plan and an error if it can't be found
func (r *SimplePlanRepository) Get(id int) (Plan, error) {
	return r.store.Plans.get(id)
}
