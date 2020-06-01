package gocom

import (
	"bufio"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestNextRational(t *testing.T) {
	type testCase struct {
		input  string
		output rational
	}
	for _, tc := range []testCase{
		testCase{input: "1.45", output: rational{frac: 29, denomi: 20}},
		testCase{input: "1.47", output: rational{frac: 147, denomi: 100}},
		testCase{input: "1.477", output: rational{frac: 1477, denomi: 1000}},
		testCase{input: ".45", output: rational{frac: 9, denomi: 20}},
		testCase{input: "0.45", output: rational{frac: 9, denomi: 20}},
		testCase{input: "10.5", output: rational{frac: 21, denomi: 2}},
		testCase{input: "145", output: rational{frac: 145, denomi: 1}},
		testCase{input: "2/3", output: rational{frac: 2, denomi: 3}},
		testCase{input: "8/35", output: rational{frac: 8, denomi: 35}},
		testCase{input: "-8/35", output: rational{neg: 1, frac: 8, denomi: 35}},
		testCase{input: "8/-36", output: rational{neg: 1, frac: 2, denomi: 9}},
		testCase{input: "-9/-35", output: rational{neg: 0, frac: 9, denomi: 35}},
		testCase{input: "-10.7", output: rational{neg: 1, frac: 107, denomi: 10}},
	} {
		got := nextRational(bufio.NewScanner(strings.NewReader(tc.input)))
		want := tc.output
		if got != want {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestRatinalOps(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := newRational(rand.Int(), rand.Int())
		ir := ratInv(r)
		want := newRational(1, 1)
		got := ratMul(r, ir)
		if !got.equal(want) {
			t.Fatalf("want %v, but %v:", want, got)
		}
	}
}
