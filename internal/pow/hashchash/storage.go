package hashchash

import (
	"sync"
)

func NewInMemmoryStorage() *InMemmoryStorage {
	return &InMemmoryStorage{
		m: &sync.Map{},
	}
}

type InMemmoryStorage struct {
	m *sync.Map
}

func (in *InMemmoryStorage) Add(resource string) error {
	in.m.Store(resource, struct{}{})
	return nil
}

func (in *InMemmoryStorage) Spent(resource string) bool {
	_, ok := in.m.Load(resource)
	return ok
}

type MockStorage struct{}

func (mo *MockStorage) Add(_ string) error {
	return nil
}

func (mo *MockStorage) Spent(_ string) bool {
	return false
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}
