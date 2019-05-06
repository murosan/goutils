package strings

import (
	"sort"
	"testing"

	sslice "github.com/murosan/goutils/slice/strings"
)

func TestEquals(t *testing.T) {
	cases := []struct {
		a, b     map[string]string
		expected bool
	}{
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			true,
		},
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go", "rb": "Ruby", "js": "Node.js"},
			false,
		},
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go"},
			false,
		},
		{nil, map[string]string{"go": "Go"}, false},
		{map[string]string{}, map[string]string{}, true},
	}

	for i, c := range cases {
		res := Equals(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[map > strings > Equals]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestKeys(t *testing.T) {
	cases := []struct {
		in       map[string]string
		expected []string
	}{
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			[]string{"go", "rb", "js"},
		},
		{map[string]string{}, []string{}},
		{nil, []string{}},
	}

	for i, c := range cases {
		res := Keys(c.in)
		sort.Strings(res)
		sort.Strings(c.expected)

		if sslice.NotEqual(res, c.expected) {
			t.Errorf(`
[map > strings > Keys]
Index:    %d
Input:    %v
Expected: %s
Actual:   %s
`, i, c.in, res, c.expected)
		}
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		a, b     map[string]string
		expected bool
	}{
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			false,
		},
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go", "rb": "Ruby", "js": "Node.js"},
			true,
		},
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			map[string]string{"go": "Go"},
			true,
		},
		{nil, map[string]string{"go": "Go"}, true},
		{map[string]string{}, map[string]string{}, false},
	}

	for i, c := range cases {
		res := NotEqual(c.a, c.b)
		if res != c.expected {
			t.Errorf(`
[map > strings > NotEqual]
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestValues(t *testing.T) {
	cases := []struct {
		in       map[string]string
		expected []string
	}{
		{
			map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"},
			[]string{"Go", "Ruby", "JavaScript"},
		},
		{map[string]string{}, []string{}},
		{nil, []string{}},
	}

	for i, c := range cases {
		res := Values(c.in)
		sort.Strings(res)
		sort.Strings(c.expected)

		if sslice.NotEqual(res, c.expected) {
			t.Errorf(`
[map > strings > Values]
Index:    %d
Input:    %v
Expected: %s
Actual:   %s
`, i, c.in, res, c.expected)
		}
	}
}
