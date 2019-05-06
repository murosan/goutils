package ints

import (
	"testing"
)

func TestContains(t *testing.T) {
	cases := []struct {
		a        []int
		b        int
		expected bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{1}, 1, true},
		{[]int{}, 0, false},
		{nil, 0, false},
	}

	for i, c := range cases {
		res := Contains(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[ints > Contains]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestEquals(t *testing.T) {
	cases := []struct {
		a, b     []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{}, []int{}, true},
		{[]int{}, nil, false},
		{nil, []int{}, false},
		{nil, nil, true},
	}

	for i, c := range cases {
		res := Equals(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[ints > Equals]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}
