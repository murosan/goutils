// Package strings provides utility functions for slice of string.
package strings

// Contains returns true if slice contains the string.
func Contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// Equals returns true if two slices are equal.
func Equals(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

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

// NotContain returns true if slice does not contain the string.
func NotContain(a []string, s string) bool {
	return !Contains(a, s)
}

// NotEqual returns true if two slices are not equal.
func NotEqual(a, b []string) bool {
	return !Equals(a, b)
}
