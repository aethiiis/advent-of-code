package utils

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		elements: make(map[T]struct{}),
	}
}

func NewSetFromList[T comparable](list []T) Set[T] {
	elements := make(map[T]struct{})
	for _, e := range list {
		elements[e] = struct{}{}
	}
	return Set[T]{
		elements: elements,
	}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.elements[value]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Empty() bool {
	return len(s.elements) == 0
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) Get() T {
	var r T
	for e := range s.elements {
		r = e
		break
	}
	return r
}

func (s *Set[T]) Pop() T {
	var r T
	for e := range s.elements {
		r = e
		delete(s.elements, e)
		break
	}
	return r
}

func (s *Set[T]) List() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return &result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return &result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return &result
}
