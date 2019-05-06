package strings

import (
	"testing"
)

func TestContains(t *testing.T) {
	cases := []struct {
		a        []string
		b        string
		expected bool
	}{
		{[]string{"one", "two", "three"}, "one", true},
		{[]string{"one", "two", "three"}, "threee", false},
		{[]string{"one"}, "one", true},
		{[]string{}, "", false},
		{nil, "", false},
	}

	for i, c := range cases {
		res := Contains(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[strings > Contains]
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
		a, b     []string
		expected bool
	}{
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}, true},
		{[]string{"one", "two", "three"}, []string{"one", "two", "threee"}, false},
		{[]string{"one"}, []string{"one"}, true},
		{[]string{"one", "two"}, []string{"one", "two", "three"}, false},
		{[]string{}, []string{}, true},
		{[]string{}, nil, false},
		{nil, []string{}, false},
		{nil, nil, true},
	}

	for i, c := range cases {
		res := Equals(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[strings > Equals]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestNotContain(t *testing.T) {
	cases := []struct {
		a        []string
		b        string
		expected bool
	}{
		{[]string{"one", "two", "three"}, "one", false},
		{[]string{"one", "two", "three"}, "threee", true},
		{[]string{"one"}, "one", false},
		{[]string{}, "", true},
		{nil, "", true},
	}

	for i, c := range cases {
		res := NotContain(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[strings > NotContain]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		a, b     []string
		expected bool
	}{
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}, false},
		{[]string{"one", "two", "three"}, []string{"one", "two", "threee"}, true},
		{[]string{"one"}, []string{"one"}, false},
		{[]string{"one", "two"}, []string{"one", "two", "three"}, true},
		{[]string{}, []string{}, false},
		{[]string{}, nil, true},
		{nil, []string{}, true},
		{nil, nil, false},
	}

	for i, c := range cases {
		res := NotEqual(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[strings > NotEqual]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}
