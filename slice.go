package threadsafe

import (
	"sync"
)

type Slice[T any] struct {
	Data []T
	lock sync.Mutex
}

// Append appends the value v into Slice.
func (s *Slice[T]) Append(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data, v)
}

func (s *Slice[T]) Insert(index int, v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:index], append([]T{v}, s.Data[index:]...)...)
}

func (s *Slice[T]) SafeInsert(index int, v T) bool {
	if index < 0 || index > len(s.Data) {
		return false
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:index], append([]T{v}, s.Data[index:]...)...)

	return true
}

func (s *Slice[T]) Replace(index int, v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data[index] = v
}

func (s *Slice[T]) SafeReplace(index int, v T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if index < 0 || index >= len(s.Data) {
		return false
	}

	s.Data[index] = v

	return true
}

func (s *Slice[T]) Get(index int) T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.Data[index]
}

func (s *Slice[T]) GetAll() []T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.Data
}

func (s *Slice[T]) SafeGet(index int) (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if index < 0 || index >= len(s.Data) {
		return *new(T), false
	}

	return s.Data[index], true
}

// Delete deletes the item at index i. Delete will panic if i is out of bounds. If a panic is undesired, use SafeDelete.
func (s *Slice[T]) Delete(index int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:index], s.Data[index+1:]...)
}

func (s *Slice[T]) SafeDelete(index int) bool {
	if index < 0 || index >= len(s.Data) {
		return false
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data = append(s.Data[:index], s.Data[index+1:]...)
	return true
}

func (s *Slice[T]) Empty() {
	s.lock.Lock()
	s.Data = nil
	s.lock.Unlock()
}

func (s *Slice[T]) IndexFunc(f func(T) bool) int {
	s.lock.Lock()
	defer s.lock.Unlock()

	for i, v := range s.Data {
		if f(v) {
			return i
		}
	}

	return -1
}

func (s *Slice[T]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.Data)
}
