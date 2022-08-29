package threadsafe

import "sync"

type Map[K comparable, V any] struct {
	Data map[K]V
	lock sync.RWMutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		Data: make(map[K]V),
	}
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	v, ok := m.Data[key]

	return v, ok
}

func (m *Map[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Data[key] = value
}

func (m *Map[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.Data[key]; ok {
		delete(m.Data, key)
	}
}

func (m *Map[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()

	keys := make([]K, len(m.Data))

	index := 0
	for k := range m.Data {
		keys[index] = k
		index++
	}

	return keys
}

func (m *Map[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	values := make([]V, len(m.Data))

	index := 0
	for _, v := range m.Data {
		values[index] = v
		index++
	}

	return values
}

func (m *Map[K, V]) Items() ([]K, []V) {
	m.lock.RLock()
	defer m.lock.RUnlock()

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

func (m *Map[K, V]) Empty() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Data = nil
	m.Data = make(map[K]V)
}
