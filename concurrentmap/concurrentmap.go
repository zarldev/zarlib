package concurrentmap

import "sync"

// Map type is a generic concurrent map.
type Map[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// NewHandlerMap returns a new map.
func NewHandlerMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		items: make(map[K]V),
	}
}

// Set sets the value for the given key.
func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.items[key] = value
	m.mu.Unlock()
}

// Get returns the value for the given key.
func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	value, ok := m.items[key]
	m.mu.RUnlock()
	return value, ok
}

// Delete deletes the value for the given key.
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.items, key)
	m.mu.Unlock()
}

// Len returns the number of items in the map.
func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	length := len(m.items)
	m.mu.RUnlock()
	return length
}

// ForEach calls f sequentially for each key and value present in the map.
// If f returns error, range stops the iteration.
func (m *Map[K, V]) ForEach(f func(key K, value V) error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.items {
		if err := f(k, v); err != nil {
			break
		}
	}
}
