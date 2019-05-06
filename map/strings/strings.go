// Package strings provides utility function for map of string.
package strings

// Keys returns keys of the map
func Keys(m map[string]string) []string {
	a := make([]string, len(m))

	i := 0
	for s := range m {
		a[i] = s
		i++
	}

	return a
}

// Values returns values of the map
func Values(m map[string]string) []string {
	a := make([]string, len(m))

	i := 0
	for _, s := range m {
		a[i] = s
		i++
	}

	return a
}
