package sortvf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCaseSortByDate struct {
	Name     string
	Input    []time.Time
	SortDesc bool
	Expected []time.Time
}

func TestSortByDate(t *testing.T) {
	t.Parallel()
	testCases := []TestCaseSortByDate{
		{
			Name:     "AscendingOrderByDate",
			Input:    []time.Time{time.Unix(3, 0), time.Unix(1, 0), time.Unix(2, 0)},
			SortDesc: false,
			Expected: []time.Time{time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0)},
		},
		{
			Name:     "DescendingOrderByDate",
			Input:    []time.Time{time.Unix(1, 0), time.Unix(3, 0), time.Unix(2, 0)},
			SortDesc: true,
			Expected: []time.Time{time.Unix(3, 0), time.Unix(2, 0), time.Unix(1, 0)},
		},
		{
			Name:     "EmptySlice",
			Input:    []time.Time{},
			SortDesc: false,
			Expected: []time.Time{},
		},
		{
			Name:     "SingleElement",
			Input:    []time.Time{time.Unix(1, 0)},
			SortDesc: false,
			Expected: []time.Time{time.Unix(1, 0)},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			SortByDate(testCase.Input, func(i time.Time) time.Time {
				return i
			}, testCase.SortDesc)
			assert.Equal(t, testCase.Expected, testCase.Input)
		})
	}
}
