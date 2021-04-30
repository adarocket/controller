package user

import (
	"errors"
	"sync"
)

// UserStore -
type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

// InMemoryUserStore -
type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

// NewInMemoryUserStore -
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

// Save -
func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errors.New("ErrAlreadyExists")
	}

	store.users[user.Username] = user.Clone()
	return nil
}

// Find -
func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
