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

func parseProblem() int {
	n, m := scanInt(sc), scanInt(sc)
	xs := nextIntSlice(sc, m)
	sort.Sort(sort.IntSlice(xs))
	dfs := []pair{}
	for i := 1; i < len(xs); i++ {
		dfs = append(dfs, pair{p: xs[i], d: xs[i] - xs[i-1]})
	}
	sort.Sort(pairs(dfs))
	sum := 0
	num := 0
	registered := map[int]struct{}{}
	j := 0
	for i := 0; i < len(dfs); i++ {
		if num >= m-n {
			break
		}
		c := dfs[j]
		if _, ok := registered[c.p]; !ok {
			registered[c.p] = struct{}{}
			sum += c.d
			num++
		}
		j++
	}
	return sum
}

type pair struct{ p, d int }
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].d < p[j].d
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {

	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}
	return a
}

func resolve(n int) int {
	return n
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
