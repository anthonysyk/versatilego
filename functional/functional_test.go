package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsZero(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, IsZero(0))
	assert.Equal(t, false, IsZero(5))

	assert.Equal(t, true, IsZero(0.0))
	assert.Equal(t, false, IsZero(5.0))

	assert.Equal(t, true, IsZero(""))
	assert.Equal(t, false, IsZero("not empty"))

	assert.Equal(t, true, IsZero(false))
	assert.Equal(t, false, IsZero(true))
}

func TestIdentity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		input int
		want  int
	}{
		{"identity zero", 0, 0},
		{"identity positive int", 5, 5},
		{"identity negative int", -3, -3},
	}

	for _, c := range tests {
		c := c
		t.Run(c.label, func(t *testing.T) {
			t.Parallel()
			got := Identity(c.input)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label  string
		input  []int
		getter MapFunc[int, int]
		want   []int
	}{
		{
			label:  "SquareIntegers",
			input:  []int{1, 2, 3, 4},
			getter: func(x int) int { return x * x },
			want:   []int{1, 4, 9, 16},
		},
		{
			label:  "IdentityIntegers",
			input:  []int{1, 2, 3, 4},
			getter: Identity[int],
			want:   []int{1, 2, 3, 4},
		},
	}

	for _, c := range tests {
		c := c
		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			got := Map(c.input, c.getter)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	evenFunc := func(val int) bool {
		return val%2 == 0
	}
	evens := Filter([]int{1, 2, 3, 4, 5, 6}, evenFunc)
	wantEvens := []int{2, 4, 6}
	assert.Equal(t, wantEvens, evens)

	nonEmptyFunc := func(val string) bool {
		return val != ""
	}
	nonEmpty := Filter([]string{"hello", "", "world"}, nonEmptyFunc)
	wantNonEmpty := []string{"hello", "world"}
	assert.Equal(t, wantNonEmpty, nonEmpty)
}

type person struct {
	Name string
	Age  int
}

func TestFind(t *testing.T) {
	t.Parallel()

	people := []person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	result, found := Find(people, func(person person) bool {
		return person.Name == "Charlie"
	})
	assert.Equal(t, person{Name: "Charlie", Age: 22}, result)
	assert.Equal(t, true, found)

	result, found = Find(people, func(person person) bool {
		return person.Name == "David"
	})
	assert.Equal(t, person{}, result)
	assert.Equal(t, false, found)
}

func TestExists(t *testing.T) {
	t.Parallel()

	people := []person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	exists := Exists(people, func(person person) bool {
		return person.Name == "Bob"
	})
	assert.Equal(t, true, exists)

	exists = Exists(people, func(person person) bool {
		return person.Name == "David"
	})
	assert.Equal(t, false, exists)
}

func TestFlatMap(t *testing.T) {
	t.Parallel()

	baskets := [][]string{
		{"Apple", "Banana"},
		{"Orange", "Pineapple"},
		{"Kiwi"},
	}

	names := FlatMap(baskets, func(fruits []string) []string {
		return fruits
	})

	assert.Equal(t, []string{"Apple", "Banana", "Orange", "Pineapple", "Kiwi"}, names)
}
