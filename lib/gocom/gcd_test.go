package gocom

import (
	"testing"
)

func TestGcd(t *testing.T) {
	type testCase struct {
		a, b, want int
	}

	for _, test := range []testCase{
		testCase{2, 3, 1},
		testCase{1, 3, 1},
		testCase{3, 3, 3},
		testCase{2, 6, 2},
		testCase{4, 6, 2},
		testCase{9, 6, 3},
	} {
		got := gcd(test.a, test.b)
		if got != test.want {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}
