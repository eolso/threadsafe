package threadsafe

import (
	"sync"
)

type Slice[T any] struct {
	Data []T
	lock sync.RWMutex
}

// Append appends the value v into Slice.
func (s *Slice[T]) Append(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data, v)
}

func (s *Slice[T]) Insert(index uint, v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data[index] = v
}

func (s *Slice[T]) Get(index int) T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.Data[index]
}

func (s *Slice[T]) GetAll() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.Data
}

func (s *Slice[T]) SafeGet(index int) (T, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if index > len(s.Data) {
		return *new(T), false
	}

	return s.Data[index], true
}

// Delete deletes the item at index i. Delete will panic if i is out of bounds. If a panic is undesired, use SafeDelete.
func (s *Slice[T]) Delete(i int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:i], s.Data[i+1:]...)
}

func (s *Slice[T]) SafeDelete(i int) bool {
	if i < 0 || i >= len(s.Data) {
		return false
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:i], s.Data[i+1:]...)
	return true
}

func (s *Slice[T]) IndexFunc(f func(T) bool) int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for i, v := range s.Data {
		if f(v) {
			return i
		}
	}

	return -1
}

func (s *Slice[T]) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.Data)
}
