package sortvf

import (
	"golang.org/x/exp/constraints"
)

func SortByOrderable[T any, U constraints.Ordered](s []T, getComparable func(T) U, isDesc bool) {
	SortFn(s, func(i, j T) bool {
		sortRes := getComparable(i) < getComparable(j)
		if isDesc {
			return !sortRes
		}
		return sortRes
	})
}

func SortByOrderableAsc[T any, U constraints.Ordered](s []T, getDate func(T) U) {
	SortByOrderable(s, getDate, false)
}

func SortByOrderableDesc[T any, U constraints.Ordered](s []T, getDate func(T) U) {
	SortByOrderable(s, getDate, true)
}
