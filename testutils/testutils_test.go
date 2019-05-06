package testutils

import (
	"testing"
)

func TestMustPanic(t *testing.T) {
	cases := []struct {
		fn                   func()
		onFailShouldBeCalled bool
	}{
		{func() { panic("panic test") }, false},
		{func() {}, true},
	}

	for i, c := range cases {
		calledOnFail := false
		onFail := func(t *testing.T) {
			calledOnFail = true
		}

		MustPanic(t, c.fn, onFail)

		if calledOnFail != c.onFailShouldBeCalled {
			t.Errorf(`
[testutils > MustPanic]
Index:    %d
Expected: %t
Actual:   %t
`, i, c.onFailShouldBeCalled, calledOnFail)
		}
	}
}
