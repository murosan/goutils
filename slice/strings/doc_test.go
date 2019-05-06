package strings

import (
	"fmt"
)

func ExampleContains() {
	a := []string{"a", "b", "c"}

	fmt.Println(Contains(a, "a"))
	fmt.Println(Contains(a, "d"))
	fmt.Println(Contains(nil, ""))
	// Output:
	// true
	// false
	// false
}

func ExampleEquals() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	c := []string{"a", "b", ""}
	d := []string{}

	fmt.Println(Equals(a, b))
	fmt.Println(Equals(a, c))
	fmt.Println(Equals(a, d))
	fmt.Println(Equals(nil, a))
	fmt.Println(Equals(nil, nil))
	// Output:
	// true
	// false
	// false
	// false
	// true
}
