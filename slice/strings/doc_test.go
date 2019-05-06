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

func ExampleNotContain() {
	a := []string{"a", "b", "c"}

	fmt.Println(NotContain(a, "a"))
	fmt.Println(NotContain(a, "d"))
	fmt.Println(NotContain(nil, ""))
	// Output:
	// false
	// true
	// true
}

func ExampleNotEqual() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	c := []string{"a", "b", ""}
	d := []string{}

	fmt.Println(NotEqual(a, b))
	fmt.Println(NotEqual(a, c))
	fmt.Println(NotEqual(a, d))
	fmt.Println(NotEqual(nil, a))
	fmt.Println(NotEqual(nil, nil))
	// Output:
	// false
	// true
	// true
	// true
	// false
}
