package store

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
