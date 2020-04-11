package store

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
