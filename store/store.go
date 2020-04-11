package store

// MemoryStore represents a data store for related records of all used types
type MemoryStore struct {
	Users          *userCollection
	Subscriptions  *subscriptionCollection
	Plans          *planCollection
	Apps           *appCollection
	LimitOverrides *limitOverrideCollection
}

// NewMemoryStore is a constructor for MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users:          newUserCollection(),
		Subscriptions:  newSubscriptionCollection(),
		Apps:           newAppCollection(),
		LimitOverrides: newLimitOverrideCollection(),
	}
}

// GetApps returns all stored Apps for testing purposes
func (s *MemoryStore) GetApps() map[int]App {
	return s.Apps.items
}
