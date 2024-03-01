package sortvf

import (
	"sort"
)

// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}

// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, less})
}
