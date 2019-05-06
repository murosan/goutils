// Package ints provides utility functions for slice of ints.
package ints

// Contains returns true if slice contains the int value.
func Contains(a []int, s int) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// Equals returns true if two slices are equal.
func Equals(a, b []int) bool {
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
