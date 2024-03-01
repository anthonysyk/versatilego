package functional

func Zero[T any]() (res T) {
	return
}

func IsEmpty[T comparable](x T) bool {
	return x == Zero[T]()
}

func IsNotEmpty[T comparable](x T) bool {
	return x != Zero[T]()
}

func Identity[T any](x T) T {
	return x
}

type MapFunc[T, R any] func(T) R

// Map applies a map function to a slice.
func Map[T any, F any](data []T, getter MapFunc[T, F]) []F {
	var res []F
	for _, e := range data {
		res = append(res, getter(e))
	}
	return res
}

type FlatMapFunc[T, R any] func(T) []R

// FlatMap applies a flatmap function to a slice.
func FlatMap[T any, F any](data []T, getter FlatMapFunc[T, F]) []F {
	var res []F
	for _, e := range data {
		res = append(res, getter(e)...)
	}
	return res
}

// FilterFunc pe of the filter function.
type FilterFunc[T any] func(T) bool

// Filter applies a filter function to a slice.
func Filter[T any](data []T, filterFunc FilterFunc[T]) []T {
	var result []T
	for _, e := range data {
		if filterFunc(e) {
			result = append(result, e)
		}
	}
	return result
}

// ReduceFunc is the type of the reduce function.
type ReduceFunc[T any, R any] func(R, T) R

// Reduce applies a reduce function on a slice.
func Reduce[T any, R any](data []T, reducer ReduceFunc[T, R], initial R) R {
	result := initial
	for _, e := range data {
		result = reducer(result, e)
	}
	return result
}

type PredicateFunc[T any] func(T) bool

func Find[T any](data []T, findFunc PredicateFunc[T]) (T, bool) {
	for _, e := range data {
		if findFunc(e) {
			return e, true
		}
	}
	var empty T
	return empty, false
}

func Exists[T any](data []T, existsFunc PredicateFunc[T]) bool {
	for _, e := range data {
		if existsFunc(e) {
			return true
		}
	}
	return false
}
