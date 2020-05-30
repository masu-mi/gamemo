package gocom

import (
	"testing"
)

func TestCompressedID(t *testing.T) {
	type testCase struct {
		input    []int
		expected map[int]int
	}
	for _, tc := range []testCase{
		testCase{
			input: []int{100, 200, 300, 200, 150},
			expected: map[int]int{
				0: 100,
				1: 150,
				2: 200,
				3: 300,
			},
		},
	} {
		ci := newCompressedID()
		for _, name := range tc.input {
			ci.addItem(name)
		}
		ci.build()
		for k, v := range ci.l {
			got := ci.name(k)
			want := tc.expected[k]
			if got != want {
				t.Fatalf("ci.name(); want %v, but %v:", want, got)
			}
			got = ci.id(v)
			want = k
			if got != want {
				t.Fatalf("ci.id(); want %v, but %v:", want, got)
			}
		}
	}
}
