package simpleset

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Set[T constraints.Ordered] struct {
	contents map[T]any
}

func New[T constraints.Ordered]() *Set[T] {
	return &Set[T]{contents: make(map[T]any)}
}

func NewFromSlice[T constraints.Ordered](slice []T) *Set[T] {
	set := New[T]()
	set.AddSlice(slice)
	return set
}

func (s *Set[T]) Add(t T) (added bool) {
	if _, found := s.contents[t]; !found {
		added = true
		s.contents[t] = struct{}{}
	}
	return
}

func (s *Set[T]) AddSlice(slice []T) (added []bool) {
	added = make([]bool, len(slice))
	for i := range slice {
		t := slice[i]
		added[i] = s.Add(t)
	}
	return
}

func (s *Set[T]) Contains(t T) (found bool) {
	_, found = s.contents[t]
	return
}

func (s *Set[T]) ContainsSlice(slice []T) (found []bool) {
	found = make([]bool, len(slice))
	for i := range slice {
		t := slice[i]
		found[i] = s.Contains(t)
	}
	return
}

func (s *Set[T]) Drop(t T) (dropped bool) {
	if _, found := s.contents[t]; found {
		dropped = true
		delete(s.contents, t)
	}
	return
}

func (s *Set[T]) DropSlice(slice []T) (dropped []bool) {
	dropped = make([]bool, len(slice))
	for i := range slice {
		t := slice[i]
		dropped[i] = s.Drop(t)
	}
	return
}

func (s *Set[T]) Len() int { return len(s.contents) }
func (s *Set[T]) Reset()   { s.contents = make(map[T]any) }

func (s *Set[T]) Slice() (contents []T) {
	contents = make([]T, 0, s.Len())
	for key := range s.contents {
		contents = append(contents, key)
	}
	slices.Sort(contents)
	return contents
}
