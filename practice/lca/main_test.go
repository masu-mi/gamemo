package lca

import (
	"testing"

	"github.com/k0kubun/pp"
)

type pair struct{ p, c int }

func TestSimple(t *testing.T) {
	type testCase struct {
		size   int
		edges  []pair
		input  pair
		output int
	}
	for _, tc := range []testCase{
		testCase{
			size:   1,
			input:  pair{0, 0},
			output: 0,
		},
		testCase{
			size:   5,
			edges:  []pair{pair{0, 1}, pair{1, 2}, pair{1, 3}, pair{2, 4}},
			input:  pair{3, 4},
			output: 1,
		},
		testCase{
			size:   5,
			edges:  []pair{pair{0, 1}, pair{1, 2}, pair{2, 3}, pair{2, 4}},
			input:  pair{3, 4},
			output: 2,
		},
	} {
		tr := newTree(tc.size)
		for _, e := range tc.edges {
			tr.addEdge(e.p, e.c)
		}
		pp.Println(tr)
		got := tr.lca(tc.input.p, tc.input.c)
		want := tc.output
		if got != want {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}
