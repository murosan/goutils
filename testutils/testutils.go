// Package testutils provides utility functions for testing.
package testutils

import (
	"testing"
)

// MustPanic checks if method 'f' causes panic.
// If there was no panic, calls 'onFail'.
//
// Actual method
//
//   func PanicMethod() {
//    	panic("")
//   }
//
// Usage
//
//   func TestPanicMethod(t *testing.T) {
//   	onFail := func(t *testing.T) {
//   		t.Helper()
//   		t.Error("Expected panic, but there wasn't.")
//   	}
//   	MustPanic(t, PanicMethod, onFail)
//   }
func MustPanic(t *testing.T, f func(), onFail func(t *testing.T)) {
	t.Helper()

	defer func() {
		// test succeeded
		recover()
	}()

	// expected panic
	f()

	// test failed for no panic
	onFail(t)
}
