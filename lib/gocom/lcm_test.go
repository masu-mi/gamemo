package gocom

import (
	"testing"
)

func TestSimple(t *testing.T) {
	type testCase struct {
		a, b, want int
	}
	for _, test := range []testCase{
		testCase{1, 2, 2},
		testCase{2, 2, 2},
		testCase{2, 3, 6},
		testCase{4, 3, 12},
		testCase{4, 6, 12},
		testCase{9, 6, 18},
	} {
		got := lcm(test.a, test.b)
		if got != test.want {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}
