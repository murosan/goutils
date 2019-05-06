package strings

import (
	"sort"
	"testing"

	sslice "github.com/murosan/goutils/slice/strings"
)

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
