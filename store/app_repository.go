package store

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
