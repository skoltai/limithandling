package store

// AppRepository specifies the possible interactions with Apps
type AppRepository interface {
	Create(app App) int
	Get(id int) (App, error)
	Update(app App) bool
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
