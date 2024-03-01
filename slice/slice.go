package slice

// Contains determines if a slice contains a given element.
func Contains[T comparable](elems []T, value T) bool {
	for _, e := range elems {
		if e == value {
			return true
		}
	}
	return false
}

// ContainsAtLeastOne checks if there's any common element between two slices.
func ContainsAtLeastOne[T comparable](sliceA, sliceB []T) bool {
	for _, a := range sliceA {
		for _, b := range sliceB {
			if a == b {
				return true
			}
		}
	}
	return false
}

// Remove removes a value from a slice and returns the modified slice.
func Remove[T comparable](elems []T, value T) []T {
	var result []T
	for _, e := range elems {
		if e != value {
			result = append(result, e)
		}
	}
	return result
}

// IndexOf returns the index of the element
func IndexOf[T comparable](elems []T, value T) int {
	for i, e := range elems {
		if e == value {
			return i
		}
	}
	return -1
}

// Intersection returns elements in both slices
func Intersection[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}

	var result []T
	for _, item := range b {
		if _, found := m[item]; found {
			result = append(result, item)
			delete(m, item)
		}
	}
	return result
}

// Union returns all elements of the two slices without duplicates
func Union[T comparable](a, b []T) []T {
	m := make(map[T]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		m[item] = true
	}

	var result []T
	for k := range m {
		result = append(result, k)
	}
	return result
}

// StrictEqual checks if two slices are strictly equals (order and content)
func StrictEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ContainSameElements checks if two slices content is the same (without order)
func ContainSameElements[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, val := range a {
		if !Contains(b, val) {
			return false
		}
	}
	return true
}

// Take returns the first n elements of a slice
func Take[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice[:]
	}
	if n <= 0 {
		return []T{}
	}
	return slice[:n]
}
