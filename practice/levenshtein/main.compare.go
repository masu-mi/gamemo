package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	comp := parseProblem(os.Stdin)
	dist, ops := comp.comparePattern()
	fmt.Printf("pat: `%s`, text: `%s`\n", comp.pat, comp.text)
	fmt.Printf("dist: %d\n", dist)
	operations(ops).WriteTo(os.Stdout)
}

type operations []opLog

func (log operations) WriteTo(w io.Writer) {
	fmt.Fprintf(w, "operation: ")
	for _, log := range log {
		switch log.operation {
		case match:
			if log.p == log.t {
				fmt.Fprintf(w, "%s", log.p)
			} else {
				fmt.Fprintf(w, "(%s->%s)", log.p, log.t)
			}
		case insert:
			fmt.Fprintf(w, "(ins:%s)", log.t)
		case del:
			fmt.Fprintf(w, "(del:%s)", log.p)
		}
	}
	fmt.Fprintln(w, "")
}

type op int

const (
	noop op = -1 + iota
	match
	insert
	del
)

type opLog struct {
	p, t      string
	operation op
}

func (comp *comparingWithLevenshtein) comparePattern() (dist int, result []opLog) {
	for i := 1; i <= len(comp.pat); i++ {
		for j := 1; j <= len(comp.text); j++ {
			comp.m[i][j] = comp.findOpHasMinimumCost(i, j)
		}
	}

	i, j := len(comp.pat), len(comp.text)
	dist = comp.m[i][j].cost
LOOP:
	for true {
		if comp.m[i][j].operation == noop {
			break LOOP
		}
		result = append(result, opLog{
			p:         comp.pat[i-1 : i],
			t:         comp.text[j-1 : j],
			operation: comp.m[i][j].operation,
		})
		switch comp.m[i][j].operation {
		case match:
			i, j = i-1, j-1
		case insert:
			j = j - 1
		case del:
			i = i - 1
		}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return
}

func (comp *comparingWithLevenshtein) findOpHasMinimumCost(patI, textI int) (result cell) {
	costs := map[op]int{}
	if comp.pat[patI-1] == comp.text[textI-1] {
		costs[match] = comp.m[patI-1][textI-1].cost + 0
	} else {
		costs[match] = comp.m[patI-1][textI-1].cost + 1
	}
	costs[insert] = comp.m[patI][textI-1].cost + 1
	costs[del] = comp.m[patI-1][textI].cost + 1
	var min int = math.MaxInt32
	for op := range costs {
		if costs[op] <= min {
			min = costs[op]
			result = cell{
				cost:      costs[op],
				operation: op,
			}
		}
	}
	return result
}

func parseProblem(r io.Reader) *comparingWithLevenshtein {
	var pat, text string
	fmt.Fscan(r, &pat, &text)
	return newComparingWithLevenshtein(pat, text)
}

type comparingWithLevenshtein struct {
	pat, text string
	m         [][]cell
}

type cell struct {
	cost      int
	operation op
}

func newComparingWithLevenshtein(pat, text string) *comparingWithLevenshtein {
	m := createMap(len(pat), len(text))
	initMap(m, len(pat), len(text))
	return &comparingWithLevenshtein{
		pat:  pat,
		text: text,
		m:    m,
	}
}

func createMap(pl, tl int) [][]cell {
	m := make([][]cell, pl+1)
	for i := range m {
		m[i] = make([]cell, tl+1)
	}
	return m
}

func initMap(m [][]cell, pl, tl int) {
	m[0][0] = cell{cost: 0, operation: noop}
	for i := 1; i <= pl; i++ {
		m[i][0].cost = i
		m[i][0].operation = del
	}
	for i := 1; i <= tl; i++ {
		m[0][i].cost = i
		m[0][i].operation = insert
	}
}
