package strings

import (
	"fmt"
	"sort"
)

func ExampleEquals() {
	m1 := map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"}
	m2 := map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"}
	m3 := make(map[string]string)

	fmt.Println(Equals(m1, m2))
	fmt.Println(Equals(m1, m3))
	fmt.Println(Equals(m3, nil))
	fmt.Println(Equals(nil, nil))
	// Output:
	// true
	// false
	// false
	// true
}

func ExampleKeys() {
	m1 := map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"}
	m2 := make(map[string]string)

	r1 := Keys(m1)
	r2 := Keys(m2)
	r3 := Keys(nil)

	sort.Strings(r1)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	// Output:
	// [go js rb]
	// []
	// []
}

func ExampleValues() {
	m1 := map[string]string{"go": "Go", "rb": "Ruby", "js": "JavaScript"}
	m2 := make(map[string]string)

	r1 := Values(m1)
	r2 := Values(m2)
	r3 := Values(nil)

	sort.Strings(r1)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	// Output:
	// [Go JavaScript Ruby]
	// []
	// []
}
