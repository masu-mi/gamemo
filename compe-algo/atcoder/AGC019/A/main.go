package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 1000000
)

var sc *bufio.Scanner

func initScanner(r io.Reader) *bufio.Scanner {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	return sc
}

func main() {
	sc = initScanner(os.Stdin)
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (int, int, int, int) {
	q, h, s, d := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
	return q, h, s, d
}

type pair struct{ f, s int }

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return (8/p[i].f)*p[i].s < (8/p[j].f)*p[j].s
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func resolve(q, h, s, d int) int {
	n := scanInt(sc) * 4
	cands := []pair{{1, q}, {2, h}, {4, s}, {8, d}}
	sort.Sort(pairs(cands))
	costs := 0
	for _, p := range cands {
		div := n / p.f
		costs += div * p.s
		n %= p.f
	}
	return costs
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
