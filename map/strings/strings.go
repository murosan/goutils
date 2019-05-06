// Package strings provides utility function for map of string.
package strings

// Equals returns if the maps have the same key-value pairs.
func Equals(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}

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

// NotEqual returns if the maps does not have the same key-value pairs.
func NotEqual(a, b map[string]string) bool {
	return !Equals(a, b)
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
