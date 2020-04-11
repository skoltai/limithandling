package store

// SubscriptionRepository specifies the possible interactions with Subscription records
type SubscriptionRepository interface {
	Get(id int) (Subscription, error)
	Create(sub Subscription) int
	Find(f func(Subscription) bool) (Subscription, bool)
}

// SimpleSubscriptionRepository implements a simple, in-memory SubscriptionRepository
type SimpleSubscriptionRepository struct {
	store *MemoryStore
}

// NewSimpleSubscriptionRepository is a constructor for SimpleSubscriptionRepository
func NewSimpleSubscriptionRepository(store *MemoryStore) SubscriptionRepository {
	return &SimpleSubscriptionRepository{store: store}
}

// Get retrieves a Subscription by ID or returns an empty Subscription and an error if it can't be found
func (r *SimpleSubscriptionRepository) Get(id int) (Subscription, error) {
	return r.store.Subscriptions.get(id)
}

// Create creates a Subscription record and returns with the ID
func (r *SimpleSubscriptionRepository) Create(sub Subscription) int {
	return r.store.Subscriptions.create(sub)
}

// Find iterates over Subscription records and returns the first where f(Subscription) is true, or an empty Subscription and a bool indicating failure
func (r *SimpleSubscriptionRepository) Find(f func(Subscription) bool) (Subscription, bool) {
	return r.store.Subscriptions.find(f)
}
