package threadsafe

import "sync"

// Map represents a generic map[comparable]any that locks itself on each operation. The underlying map Data is left
// exposed to not block any potential operations that might be needed, but should generally not be touched directly.
type Map[K comparable, V any] struct {
	Data map[K]V
	lock sync.Mutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		Data: make(map[K]V),
	}
}

// Get returns the value V at key K. Also returns a boolean representing if the value was found or not.
func (m *Map[K, V]) Get(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	v, ok := m.Data[key]

	return v, ok
}

// Pull behaves like Get but will also delete the key from the map before returning and unlocking the map. This can be
// useful for singleton operations.
func (m *Map[K, V]) Pull(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	v, ok := m.Data[key]
	if !ok {
		return v, ok
	}

	m.Delete(key)

	return v, ok
}

// Set writes the value V at key K.
func (m *Map[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Data[key] = value
}

// Delete deletes the key K, if it exists.
func (m *Map[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.Data[key]; ok {
		delete(m.Data, key)
	}
}

// Keys returns a slice of K keys.
func (m *Map[K, V]) Keys() []K {
	m.lock.Lock()
	defer m.lock.Unlock()

	keys := make([]K, len(m.Data))

	index := 0
	for k := range m.Data {
		keys[index] = k
		index++
	}

	return keys
}

// Values returns a slice V values.
func (m *Map[K, V]) Values() []V {
	m.lock.Lock()
	defer m.lock.Unlock()

	values := make([]V, len(m.Data))

	index := 0
	for _, v := range m.Data {
		values[index] = v
		index++
	}

	return values
}

// Items returns both the slice of keys and values.
func (m *Map[K, V]) Items() ([]K, []V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	keys := make([]K, len(m.Data))
	values := make([]V, len(m.Data))

	index := 0
	for k, v := range m.Data {
		keys[index] = k
		values[index] = v
		index++
	}

	return keys, values
}

// Empty deletes all keys in the map.
func (m *Map[K, V]) Empty() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Data = nil
	m.Data = make(map[K]V)
}

// Len returns the length of the map.
func (m *Map[K, V]) Len() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.Data)
}
