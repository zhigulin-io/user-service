package user

import (
	"errors"
	"sync"
)

type InMemoryStorage struct {
	storage sync.Map
}

func NewInMemStorage() *InMemoryStorage {
	return &InMemoryStorage{
		storage: sync.Map{},
	}
}

func (s *InMemoryStorage) Write(username string, password string) error {
	_, loaded := s.storage.LoadOrStore(username, password)

	if loaded {
		return errors.New("username already exists")
	}

	return nil
}
