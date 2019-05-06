package ints

import (
	"fmt"
)

func ExampleContains() {
	a := []int{1, 2, 3}

	fmt.Println(Contains(a, 1))
	fmt.Println(Contains(a, 4))
	fmt.Println(Contains(nil, 0))
	// Output:
	// true
	// false
	// false
}

func ExampleEquals() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	d := []int{}

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
