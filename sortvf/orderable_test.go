package sortvf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSortByInt struct {
	Name     string
	Input    []int
	SortDesc bool
	Expected []int
}

func TestSortByInt(t *testing.T) {
	t.Parallel()
	testCases := []testCaseSortByInt{
		{
			Name:     "AscendingOrder",
			Input:    []int{3, 1, 4, 5, 2},
			SortDesc: false,
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Name:     "DescendingOrder",
			Input:    []int{3, 1, 4, 5, 2},
			SortDesc: true,
			Expected: []int{5, 4, 3, 2, 1},
		},
		{
			Name:     "EmptySlice",
			Input:    []int{},
			SortDesc: false,
			Expected: []int{},
		},
		{
			Name:     "SingleElement",
			Input:    []int{0},
			SortDesc: false,
			Expected: []int{0},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			SortByOrderable(testCase.Input, func(i int) int {
				return i
			}, testCase.SortDesc)
			assert.Equal(t, testCase.Expected, testCase.Input)
		})
	}
}
