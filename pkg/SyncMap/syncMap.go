package SyncMap

import (
	"errors"
	"sync"
)

type SyncMap struct {
	mu sync.Mutex
	m  map[any]any
}

func (m *SyncMap) Add(key, value any) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *SyncMap) Delete(key any) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *SyncMap) GetValues() []any {
	var array []any
	m.mu.Lock()
	for key := range m.m {
		array = append(array, m.m[key])
	}
	m.mu.Unlock()
	return array
}

func (m *SyncMap) Get(key any) (any, error) {
	m.mu.Lock()
	value, exist := m.m[key]
	if !exist {
		m.mu.Unlock()
		return nil, errors.New("Key missed")
	}
	m.mu.Unlock()
	return value, nil
}

func NewSyncMap() *SyncMap {
	return &SyncMap{m: make(map[any]any), mu: sync.Mutex{}}
}
