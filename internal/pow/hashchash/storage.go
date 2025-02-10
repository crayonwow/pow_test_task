package hashchash

import (
	"sync"
)

func newInMemmoryStorage() *inMemmoryStorage {
	return &inMemmoryStorage{
		m: sync.Map{},
	}
}

type inMemmoryStorage struct {
	m sync.Map
}

func (in *inMemmoryStorage) Add(resource string) error {
	in.m.Store(resource, struct{}{})
	return nil
}

func (in *inMemmoryStorage) Spent(resource string) bool {
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
