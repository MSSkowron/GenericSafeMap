package genericsafemap

import "sync"

type GenericSafeMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func New[K comparable, V any]() *GenericSafeMap[K, V] {
	return &GenericSafeMap[K, V]{
		data: make(map[K]V),
		mu:   sync.RWMutex{},
	}
}

func (m *GenericSafeMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	if !ok {
		return value, false
	}

	return value, true
}

func (m *GenericSafeMap[K, V]) Put(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *GenericSafeMap[K, V]) Remove(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
}
