package slice

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
	}

	for _, tt := range tests {
		result := Contains(tt.slice, tt.value)
		if result != tt.expected {
			t.Errorf("SliceContains(%v, %d) = %v; want %v", tt.slice, tt.value, result, tt.expected)
		}
	}
}

func TestSliceContainsAtLeastOne(t *testing.T) {
	tests := []struct {
		sliceA []string
		sliceB []string
		result bool
	}{
		{[]string{"a", "b", "c"}, []string{"c", "d", "e"}, true},
		{[]string{"a", "b", "c"}, []string{"d", "e", "f"}, false},
	}

	for _, tt := range tests {
		contains := ContainsAtLeastOne(tt.sliceA, tt.sliceB)
		if contains != tt.result {
			t.Errorf("SliceContainsAtLeastOne(%v, %v) = %v; want %v", tt.sliceA, tt.sliceB, contains, tt.result)
		}
	}
}

func TestSliceRemove(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		result := Remove(tt.slice, tt.value)
		if !StrictEqual(result, tt.expected) {
			t.Errorf("SliceRemove(%v, %d) = %v; want %v", tt.slice, tt.value, result, tt.expected)
		}
	}
}

func TestSliceIndexOf(t *testing.T) {
	tests := []struct {
		slice    []int
		value    int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, tt := range tests {
		result := IndexOf(tt.slice, tt.value)
		if result != tt.expected {
			t.Errorf("expected %d but got %d", tt.expected, result)
		}
	}
}

func TestSliceIntersection(t *testing.T) {
	tests := []struct {
		sliceA []string
		sliceB []string
		result []string
	}{
		{[]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "c"}},
	}

	for _, tt := range tests {
		intersection := Intersection(tt.sliceA, tt.sliceB)
		if !StrictEqual(intersection, tt.result) {
			t.Errorf("expected %v but got %v", tt.result, intersection)
		}
	}
}

func TestSliceUnion(t *testing.T) {
	tests := []struct {
		sliceA []int
		sliceB []int
		result []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		union := Union(tt.sliceA, tt.sliceB)
		if !ContainSameElements(union, tt.result) {
			t.Errorf("expected %v but got %v", tt.result, union)
		}
	}
}

func TestStrictEqual(t *testing.T) {
	tests := []struct {
		sliceA []int
		sliceB []int
		result bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
	}

	for _, tt := range tests {
		strictEqual := StrictEqual(tt.sliceA, tt.sliceB)
		if strictEqual != tt.result {
			t.Errorf("expected %v but got %v", tt.result, strictEqual)
		}
	}
}

func TestContainSameElements(t *testing.T) {
	tests := []struct {
		sliceA []int
		sliceB []int
		result bool
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{1, 2, 5}, false},
	}

	for _, tt := range tests {
		containsSameElements := ContainSameElements(tt.sliceA, tt.sliceB)
		if containsSameElements != tt.result {
			t.Errorf("expected %v but got %v", tt.result, containsSameElements)
		}
	}
}
