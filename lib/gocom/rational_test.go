package gocom

import (
	"bufio"
	"strings"
	"testing"
)

func TestNextRational(t *testing.T) {
	type testCase struct {
		input  string
		output rational
	}
	for _, tc := range []testCase{
		testCase{input: "1.45", output: rational{frac: 145, denomi: 100}},
		testCase{input: ".45", output: rational{frac: 45, denomi: 100}},
		testCase{input: ".45", output: rational{frac: 45, denomi: 100}},
		testCase{input: "0.45", output: rational{frac: 45, denomi: 100}},
		testCase{input: "10.5", output: rational{frac: 105, denomi: 10}},
		testCase{input: "145", output: rational{frac: 145, denomi: 1}},
		testCase{input: "2/3", output: rational{frac: 2, denomi: 3}},
		testCase{input: "8/35", output: rational{frac: 8, denomi: 35}},
	} {
		got := nextRational(bufio.NewScanner(strings.NewReader(tc.input)))
		want := tc.output
		if got != want {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}
