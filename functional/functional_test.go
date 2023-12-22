package functional

import (
	"reflect"
	"testing"
)

func TestIdentity(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"IdentityPositiveInt", 5, 5},
		{"IdentityZero", 0, 0},
		{"IdentityNegativeInt", -3, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Identity(tt.input); got != tt.want {
				t.Errorf("Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		getter MapFunc[int, int]
		want   []int
	}{
		{
			name:   "SquareIntegers",
			input:  []int{1, 2, 3, 4},
			getter: func(x int) int { return x * x },
			want:   []int{1, 4, 9, 16},
		},
		{
			name:   "IdentityIntegers",
			input:  []int{1, 2, 3, 4},
			getter: Identity[int],
			want:   []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.input, tt.getter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	sumFunc := func(acc, val int) int {
		return acc + val
	}
	sum := Reduce([]int{1, 2, 3, 4}, sumFunc, 0)
	if sum != 10 {
		t.Errorf("Reduce() for sum returned %v, want %v", sum, 10)
	}

	concatFunc := func(acc string, val string) string {
		return acc + val
	}
	concat := Reduce([]string{"a", "b", "c"}, concatFunc, "")
	if concat != "abc" {
		t.Errorf("Reduce() for concatenation returned %v, want %v", concat, "abc")
	}
}

func TestFilter(t *testing.T) {
	evenFunc := func(val int) bool {
		return val%2 == 0
	}
	evens := Filter([]int{1, 2, 3, 4, 5, 6}, evenFunc)
	wantEvens := []int{2, 4, 6}
	if !reflect.DeepEqual(evens, wantEvens) {
		t.Errorf("Filter() for even numbers returned %v, want %v", evens, wantEvens)
	}

	nonEmptyFunc := func(val string) bool {
		return val != ""
	}
	nonEmpty := Filter([]string{"hello", "", "world"}, nonEmptyFunc)
	wantNonEmpty := []string{"hello", "world"}
	if !reflect.DeepEqual(nonEmpty, wantNonEmpty) {
		t.Errorf("Filter() for non-empty strings returned %v, want %v", nonEmpty, wantNonEmpty)
	}
}
