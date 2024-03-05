package sortvf

import "time"

func SortByDate[T any](s []T, getDate func(T) time.Time, isDesc bool) {
	SortFn(s, func(i, j T) bool {
		sortRes := getDate(i).Before(getDate(j))
		if isDesc {
			return !sortRes
		}
		return sortRes
	})
}

func SortByDateAsc[T any](s []T, getDate func(T) time.Time) {
	SortByDate(s, getDate, false)
}

func SortByDateDesc[T any](s []T, getDate func(T) time.Time) {
	SortByDate(s, getDate, true)
}
